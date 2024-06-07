package service

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/sanmuyan/xpkg/xjwt"
	"github.com/sanmuyan/xpkg/xutil"
	"strings"
	"time"
	"wukong/pkg/config"
	"wukong/pkg/datastorage"
	"wukong/pkg/db"
	"wukong/server/model"
)

const (
	accessTokenExpiration = 3600
)

func (s *Service) GetOauthRedirectURI(req *model.OauthCodeRequest) (string, *model.OauthErrorResponse) {
	var oauthAPP model.OauthAPP
	if req.ResponseType != "code" {
		return "", model.NewOauthErrorResponse("unsupported_response_type")
	}
	tx := db.DB.Where("client_id = ?", req.ClientID).First(&oauthAPP)
	if tx.RowsAffected == 0 {
		return "", model.NewOauthErrorResponse("invalid_client_id")
	}
	if !xutil.IsContains(req.RedirectURI, strings.Split(oauthAPP.RedirectURI, ",")) {
		return "", model.NewOauthErrorResponse("invalid_redirect_uri")
	}
	appScopes := strings.Split(oauthAPP.Scope, ",")
	scopes := strings.Split(req.Scope, " ")
	for _, _scope := range scopes {
		if !xutil.IsContains(_scope, appScopes) {
			return "", model.NewOauthErrorResponse("invalid_scope")
		}
	}
	if req.State != "" {
		return fmt.Sprintf("%s?state=%s", req.RedirectURI, req.State), nil
	}
	return fmt.Sprintf("%s?state=", req.RedirectURI), nil
}

func (s *Service) GetOauthCode(token *model.Token, req *model.OauthCodeRequest) (string, *model.OauthErrorResponse) {
	var code string
	var oauthAPP model.OauthAPP
	if req.ResponseType != "code" {
		return code, model.NewOauthErrorResponse("unsupported_response_type")
	}
	tx := db.DB.Where("client_id = ?", req.ClientID).First(&oauthAPP)
	if tx.RowsAffected == 0 {
		return code, model.NewOauthErrorResponse("invalid_client_id")
	}
	if !xutil.IsContains(req.RedirectURI, strings.Split(oauthAPP.RedirectURI, ",")) {
		return code, model.NewOauthErrorResponse("invalid_redirect_uri")
	}
	appScopes := strings.Split(oauthAPP.Scope, ",")
	scopes := strings.Split(req.Scope, " ")
	for _, _scope := range scopes {
		if !xutil.IsContains(_scope, appScopes) {
			return code, model.NewOauthErrorResponse("invalid_scope")
		}
	}
	code = uuid.NewString()
	oauthCode := model.OauthCode{
		Code:         code,
		Username:     token.Username,
		ClientID:     req.ClientID,
		ClientSecret: oauthAPP.ClientSecret,
		RedirectURI:  req.RedirectURI,
		Scope:        strings.Replace(req.Scope, " ", ",", -1),
		ExpiresAt:    time.Now().UTC().Add(5 * time.Minute),
	}
	err := datastorage.DS.StoreCode(&oauthCode)
	if err != nil {
		return code, model.NewOauthErrorResponse("server_error")
	}
	if req.State != "" {
		return fmt.Sprintf("%s?code=%s&state=%s", req.RedirectURI, code, req.State), nil
	}
	return fmt.Sprintf("%s?code=%s", req.RedirectURI, code), nil
}

func (s *Service) GetOauthToken(req *model.OauthTokenRequest) (*model.OauthTokenResponse, *model.OauthErrorResponse) {
	if req.GrantType != "authorization_code" {
		return nil, model.NewOauthErrorResponse("unsupported_grant_type")
	}
	oauthCode, err := datastorage.DS.LoadCode(req.Code, req.ClientID)
	if err != nil {
		return nil, model.NewOauthErrorResponse("invalid_code")
	}
	defer func() {
		_ = datastorage.DS.DeleteCode(req.Code, req.ClientID)
	}()
	if oauthCode.ExpiresAt.Before(time.Now().UTC()) {
		return nil, model.NewOauthErrorResponse("invalid_code")
	}
	if oauthCode.ClientID != req.ClientID {
		return nil, model.NewOauthErrorResponse("invalid_client")
	}
	if oauthCode.RedirectURI != req.RedirectURI {
		return nil, model.NewOauthErrorResponse("invalid_redirect_uri")
	}
	if req.ClientSecret != "" {
		if req.ClientSecret != oauthCode.ClientSecret {
			return nil, model.NewOauthErrorResponse("invalid_client_secret")
		}
	}
	oauthTokenResponse, err := s.createOauthToken(oauthCode)
	if err != nil {
		return nil, model.NewOauthErrorResponse("server_error")
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

func (s *Service) validateOauthRefreshToken(refreshToken, clientID, clientSecret string) (*model.Token, *model.OauthErrorResponse) {
	var oauthAPP model.OauthAPP
	var token model.Token
	_, err := xjwt.ParseToken(refreshToken, config.Conf.Secret.TokenID, &token)
	if err != nil {
		return nil, model.NewOauthErrorResponse("invalid_token")
	}
	err = db.DB.Where("client_id = ?", clientID).First(&oauthAPP).Error
	if err != nil {
		return nil, model.NewOauthErrorResponse("invalid_client")
	}
	if clientSecret != "" {
		if clientSecret != oauthAPP.ClientSecret {
			return nil, model.NewOauthErrorResponse("invalid_client_secret")
		}
	}
	return &token, nil
}

func (s *Service) RefreshOauthToken(req *model.OauthTokenRequest) (*model.OauthTokenResponse, *model.OauthErrorResponse) {
	if req.GrantType != "refresh_token" {
		return nil, model.NewOauthErrorResponse("unsupported_grant_type")
	}
	token, _err := s.validateOauthRefreshToken(req.RefreshToken, req.ClientID, req.ClientSecret)
	if _err != nil {
		return nil, _err
	}
	accessToken, err := s.createOrSetOauthToken(token.Username, model.OauthAccessToken, token.Scope, req.ClientID, accessTokenExpiration)
	if err != nil {
		return nil, model.NewOauthErrorResponse("server_error")
	}
	return &model.OauthTokenResponse{
		AccessToken: accessToken,
		ExpiresIn:   accessTokenExpiration,
		TokenType:   "Bearer",
	}, nil
}

func (s *Service) RevokeOauthToken(req *model.OauthRevokeTokenRequest) *model.OauthErrorResponse {
	token, _err := s.validateOauthRefreshToken(req.Token, req.ClientID, req.ClientSecret)
	if _err != nil {
		return _err
	}
	err := datastorage.DS.DeleteToken(s.generateOauthTokenID(token), model.OauthRefreshToken)
	if err != nil {
		return model.NewOauthErrorResponse("server_error")
	}
	return nil
}

func (s *Service) generateOauthTokenID(token *model.Token) string {
	return fmt.Sprintf("%s_app_%s", token.Username, token.ClientID)

}
