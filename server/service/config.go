package service

import (
	"encoding/json"
	"github.com/sanmuyan/xpkg/xutil"
	"wukong/pkg/config"
	"wukong/pkg/db"
	"wukong/server/model"
)

func (s *Service) saveConfig() {
	db.DB.Where("name = ?", "config").Updates(&model.Config{
		Name:    "config",
		Content: string(xutil.RemoveError(json.Marshal(config.Conf))),
	})
}

func (s *Service) GetBasicConfig() config.Basic {
	return config.Conf.Basic
}

func (s *Service) UpdateBasicConfig(req *config.Basic) {
	_ = xutil.FillObj(req, &config.Conf.Basic)
	s.saveConfig()
}

func (s *Service) GetSecurityConfig() *config.Security {
	return &config.Conf.Security
}

func (s *Service) UpdateSecurityConfig(req *config.Security) {
	_ = xutil.FillObj(req, &config.Conf.Security)
	s.saveConfig()
}

func (s *Service) GetLDAPConfig() *config.LDAP {
	resp := config.Conf.LDAP
	resp.AdminPassword = ""
	return &resp
}

func (s *Service) UpdateLDAPConfig(req *config.LDAP) {
	_ = xutil.FillObj(req, &config.Conf.LDAP)
	s.saveConfig()
}

func (s *Service) GetOauthProvidersConfig() *[]config.OauthProvider {
	return &config.Conf.OauthProviders
}

func (s *Service) UpdateOauthProvidersConfig(req *[]config.OauthProvider) {
	_ = xutil.FillObj(req, &config.Conf.OauthProviders)
	s.saveConfig()
}
