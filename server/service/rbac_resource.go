package service

import (
	"wukong/server/model"
)

func (s *Service) GetResources(query *model.Query) (*model.Resources, *model.Error) {
	var resources model.Resources
	opt := setQuery(query)
	err := dalf().SetQuery(opt).Query(&resources.Resources)
	if err != nil {
		return nil, model.NewError(err.Error())
	}
	resources.Page = *opt.Page
	return &resources, nil
}

func (s *Service) CreateResource(resource *model.Resource) *model.Error {
	if err := dalf().Create(&resource); err != nil {
		return &model.Error{Err: err}
	}
	return nil
}

func (s *Service) UpdateResource(resource *model.Resource) *model.Error {
	if err := dalf().Save(&resource); err != nil {
		return model.NewError(err.Error())
	}
	return nil
}

func (s *Service) DeleteResource(resource *model.Resource) *model.Error {
	if err := dalf().Delete(&resource); err != nil {
		return model.NewError(err.Error())
	}
	return nil
}
