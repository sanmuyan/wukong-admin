package service

import (
	"wukong/pkg/db"
	"wukong/pkg/util"
	"wukong/server/model"
)

func (s *Service) GetOauthAPPS(query *model.Query) (*model.OauthAPPS, util.RespError) {
	var oauthAPPS model.OauthAPPS
	err := queryData(query, &oauthAPPS)
	if err != nil {
		return nil, util.NewRespError(err)
	}
	return &oauthAPPS, nil
}

func (s *Service) CreateOauthAPP(oauthAPP *model.OauthAPP) util.RespError {
	if err := db.DB.Create(&oauthAPP).Error; err != nil {
		return util.NewRespError(err)
	}
	return nil
}

func (s *Service) UpdateOauthAPP(oauthAPP *model.OauthAPP) util.RespError {
	if err := db.DB.Updates(&oauthAPP).Error; err != nil {
		return util.NewRespError(err)
	}
	return nil
}

func (s *Service) DeleteOauthAPP(oauthAPP *model.OauthAPP) util.RespError {
	if err := db.DB.Delete(&model.OauthAPP{}, oauthAPP.ID).Error; err != nil {
		return util.NewRespError(err)
	}
	return nil
}
