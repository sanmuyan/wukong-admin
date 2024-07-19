package oauthlogin

import (
	"sync"
	"wukong/pkg/config"
	"wukong/server/model"
)

// OauthUserProvider 解析用户信息
type OauthUserProvider interface {
	// GetAuthURL 获取授权地址
	GetAuthURL(config.OauthProvider, string) string
	// GetUser 获取用户信息
	GetUser(config.OauthProvider, string) (OauthUserProvider, error)
	// GetUserIDField 获取用户关联字段
	GetUserIDField() string
	// GetUserID 用户在第三方的唯一ID
	GetUserID() string
}

var OauthUserProviders = make(map[string]OauthUserProvider)

func init() {
	OauthUserProviders[model.OauthProviderGitLab] = NewGitlabUser()
	OauthUserProviders[model.OauthProviderWecom] = NewWecomUser()
	OauthUserProviders[model.OauthProviderDingtalk] = NewDingtalkUser()
}

var OauthProviderConfig *sync.Map

func InitOauthProviderConfig() {
	OauthProviderConfig = new(sync.Map)
	OauthProviderConfig.Store(model.OauthProviderGitLab, config.Conf.OauthProviders.Gitlab)
	OauthProviderConfig.Store(model.OauthProviderWecom, config.Conf.OauthProviders.Wecom)
	OauthProviderConfig.Store(model.OauthProviderDingtalk, config.Conf.OauthProviders.Dingtalk)
}
