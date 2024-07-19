package service

import (
	"fmt"
	"github.com/go-ldap/ldap/v3"
	"github.com/sanmuyan/xpkg/xutil"
	"github.com/sirupsen/logrus"
	"wukong/pkg/config"
	"wukong/pkg/db"
	"wukong/pkg/util"
	"wukong/server/model"
)

func (s *Service) SyncLDAPUsers() (string, util.RespError) {
	conf := config.Conf.LDAP
	l, err := ldap.DialURL(conf.URL)
	if err != nil {
		return "", util.NewRespError(err)
	}
	err = l.Bind(conf.AdminDN, conf.AdminPassword)
	if err != nil {
		return "", util.NewRespError(err)
	}
	sr := ldap.SearchRequest{
		BaseDN: conf.SearchBase,
		Scope:  ldap.ScopeWholeSubtree,
		Filter: conf.SearchFilter,
		Attributes: []string{
			conf.AttributeMap.DisplayName,
			conf.AttributeMap.Email,
			conf.AttributeMap.Mobile,
			conf.UsernameAttribute,
		},
	}
	sp, err := l.Search(&sr)
	if err != nil {
		return "", util.NewRespError(err)
	}
	totalUserCount := len(sp.Entries)
	var newUserCount int
	for _, entry := range sp.Entries {
		displayName := entry.GetAttributeValue(conf.AttributeMap.DisplayName)
		email := entry.GetAttributeValue(conf.AttributeMap.Email)
		mobile := entry.GetAttributeValue(conf.AttributeMap.Mobile)
		username := entry.GetAttributeValue(conf.UsernameAttribute)
		tx := db.DB.Select("username").Where(&model.User{Username: username}).First(&model.User{})
		if tx.RowsAffected == 0 {
			if username == "" {
				logrus.Warnf("ldap user has no username: dn=%s", entry.DN)
				continue
			}
			newUserCount++
			logrus.Infof("new ldap user: username=%s", username)
			db.DB.Create(&model.User{
				Username:    username,
				DisplayName: displayName,
				Email:       email,
				Mobile:      mobile,
				Source:      "ldap",
				IsActive:    1,
			})
		}
	}
	logrus.Infof("sync ldap users complete: totalUserCount=%d, newUserCount=%d", totalUserCount, newUserCount)
	return fmt.Sprintf("LDAP 用户同步完成：用户总计=%d, 新增用户=%d", totalUserCount, newUserCount), nil
}

func (s *Service) LDAPConnTest(req *config.LDAP) util.RespError {
	l, err := ldap.DialURL(req.URL)
	if err != nil {
		return util.NewRespError(err, true)
	}
	if req.AdminPassword == "" {
		req.AdminPassword = config.Conf.LDAP.AdminPassword
	}
	err = l.Bind(req.AdminDN, req.AdminPassword)
	if err != nil {
		util.NewRespError(err, true)
	}
	return nil
}

func (s *Service) LDAPLoginTest(req *model.LoginRequest) util.RespError {
	l, err := ldap.DialURL(config.Conf.LDAP.URL)
	if err != nil {
		return util.NewRespError(err, true)
	}
	err = l.Bind(fmt.Sprintf("%s=%s,%s", config.Conf.LDAP.UsernameAttribute, req.Username, config.Conf.LDAP.SearchBase), xutil.RemoveError(s.DecryptClientData(req.Password)))
	if err != nil {
		return util.NewRespError(err, true)
	}
	return nil
}
