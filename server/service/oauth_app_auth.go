package service

import (
	"fmt"
	"github.com/sanmuyan/xpkg/xjwt"
	"github.com/sanmuyan/xpkg/xutil"
	"strings"
	"time"
	"wukong/pkg/config"
	"wukong/pkg/datastore"
	"wukong/pkg/db"
	"wukong/pkg/util"
	"wukong/server/model"
)

func (s *Service) GetOauthCodeSession(token *model.Token, req *model.OauthCodeSessionRequest) (string, *model.OauthErrorResponse) {
	var code string
	var oauthApp model.OauthApp
	if req.ResponseType != "code" {
		return code, model.NewOauthErrorResponse("unsupported_response_type")
	}
	tx := db.DB.Where("client_id = ?", req.ClientID).First(&oauthApp)
	if tx.RowsAffected == 0 {
		return code, model.NewOauthErrorResponse("invalid_client_id")
	}
	if !xutil.IsContains(req.RedirectURI, strings.Split(oauthApp.RedirectURI, ",")) {
		return code, model.NewOauthErrorResponse("invalid_redirect_uri")
	}
	appScopes := strings.Split(oauthApp.Scope, ",")
	scopes := strings.Split(req.Scope, " ")
	for _, _scope := range scopes {
		if !xutil.IsContains(_scope, appScopes) {
			return code, model.NewOauthErrorResponse("invalid_scope")
		}
	}
	sessionID := util.GetRandomID()
	code = sessionID
	oauthCode := model.OauthCodeSession{
		Code:         sessionID,
		ClientID:     req.ClientID,
		ClientSecret: oauthApp.ClientSecret,
		RedirectURI:  req.RedirectURI,
		Scope:        strings.Replace(req.Scope, " ", ",", -1),
	}
	err := datastore.DS.StoreSession(model.NewSession(sessionID, model.SessionTypeOAuthCode, token.GetUserID(), token.Username,
		&oauthCode).SetTimeout(config.OauthCodeTimeoutMin * time.Minute))
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
	var oauthCodeSession model.OauthCodeSession
	session, ok := datastore.DS.LoadSession(req.Code, model.SessionTypeOAuthCode, &oauthCodeSession)
	if !ok {
		return nil, model.NewOauthErrorResponse("invalid_code")
	}
	defer func() {
		_ = datastore.DS.DeleteSession(req.Code, model.SessionTypeOAuthCode)
	}()
	if oauthCodeSession.ClientID != req.ClientID {
		return nil, model.NewOauthErrorResponse("invalid_client")
	}
	if oauthCodeSession.RedirectURI != req.RedirectURI {
		return nil, model.NewOauthErrorResponse("invalid_redirect_uri")
	}
	if req.ClientSecret != "" {
		if req.ClientSecret != oauthCodeSession.ClientSecret {
			return nil, model.NewOauthErrorResponse("invalid_client_secret")
		}
	}
	oauthTokenResponse, err := s.createOauthToken(&oauthCodeSession, session.Username)
	if err != nil {
		return nil, model.NewOauthErrorResponse("server_error")
	}
	return oauthTokenResponse, nil
}

func (s *Service) createOauthToken(oauthCode *model.OauthCodeSession, username string) (*model.OauthTokenResponse, error) {
	accessToken, err := s.createOrSetOauthToken(username, model.TokenTypeOauthAccess, oauthCode.Scope, oauthCode.ClientID, config.AppAccessTokenExpiration)
	if err != nil {
		return nil, err
	}
	refreshToken, err := s.createOrSetOauthToken(username, model.TokenTypeOauthRefresh, oauthCode.Scope, oauthCode.ClientID, config.Conf.TokenTTL)
	if err != nil {
		return nil, err
	}
	oauthToken := model.OauthTokenResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
		ExpiresIn:    config.AppAccessTokenExpiration,
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
	tokenStr, err := s.CreateOrSetToken(&token, expiresAt)
	if err != nil {
		return tokenStr, err
	}
	return tokenStr, nil
}

func (s *Service) validateOauthRefreshToken(refreshToken, clientID, clientSecret string) (*model.Token, *model.OauthErrorResponse) {
	var oauthApp model.OauthApp
	var token model.Token
	_, err := xjwt.ParseToken(refreshToken, config.Conf.Secret.TokenKey, &token)
	if err != nil {
		return nil, model.NewOauthErrorResponse("invalid_token")
	}
	err = db.DB.Where("client_id = ?", clientID).First(&oauthApp).Error
	if err != nil {
		return nil, model.NewOauthErrorResponse("invalid_client")
	}
	if clientSecret != "" {
		if clientSecret != oauthApp.ClientSecret {
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
	if token.TokenType != model.TokenTypeOauthRefresh {
		return nil, model.NewOauthErrorResponse("invalid_token")
	}
	accessToken, err := s.createOrSetOauthToken(token.Username, model.TokenTypeOauthAccess, token.Scope, req.ClientID, config.AppAccessTokenExpiration)
	if err != nil {
		return nil, model.NewOauthErrorResponse("server_error")
	}
	return &model.OauthTokenResponse{
		AccessToken: accessToken,
		ExpiresIn:   config.AppAccessTokenExpiration,
		TokenType:   "Bearer",
	}, nil
}

func (s *Service) RevokeOauthToken(req *model.OauthRevokeTokenRequest) *model.OauthErrorResponse {
	token, _err := s.validateOauthRefreshToken(req.Token, req.ClientID, req.ClientSecret)
	if _err != nil {
		return _err
	}
	if token.TokenType != model.TokenTypeOauthRefresh {
		return model.NewOauthErrorResponse("invalid_token")
	}
	err := datastore.DS.DeleteSession(token.TokenID, token.TokenType, fmt.Sprintf("%s:%s:%s", token.TokenType, token.Username, token.TokenID))
	if err != nil {
		return model.NewOauthErrorResponse("server_error")
	}
	return nil
}
