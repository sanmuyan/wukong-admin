package service

import (
	"wukong/server/model"
)

func (s *Service) GetResources(query *model.Query) (*model.Resources, *model.Error) {
	var resources model.Resources
	var opt options
	opt.setQuery(query)
	dal := newDal(&opt)
	err := dal.Query(&resources.Resources)
	if err != nil {
		return nil, model.NewError(err.Error())
	}
	resources.Page = opt.Page
	return &resources, nil
}

func (s *Service) CreateResource(resource *model.Resource) *model.Error {
	dal := newDal(&options{})
	if err := dal.Create(&resource); err != nil {
		return &model.Error{Err: err}
	}
	return nil
}

func (s *Service) UpdateResource(resource *model.Resource) *model.Error {
	dal := newDal(&options{})
	if err := dal.Update(&resource); err != nil {
		return model.NewError(err.Error())
	}
	return nil
}

func (s *Service) DeleteResource(resource *model.Resource) *model.Error {
	dal := newDal(&options{})
	if err := dal.Delete(&resource); err != nil {
		return model.NewError(err.Error())
	}
	return nil
}
