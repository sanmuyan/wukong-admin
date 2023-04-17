package service

import (
	"wukong/server/model"
)

func (s *Service) GetUserBinds(query *model.Query) (*model.UserBinds, *model.Error) {
	var userBinds model.UserBinds
	opt := setQuery(query)
	err := dalf().SetQuery(opt).Query(&userBinds.UserBinds)
	if err != nil {
		return nil, model.NewError(err.Error())
	}
	userBinds.Page = *opt.Page
	return &userBinds, nil
}

func (s *Service) CreateUserBind(userBind *model.UserBind) *model.Error {
	_ = dalf().Get(&userBind)
	if userBind.Id != 0 {
		return model.NewError("已存在该用户绑定", true)
	}
	if err := dalf().Create(&userBind); err != nil {
		return model.NewError(err.Error())
	}
	return nil
}

func (s *Service) DeleteUserBind(userBind *model.UserBind) *model.Error {
	if err := dalf().Delete(&userBind); err != nil {
		return model.NewError(err.Error())
	}
	return nil
}
