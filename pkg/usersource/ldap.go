package usersource

import (
	"fmt"
	"github.com/go-ldap/ldap/v3"
	"github.com/sirupsen/logrus"
)

type LDAPUser struct {
	url string
	sd  string
	cn  string
}

func NewLDAPUser(url string, searchBase string, cnAttribute string) *LDAPUser {
	return &LDAPUser{
		url: url,
		sd:  searchBase,
		cn:  cnAttribute,
	}
}

func (c *LDAPUser) Login(username string, password string) bool {
	l, err := ldap.DialURL(c.url)
	if err != nil {
		logrus.Errorf("failed to dial: %s", err)
		return false
	}
	err = l.Bind(fmt.Sprintf("%s=%s,%s", c.cn, username, c.sd), password)
	if err != nil {
		return false
	}
	return true
}
