package usersource

import "wukong/pkg/config"

type UserSource interface {
	Login(username string, password string) bool
}

var UserSources map[string]UserSource

func InitUserSource() {
	UserSources = make(map[string]UserSource)
	UserSources["local"] = NewLocalUser()
	UserSources["ldap"] = NewLDAPUser(config.Conf.LDAP.URL, config.Conf.LDAP.SearchBase, config.Conf.LDAP.UsernameAttribute)
}
