package service

import (
	"wukong/pkg/db"
	"wukong/pkg/util"
	"wukong/server/model"
)

func (s *Service) GetRoles(query *model.Query) (*model.Roles, util.RespError) {
	var roles model.Roles
	err := queryData(query, &roles)
	if err != nil {
		return nil, util.NewRespError(err)
	}
	return &roles, nil
}

func (s *Service) CreateRole(role *model.Role) util.RespError {
	if err := db.DB.Create(&role).Error; err != nil {
		return util.NewRespError(err)
	}
	return nil
}

func (s *Service) UpdateRole(role *model.Role) util.RespError {
	if err := db.DB.Updates(&role).Error; err != nil {
		return util.NewRespError(err)
	}
	return nil
}

func (s *Service) DeleteRole(role *model.Role) util.RespError {
	if err := db.DB.Delete(&model.Role{}, role.ID).Error; err != nil {
		return util.NewRespError(err)
	}
	return nil
}
