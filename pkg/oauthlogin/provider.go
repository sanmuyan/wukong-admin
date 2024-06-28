package oauthlogin

import (
	"sync"
	"wukong/pkg/config"
)

type OauthUserProvider interface {
	GetUserInfo(user []byte) (OauthUserProvider, error)
	GetUsername() string
	GetEmail() string
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
