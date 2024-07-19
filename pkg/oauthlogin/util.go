package oauthlogin

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/sanmuyan/xpkg/xrequest"
	"github.com/sanmuyan/xpkg/xutil"
	"github.com/sirupsen/logrus"
	"golang.org/x/oauth2"
	"io"
	"strings"
	"time"
	"wukong/pkg/config"
)

func getUser(providerConf config.OauthProvider, code string) ([]byte, error) {
	conf := getOauthConfig(providerConf)
	token, err := conf.Exchange(context.Background(), code)
	if err != nil {
		return nil, err
	}
	logrus.Debugf("Oauth token: %+v", token)
	client := conf.Client(context.Background(), token)
	resp, err := client.Get(providerConf.UserInfoURL)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != 200 {
		return nil, errors.New(fmt.Sprintf("get user info status code: %d", resp.StatusCode))
	}
	defer func() {
		_ = resp.Body.Close()
	}()
	return io.ReadAll(resp.Body)
}

func getOauthConfig(providerConf config.OauthProvider) *oauth2.Config {
	return &oauth2.Config{
		ClientID:     providerConf.ClientID,
		ClientSecret: providerConf.ClientSecret,
		Endpoint: oauth2.Endpoint{
			AuthURL:  providerConf.AuthURL,
			TokenURL: providerConf.TokenURL,
		},
		RedirectURL: providerConf.RedirectURL,
		Scopes:      strings.Split(providerConf.Scopes, ","),
	}
}

func getWecomAccessToken(conf config.OauthProvider) (*WecomAccessToken, error) {
	dsKey := conf.CorpID + conf.ClientSecret
	getAt := func() (*WecomAccessToken, error) {
		client := xrequest.Request{
			Config: &xrequest.Options{
				Method: "GET",
				URL:    fmt.Sprintf("%s?corpid=%s&corpsecret=%s", conf.TokenURL, conf.CorpID, conf.ClientSecret),
			},
		}
		req, err := client.Request()
		if err != nil {
			return nil, err
		}
		var accessToken WecomAccessToken
		err = json.Unmarshal(req.Body, &accessToken)
		if err != nil {
			return nil, err
		}
		if accessToken.ErrCode != 0 {
			return nil, errors.New(string(req.Body))
		}
		accessToken.ExpiresAt = time.Now().UTC().Add(time.Duration(accessToken.ExpiresIn-10) * time.Second)
		WecomAccessTokenStore.Store(dsKey, &accessToken)
		return &accessToken, nil
	}
	_at, ok := WecomAccessTokenStore.Load(dsKey)
	if !ok {
		return getAt()
	}
	at := _at.(*WecomAccessToken)
	if at.ExpiresAt.Before(time.Now().UTC()) {
		return getAt()
	}
	return at, nil
}

func getWecomUser(conf config.OauthProvider, code, at string) ([]byte, error) {
	client := xrequest.Request{
		Config: &xrequest.Options{
			Method: "GET",
			URL:    fmt.Sprintf("%s?code=%s&access_token=%s", conf.UserInfoURL, code, at),
		},
	}
	req, err := client.Request()
	if err != nil {
		return nil, err
	}
	return req.Body, nil
}

func getDingtalkAccessToken(conf config.OauthProvider, code string) (*DingtalkUserAccessToken, error) {
	var ak *DingtalkAccessToken
	var err error
	dsKey := conf.ClientID + conf.ClientSecret
	getAt := func() (*DingtalkAccessToken, error) {
		type body struct {
			AppKey    string `json:"appKey"`
			AppSecret string `json:"appSecret"`
		}
		client := xrequest.Request{
			Config: &xrequest.Options{
				Method: "POST",
				URL:    conf.TokenURL + "/accessToken",
				Head:   map[string]string{"Content-Type": "application/json"},
				Body: xutil.RemoveError(json.Marshal(body{
					AppKey:    conf.ClientID,
					AppSecret: conf.ClientSecret,
				})),
			},
		}
		req, err := client.Request()
		if err != nil {
			return nil, err
		}
		var accessToken DingtalkAccessToken
		err = json.Unmarshal(req.Body, &accessToken)
		if err != nil {
			return nil, err
		}
		if accessToken.DingtalkError != nil {
			return nil, errors.New(string(req.Body))
		}
		accessToken.ExpiresAt = time.Now().UTC().Add(time.Duration(accessToken.ExpiresIn-10) * time.Second)
		WecomAccessTokenStore.Store(dsKey, &accessToken)
		return &accessToken, nil
	}
	_at, ok := DingtalkAccessTokenStore.Load(dsKey)
	if ok {
		ak = _at.(*DingtalkAccessToken)
		if ak.ExpiresAt.Before(time.Now().UTC()) {
			ak, err = getAt()
			if err != nil {
				return nil, err
			}
		}
	} else {
		ak, err = getAt()
		if err != nil {
			return nil, err
		}
	}
	type body struct {
		ClientId     string `json:"clientId"`
		ClientSecret string `json:"clientSecret"`
		Code         string `json:"code"`
		GrantType    string `json:"grantType"`
	}
	client := xrequest.Request{
		Config: &xrequest.Options{
			Method: "POST",
			Head:   map[string]string{"Content-Type": "application/json", "x-acs-dingtalk-access-token": ak.AccessToken},
			URL:    conf.TokenURL + "/userAccessToken",
			Body: xutil.RemoveError(json.Marshal(body{
				ClientId:     conf.ClientID,
				ClientSecret: conf.ClientSecret,
				Code:         code,
				GrantType:    "authorization_code",
			})),
		}}
	req, err := client.Request()
	if err != nil {
		return nil, err
	}
	var uak DingtalkUserAccessToken
	err = json.Unmarshal(req.Body, &uak)
	if err != nil {
		return nil, err
	}
	if uak.DingtalkUserError != nil {
		return nil, errors.New(string(req.Body))
	}
	return &uak, nil
}

func getDingtalkUser(conf config.OauthProvider, uat string) ([]byte, error) {
	client := xrequest.Request{
		Config: &xrequest.Options{
			Method: "GET",
			URL:    fmt.Sprintf("%s", conf.UserInfoURL),
			Head:   map[string]string{"x-acs-dingtalk-access-token": uat, "Content-Type": "application/json"},
		},
	}
	req, err := client.Request()
	if err != nil {
		return nil, err
	}
	return req.Body, nil
}
