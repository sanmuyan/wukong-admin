package oauthlogin

import (
	"sync"
	"wukong/pkg/config"
)

// OauthUserProvider 第三方 OAuth 登录获取用户信息
type OauthUserProvider interface {
	// GetUserInfo 获取用户信息
	GetUserInfo(user []byte) (OauthUserProvider, error)
	// GetUsername 获取用户名
	GetUsername() string
	// GetEmail 获取邮箱
	GetEmail() string
	// GetDisplayName 获取显示名称
	GetDisplayName() string
}

var OauthUserProviders = make(map[string]OauthUserProvider)

func init() {
	OauthUserProviders["gitlab"] = NewGitlabUser()
}

var OauthProviders *sync.Map

func InitOauthProviders() {
	OauthProviders = new(sync.Map)
	for _, oauthProvider := range config.Conf.OauthProviders {
		_, ok := OauthUserProviders[oauthProvider.Provider]
		if oauthProvider.Enable && ok {
			OauthProviders.Store(oauthProvider.Provider, oauthProvider)
		}
	}
}
