package service

import (
	"wukong/pkg/db"
	"wukong/pkg/util"
	"wukong/server/model"
)

func (s *Service) GetResources(query *model.Query) (*model.Resources, util.RespError) {
	var resources model.Resources
	err := queryData(query, &resources)
	if err != nil {
		return nil, util.NewRespError(err)
	}
	return &resources, nil
}

func (s *Service) CreateResource(resource *model.Resource) util.RespError {
	if err := db.DB.Create(&resource).Error; err != nil {
		return util.NewRespError(err)
	}
	return nil
}

func (s *Service) UpdateResource(resource *model.Resource) util.RespError {
	if err := db.DB.Updates(&resource).Error; err != nil {
		return util.NewRespError(err)
	}
	return nil
}

func (s *Service) DeleteResource(resource *model.Resource) util.RespError {
	if err := db.DB.Delete(&model.Resource{}, resource.ID).Error; err != nil {
		return util.NewRespError(err)
	}
	return nil
}
