package usersource

import (
	"sync"
	"wukong/pkg/config"
)

// UserSource 用户来源
type UserSource interface {
	// Login 用户登录
	Login(username string, password string) bool
}

var UserSources *sync.Map

func InitUserSource() {
	UserSources = new(sync.Map)
	UserSources.Store("local", NewLocalUser())
	if config.Conf.LDAP.Enable {
		UserSources.Store("ldap", NewLDAPUser(config.Conf.LDAP.URL, config.Conf.LDAP.SearchBase, config.Conf.LDAP.UsernameAttribute))
	}
}
