package service

import (
	model "wukong/server/model"
)

func (s *Service) GetRoles(query *model.Query) (*model.Roles, *model.Error) {
	var roles model.Roles
	var opt options
	opt.setQuery(query)
	dal := newDal(&opt)
	err := dal.Query(&roles.Roles)
	if err != nil {
		return nil, model.NewError(err.Error())
	}
	roles.Page = opt.Page
	return &roles, nil
}

func (s *Service) CreateRole(role *model.Role) *model.Error {
	dal := newDal(&options{})
	if err := dal.Create(&role); err != nil {
		return model.NewError(err.Error())
	}
	return nil
}

func (s *Service) UpdateRole(role *model.Role) *model.Error {
	dal := newDal(&options{})
	if err := dal.Update(&role); err != nil {
		return &model.Error{Err: err}
	}
	return nil
}

func (s *Service) DeleteRole(role *model.Role) *model.Error {
	dal := newDal(&options{})
	if err := dal.Delete(&role); err != nil {
		return &model.Error{Err: err}
	}
	return nil
}
