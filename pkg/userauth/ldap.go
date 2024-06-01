package userauth

import (
	"fmt"
	"github.com/go-ldap/ldap/v3"
	"github.com/sirupsen/logrus"
)

type LDAPAuth struct {
	url string
	sd  string
	cn  string
}

func NewLDAPAuth(url string, searchBase string, cnAttribute string) *LDAPAuth {
	return &LDAPAuth{
		url: url,
		sd:  searchBase,
		cn:  cnAttribute,
	}
}

func (c *LDAPAuth) Login(username string, password string) bool {
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
