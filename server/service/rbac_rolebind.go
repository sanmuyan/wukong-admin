package service

import (
	"wukong/pkg/db"
	"wukong/pkg/util"
	"wukong/server/model"
)

func (s *Service) GetRoleBinds(query *model.Query) (*model.RoleBinds, util.RespError) {
	var roleBinds model.RoleBinds
	err := queryData(query, &roleBinds)
	if err != nil {
		return nil, util.NewRespError(err)
	}
	return &roleBinds, nil
}

func (s *Service) CreateRoleBind(roleBind *model.RoleBind) util.RespError {
	if err := db.DB.Create(&roleBind).Error; err != nil {
		return util.NewRespError(err)
	}
	return nil
}

func (s *Service) DeleteRoleBind(roleBind *model.RoleBind) util.RespError {
	if err := db.DB.Delete(&model.RoleBind{}, roleBind.ID).Error; err != nil {
		return util.NewRespError(err)
	}
	return nil
}
