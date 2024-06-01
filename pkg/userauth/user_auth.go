package userauth

import "wukong/pkg/config"

type AuthSource interface {
	Login(username string, password string) bool
}

var AuthSources map[string]AuthSource

func InitUserAuth() {
	AuthSources = make(map[string]AuthSource)
	AuthSources["local"] = NewLocalAuth()
	AuthSources["ldap"] = NewLDAPAuth(config.Conf.LDAP.URL, config.Conf.LDAP.SearchBase, config.Conf.LDAP.UsernameAttribute)
}
