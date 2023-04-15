package service

import (
	"wukong/server/model"
)

func (s *Service) GetUserBinds(query *model.Query) (*model.UserBinds, *model.Error) {
	var userBinds model.UserBinds
	var opt options
	opt.setQuery(query)
	db := newDal(&opt)
	err := db.Query(&userBinds.UserBinds)
	if err != nil {
		return nil, model.NewError(err.Error())
	}
	userBinds.Page = opt.Page
	return &userBinds, nil
}

func (s *Service) CreateUserBind(userBind *model.UserBind) *model.Error {
	dal := newDal(&options{})
	_ = dal.Get(&userBind)
	if userBind.Id != 0 {
		return model.NewError("已存在该用户绑定", true)
	}
	if err := dal.Create(&userBind); err != nil {
		return model.NewError(err.Error())
	}
	return nil
}

func (s *Service) DeleteUserBind(userBind *model.UserBind) *model.Error {
	dal := newDal(&options{})
	if err := dal.Delete(&userBind); err != nil {
		return model.NewError(err.Error())
	}
	return nil
}
