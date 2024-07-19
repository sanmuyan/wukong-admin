package oauthlogin

import (
	"encoding/json"
	"errors"
	"fmt"
	"sync"
	"time"
	"wukong/pkg/config"
	"wukong/server/model"
)

// 1. 用户登录获取 code
// 2. 获取应用 access_token 并缓存
// 3. 通过 code 和 access_token 换取用户 user_access_token
// 4. 通过 user_access_token 获取用户信息

var WecomAccessTokenStore = new(sync.Map)

type WecomError struct {
	ErrCode int    `json:"errcode"`
	ErrMsg  string `json:"errmsg"`
}

type WecomAccessToken struct {
	WecomError
	AccessToken string    `json:"access_token"`
	ExpiresIn   int       `json:"expires_in"`
	ExpiresAt   time.Time `json:"expires_at"`
}

type WecomUser struct {
	WecomError
	// 用户在该企业的唯一标识
	UserID string `json:"userid"`
}

func NewWecomUser() *WecomUser {
	return &WecomUser{}
}

func (c *WecomUser) GetAuthURL(conf config.OauthProvider, state string) string {
	url := fmt.Sprintf("%s?login_type=CorpApp&appid=%s&agentid=%s&redirect_uri=%s", conf.AuthURL, conf.CorpID, conf.ClientID, conf.RedirectURL)
	if state != "" {
		url = fmt.Sprintf("%s&state=%s", url, state)
	}
	return url
}

func (c *WecomUser) GetUserIDField() string {
	return model.UserWecomIDField
}

func (c *WecomUser) GetUserID() string {
	return c.UserID
}

func (c *WecomUser) GetUser(conf config.OauthProvider, code string) (OauthUserProvider, error) {
	at, err := getWecomAccessToken(conf)
	if err != nil {
		return nil, err
	}
	userRaw, err := getWecomUser(conf, code, at.AccessToken)
	if err != nil {
		return nil, err
	}
	user := &WecomUser{}
	err = json.Unmarshal(userRaw, user)
	if err != nil {
		return nil, err
	}
	if user.ErrCode != 0 {
		return nil, errors.New(string(userRaw))
	}
	return user, nil
}
