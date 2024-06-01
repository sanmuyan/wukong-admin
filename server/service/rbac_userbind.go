package service

import (
	"wukong/pkg/db"
	"wukong/pkg/util"
	"wukong/server/model"
)

func (s *Service) GetUserBinds(query *model.Query) (*model.UserBinds, util.RespError) {
	var userBinds model.UserBinds
	err := queryData(query, &userBinds)
	if err != nil {
		return nil, util.NewRespError(err)
	}
	return &userBinds, nil
}

func (s *Service) CreateUserBind(userBind *model.UserBind) util.RespError {
	if err := db.DB.Create(&userBind).Error; err != nil {
		return util.NewRespError(err)
	}
	return nil
}

func (s *Service) DeleteUserBind(userBind *model.UserBind) util.RespError {
	if err := db.DB.Delete(&model.UserBind{}, userBind.ID).Error; err != nil {
		return util.NewRespError(err)
	}
	return nil
}
