package service

import (
	"github.com/sanmuyan/xpkg/xcrypto"
	"wukong/pkg/db"
	"wukong/pkg/security"
	"wukong/pkg/util"
	"wukong/server/model"
)

func (s *Service) GetOauthApps(query *model.Query) (*model.OauthApps, util.RespError) {
	var oauthApps model.OauthApps
	err := queryData(query, &oauthApps)
	if err != nil {
		return nil, util.NewRespError(err)
	}
	return &oauthApps, nil
}

func (s *Service) CreateOauthApp(oauthApp *model.OauthApp) util.RespError {
	oauthApp.ClientID = security.GetRandomID()
	oauthApp.ClientSecret = xcrypto.GenerateRandomString(64, true, true, true, false)
	if err := db.DB.Create(&oauthApp).Error; err != nil {
		return util.NewRespError(err)
	}
	return nil
}

func (s *Service) UpdateOauthApp(oauthApp *model.OauthApp) util.RespError {
	if err := db.DB.Updates(&oauthApp).Error; err != nil {
		return util.NewRespError(err)
	}
	return nil
}

func (s *Service) DeleteOauthApp(oauthApp *model.OauthApp) util.RespError {
	if err := db.DB.Delete(&model.OauthApp{}, oauthApp.ID).Error; err != nil {
		return util.NewRespError(err)
	}
	return nil
}
