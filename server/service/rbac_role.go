package service

import (
	model "wukong/server/model"
)

func (s *Service) GetRoles(query *model.Query) (*model.Roles, *model.Error) {
	var roles model.Roles
	opt := setQuery(query)
	err := dalf().SetQuery(opt).Query(&roles.Roles)
	if err != nil {
		return nil, model.NewError(err.Error())
	}
	roles.Page = *opt.Page
	return &roles, nil
}

func (s *Service) CreateRole(role *model.Role) *model.Error {
	if err := dalf().Create(&role); err != nil {
		return model.NewError(err.Error())
	}
	return nil
}

func (s *Service) UpdateRole(role *model.Role) *model.Error {
	if err := dalf().Update(&role); err != nil {
		return &model.Error{Err: err}
	}
	return nil
}

func (s *Service) DeleteRole(role *model.Role) *model.Error {
	if err := dalf().Delete(&role); err != nil {
		return &model.Error{Err: err}
	}
	return nil
}
