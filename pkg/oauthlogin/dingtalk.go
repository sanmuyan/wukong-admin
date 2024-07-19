package oauthlogin

import (
	"encoding/json"
	"errors"
	"sync"
	"time"
	"wukong/pkg/config"
	"wukong/server/model"
)

// 1. 用户登录获取 code
// 2. 获取应用 access_token 并缓存
// 3. 通过 code 和 access_token 直接获取用户信息

var DingtalkAccessTokenStore = new(sync.Map)

type DingtalkError struct {
	ErrCode string `json:"errcode"`
	ErrMsg  string `json:"errmsg"`
}

type DingtalkUserError struct {
	ErrCode string `json:"code"`
	ErrMsg  string `json:"message"`
}

type DingtalkUserAccessToken struct {
	AccessToken string `json:"accessToken"`
	*DingtalkUserError
}

type DingtalkAccessToken struct {
	*DingtalkError
	AccessToken string    `json:"access_token"`
	ExpiresIn   int       `json:"expires_in"`
	ExpiresAt   time.Time `json:"expires_at"`
}

type DingtalkUser struct {
	*DingtalkUserError
	// UnionId 用户在该企业的唯一标识
	UnionID string `json:"unionId"`
}

func NewDingtalkUser() *DingtalkUser {
	return &DingtalkUser{}
}

func (c *DingtalkUser) GetAuthURL(conf config.OauthProvider, state string) string {
	return getOauthConfig(conf).AuthCodeURL(state) + "&prompt=consent"
}

func (c *DingtalkUser) GetUserIDField() string {
	return model.UserDingtalkIDField
}

func (c *DingtalkUser) GetUserID() string {
	return c.UnionID
}

func (c *DingtalkUser) GetUser(conf config.OauthProvider, code string) (OauthUserProvider, error) {
	uat, err := getDingtalkAccessToken(conf, code)
	if err != nil {
		return nil, err
	}
	userRaw, err := getDingtalkUser(conf, uat.AccessToken)
	if err != nil {
		return nil, err
	}
	user := &DingtalkUser{}
	err = json.Unmarshal(userRaw, user)
	if err != nil {
		return nil, err
	}
	if user.DingtalkUserError != nil {
		return nil, errors.New(string(userRaw))
	}
	return user, nil
}
