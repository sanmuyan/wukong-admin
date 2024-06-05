package service

import (
	"context"
	"errors"
	"fmt"
	"github.com/sanmuyan/xpkg/xresponse"
	"github.com/sirupsen/logrus"
	"golang.org/x/oauth2"
	"io"
	"wukong/pkg/config"
	"wukong/pkg/db"
	useroAuth "wukong/pkg/useroauth"
	"wukong/pkg/util"
	"wukong/server/model"
)

func (s *Service) OAuthLogin(provider string) string {
	return s.getOAuthConfig().AuthCodeURL(provider)
}

func (s *Service) OAuthCallback(code string, state string) (string, util.RespError) {
	oAuthUser, err := s.getOAuthUser(s.getOAuthConfig(), code, state)
	if err != nil {
		return "", util.NewRespError(err).WithCode(xresponse.HttpUnauthorized)
	}
	if oAuthUser.GetUsername() == "" {
		return "", util.NewRespError(errors.New("username not found"))
	}
	user := model.User{
		Username:    oAuthUser.GetUsername(),
		DisplayName: oAuthUser.GetDisplayName(),
		Email:       oAuthUser.GetEmail(),
		Source:      state,
		IsActive:    1,
	}
	tx := db.DB.Select("id,username,is_active").Where(&model.User{Username: user.Username}).First(&user)
	if tx.RowsAffected == 0 {
		err = db.DB.Create(&user).Error
		if err != nil {
			return "", util.NewRespError(err)
		}
	}
	if user.ID == 0 || user.IsActive != 1 {
		return "", util.NewRespError(errors.New("用户已禁用"), true).WithCode(xresponse.HttpUnauthorized)
	}
	token := model.Token{
		Username:    user.Username,
		AccessLevel: s.GetMaxAccessLevel(s.GetUserRoles(user.ID)),
		TokenType:   model.SessionToken,
	}
	tokenStr, err := s.CreateOrSetToken(&token)
	if err != nil {
		return "", util.NewRespError(err)
	}
	logrus.Infof("OAuth 用户登陆成功: %s", user.Username)
	return tokenStr, nil
}

func (s *Service) getOAuthUser(conf *oauth2.Config, code string, state string) (useroAuth.OAuthProvider, error) {
	token, err := conf.Exchange(context.Background(), code)
	if err != nil {
		return nil, err
	}
	logrus.Debugf("OAuth token: %+v", token)
	client := conf.Client(context.Background(), token)
	resp, err := client.Get(config.Conf.OAuth.UserInfoURL)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != 200 {
		return nil, errors.New(fmt.Sprintf("get user info status code: %d", resp.StatusCode))
	}
	defer func() {
		_ = resp.Body.Close()
	}()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	op, ok := useroAuth.OAuthProviders[state]
	if !ok {
		return nil, errors.New("provider not found")
	}
	return op.GetUserInfo(body)
}

func (s *Service) getOAuthConfig() *oauth2.Config {
	return &oauth2.Config{
		ClientID:     config.Conf.OAuth.ClientID,
		ClientSecret: config.Conf.OAuth.ClientSecret,
		Endpoint: oauth2.Endpoint{
			AuthURL:  config.Conf.OAuth.AuthURL,
			TokenURL: config.Conf.OAuth.TokenURL,
		},
		RedirectURL: config.Conf.OAuth.RedirectURL,
		Scopes:      config.Conf.OAuth.Scopes,
	}
}
