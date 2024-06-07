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
	useroauth "wukong/pkg/oauthlogin"
	"wukong/pkg/util"
	"wukong/server/model"
)

func (s *Service) OauthLogin(provider string) (string, util.RespError) {
	providerConf, ok := config.OauthProviders[provider]
	if !ok {
		return "", util.NewRespError(errors.New("provider not found"))
	}
	return s.getOauthConfig(providerConf).AuthCodeURL(provider), nil
}

func (s *Service) OauthCallback(code string, state string) (string, util.RespError) {
	providerConf, ok := config.OauthProviders[state]
	if !ok {
		return "", util.NewRespError(errors.New("provider not found"))
	}
	oauthUser, err := s.getOauthUser(s.getOauthConfig(providerConf), code, state, providerConf.UserInfoURL)
	if err != nil {
		return "", util.NewRespError(err).WithCode(xresponse.HttpUnauthorized)
	}
	if oauthUser.GetUsername() == "" {
		return "", util.NewRespError(errors.New("username not found"))
	}
	user := model.User{
		Username:    oauthUser.GetUsername(),
		DisplayName: oauthUser.GetDisplayName(),
		Email:       oauthUser.GetEmail(),
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
	tokenStr, err := s.CreateOrSetToken(&token, token.Username, config.Conf.TokenTTL)
	if err != nil {
		return "", util.NewRespError(err)
	}
	logrus.Infof("Oauth 用户登陆成功: %s", user.Username)
	return tokenStr, nil
}

func (s *Service) getOauthUser(conf *oauth2.Config, code string, state string, userInfoURL string) (useroauth.OauthProvider, error) {
	token, err := conf.Exchange(context.Background(), code)
	if err != nil {
		return nil, err
	}
	logrus.Debugf("Oauth token: %+v", token)
	client := conf.Client(context.Background(), token)
	resp, err := client.Get(userInfoURL)
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
	op, ok := useroauth.OauthProviders[state]
	if !ok {
		return nil, errors.New("provider not found")
	}
	return op.GetUserInfo(body)
}

func (s *Service) getOauthConfig(providerConf config.OauthProvider) *oauth2.Config {
	return &oauth2.Config{
		ClientID:     providerConf.ClientID,
		ClientSecret: providerConf.ClientSecret,
		Endpoint: oauth2.Endpoint{
			AuthURL:  providerConf.AuthURL,
			TokenURL: providerConf.TokenURL,
		},
		RedirectURL: providerConf.RedirectURL,
		Scopes:      providerConf.Scopes,
	}
}