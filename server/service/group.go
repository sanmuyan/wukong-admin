package service

import (
	"wukong/pkg/db"
	"wukong/pkg/util"
	"wukong/server/model"
)

func (s *Service) GetGroups(query *model.Query) (*model.Groups, util.RespError) {
	var groups model.Groups
	err := queryData(query, &groups)
	if err != nil {
		return nil, util.NewRespError(err)
	}
	return &groups, nil
}

func (s *Service) CreateGroup(group *model.Group) util.RespError {
	if err := db.DB.Create(&group).Error; err != nil {
		return util.NewRespError(err)
	}
	return nil
}

func (s *Service) UpdateGroup(group *model.Group) util.RespError {
	if err := db.DB.Updates(group).Error; err != nil {
		return util.NewRespError(err)
	}
	return nil
}

func (s *Service) DeleteGroup(group *model.Group) util.RespError {
	if err := db.DB.Delete(&model.Group{}, group.ID).Error; err != nil {
		return util.NewRespError(err)
	}
	return nil
}

func (s *Service) GetUserGroupBinds(query *model.Query) (*model.UserGroupBinds, util.RespError) {
	var userGroupBinds model.UserGroupBinds
	err := queryData(query, &userGroupBinds)
	if err != nil {
		return nil, util.NewRespError(err)
	}
	return &userGroupBinds, nil
}

func (s *Service) CreateUserGroupBind(userGroupBind *model.UserGroupBind) util.RespError {
	if err := db.DB.Create(&userGroupBind).Error; err != nil {
		return util.NewRespError(err)
	}
	return nil
}

func (s *Service) DeleteUserGroupBind(userGroupBind *model.UserGroupBind) util.RespError {
	if err := db.DB.Delete(&model.UserGroupBind{}, userGroupBind.ID).Error; err != nil {
		return util.NewRespError(err)
	}
	return nil
}
