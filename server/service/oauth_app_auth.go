package service

import (
	"errors"
	"fmt"
	"github.com/google/uuid"
	"github.com/sanmuyan/xpkg/xjwt"
	"github.com/sanmuyan/xpkg/xresponse"
	"github.com/sanmuyan/xpkg/xutil"
	"strings"
	"time"
	"wukong/pkg/config"
	"wukong/pkg/datastorage"
	"wukong/pkg/db"
	"wukong/pkg/util"
	"wukong/server/model"
)

const (
	accessTokenExpiration = 3600
)

func (s *Service) GetOauthCode(token *model.Token, clientID, redirectURI, responseType, scope, state string) (string, util.RespError) {
	var code string
	var oauthAPP model.OauthAPP
	if responseType != "code" {
		return code, util.NewRespError(errors.New("不支持的 response_type")).WithCode(xresponse.HttpBadRequest)
	}
	tx := db.DB.Where("client_id = ?", clientID).First(&oauthAPP)
	if tx.RowsAffected == 0 {
		return code, util.NewRespError(errors.New("应用不存在")).WithCode(xresponse.HttpBadRequest)
	}
	if !xutil.IsContains(redirectURI, strings.Split(oauthAPP.RedirectURI, ",")) {
		return code, util.NewRespError(errors.New("回调地址不正确")).WithCode(xresponse.HttpBadRequest)
	}
	appScopes := strings.Split(oauthAPP.Scope, ",")
	scopes := strings.Split(scope, " ")
	for _, _scope := range scopes {
		if !xutil.IsContains(_scope, appScopes) {
			return code, util.NewRespError(errors.New("不支持的 scope")).WithCode(xresponse.HttpBadRequest)
		}
	}
	code = uuid.NewString()
	oauthCode := model.OauthCode{
		Code:         code,
		Username:     token.Username,
		ClientID:     clientID,
		ClientSecret: oauthAPP.ClientSecret,
		RedirectURI:  redirectURI,
		Scope:        strings.Replace(scope, " ", ",", -1),
		ExpiresAt:    time.Now().UTC().Add(5 * time.Minute),
	}
	err := datastorage.DS.StoreCode(&oauthCode)
	if err != nil {
		return code, util.NewRespError(err)
	}
	if state != "" {
		return fmt.Sprintf("%s?code=%s&state=%s", redirectURI, code, state), nil
	}
	return fmt.Sprintf("%s?code=%s", redirectURI, code), nil
}

func (s *Service) GetOauthToken(code, clientID, clientSecret, redirectURI, grantType string) (*model.OauthTokenResponse, util.RespError) {
	if grantType != "authorization_code" {
		return nil, util.NewRespError(errors.New("不支持的 grant_type")).WithCode(xresponse.HttpBadRequest)
	}
	oauthCode, err := datastorage.DS.LoadCode(code, clientID)
	if err != nil {
		return nil, util.NewRespError(errors.New("无效的 code")).WithCode(xresponse.HttpBadRequest)
	}
	defer func() {
		_ = datastorage.DS.DeleteCode(code, clientID)
	}()
	if oauthCode.ExpiresAt.Before(time.Now().UTC()) {
		return nil, util.NewRespError(errors.New("code 已过期")).WithCode(xresponse.HttpBadRequest)
	}
	if oauthCode.ClientID != clientID {
		return nil, util.NewRespError(errors.New("无效的 client_id")).WithCode(xresponse.HttpBadRequest)
	}
	if oauthCode.RedirectURI != redirectURI {
		return nil, util.NewRespError(errors.New("无效的 redirect_uri")).WithCode(xresponse.HttpBadRequest)
	}
	if clientSecret != "" {
		if clientSecret != oauthCode.ClientSecret {
			return nil, util.NewRespError(errors.New("无效的 client_secret")).WithCode(xresponse.HttpBadRequest)
		}
	}
	oauthTokenResponse, err := s.createOauthToken(oauthCode)
	if err != nil {
		return nil, util.NewRespError(err)
	}
	return oauthTokenResponse, nil
}

func (s *Service) createOauthToken(oauthCode *model.OauthCode) (*model.OauthTokenResponse, error) {
	accessToken, err := s.createOrSetOauthToken(oauthCode.Username, model.OauthAccessToken, oauthCode.Scope, oauthCode.ClientID, accessTokenExpiration)
	if err != nil {
		return nil, err
	}
	refreshToken, err := s.createOrSetOauthToken(oauthCode.Username, model.OauthRefreshToken, oauthCode.Scope, oauthCode.ClientID, config.Conf.TokenTTL)
	if err != nil {
		return nil, err
	}
	oauthToken := model.OauthTokenResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
		ExpiresIn:    accessTokenExpiration,
		TokenType:    "Bearer",
	}
	return &oauthToken, nil
}

func (s *Service) createOrSetOauthToken(username, tokenType, scope, clientID string, expiresAt int) (string, error) {
	var token model.Token
	var tokenStr string
	token.Username = username
	token.TokenType = tokenType
	token.Scope = scope
	token.ClientID = clientID
	tokenStr, err := s.CreateOrSetToken(&token, s.generateOauthTokenID(&token), expiresAt)
	if err != nil {
		return tokenStr, err
	}
	return tokenStr, nil
}

func (s *Service) validateOauthRefreshToken(refreshToken, clientID, clientSecret string) (*model.Token, error) {
	var oauthAPP model.OauthAPP
	var token model.Token
	_, err := xjwt.ParseToken(refreshToken, config.Conf.Secret.TokenID, &token)
	if err != nil {
		return nil, err
	}
	err = db.DB.Where("client_id = ?", clientID).First(&oauthAPP).Error
	if err != nil {
		return nil, err
	}
	if clientSecret != "" {
		if clientSecret != oauthAPP.ClientSecret {
			return nil, errors.New("无效的 client_secret")
		}
	}
	return &token, nil
}

func (s *Service) RefreshOauthToken(refreshToken, clientID, grantType, clientSecret string) (*model.OauthTokenResponse, util.RespError) {
	if grantType != "refresh_token" {
		return nil, util.NewRespError(errors.New("不支持的 grant_type")).WithCode(xresponse.HttpBadRequest)
	}
	token, err := s.validateOauthRefreshToken(refreshToken, clientID, clientSecret)
	if err != nil {
		return nil, util.NewRespError(err)
	}
	accessToken, err := s.createOrSetOauthToken(token.Username, model.OauthAccessToken, token.Scope, clientID, accessTokenExpiration)
	if err != nil {
		return nil, util.NewRespError(err)
	}
	return &model.OauthTokenResponse{
		AccessToken: accessToken,
		ExpiresIn:   accessTokenExpiration,
		TokenType:   "Bearer",
	}, nil
}

func (s *Service) RevokeOauthToken(refreshToken, clientID, clientSecret string) util.RespError {
	token, err := s.validateOauthRefreshToken(refreshToken, clientID, clientSecret)
	if err != nil {
		return util.NewRespError(err)
	}
	err = datastorage.DS.DeleteToken(s.generateOauthTokenID(token), model.OauthRefreshToken)
	if err != nil {
		return util.NewRespError(err)
	}
	return nil
}

func (s *Service) generateOauthTokenID(token *model.Token) string {
	return fmt.Sprintf("%s_app_%s", token.Username, token.ClientID)

}
