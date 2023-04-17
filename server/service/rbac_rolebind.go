package service

import (
	"wukong/server/model"
)

func (s *Service) GetRoleBinds(query *model.Query) (*model.RoleBinds, *model.Error) {
	var roleBinds model.RoleBinds
	opt := setQuery(query)
	err := dalf().SetQuery(opt).Query(&roleBinds.RoleBinds)
	if err != nil {
		return nil, &model.Error{Err: err}
	}
	roleBinds.Page = *opt.Page
	return &roleBinds, nil
}

func (s *Service) CreateRoleBind(roleBind *model.RoleBind) *model.Error {
	_ = dalf().Get(&roleBind)
	if roleBind.Id != 0 {
		return model.NewError("已存在该角色绑定", true)
	}
	if err := dalf().Create(&roleBind); err != nil {
		return model.NewError(err.Error())
	}
	return nil
}

func (s *Service) DeleteRoleBind(roleBind *model.RoleBind) *model.Error {
	if err := dalf().Delete(&roleBind); err != nil {
		return model.NewError(err.Error())
	}
	return nil
}
