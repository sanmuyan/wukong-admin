package service

import (
	"time"
	"wukong/pkg/util"
	"wukong/server/model"
)

func (s *Service) GetUsers(query *model.Query) (*model.Users, *model.Error) {
	var users model.Users
	opt := setQuery(query)
	err := dalf().SetQuery(opt).Query(&users.Users)
	if err != nil {
		return nil, &model.Error{Err: err}
	}
	users.Page = *opt.Page
	for i := range users.Users {
		users.Users[i].Password = ""
	}
	return &users, nil
}

func (s *Service) CreateUser(user *model.User) *model.Error {
	nowTime := time.Now().Format("2006-01-02 15:04:05")
	if !util.IsUserName(user.Username) {
		return model.NewError("用户名不符合要求", true)
	}
	if !util.IsPassword(user.Password) {
		return model.NewError("密码不符合要求", true)
	}
	_ = dalf().Get(&model.User{Username: user.Username})
	if user.Id != 0 {
		return model.NewError("用户名已存在", true)
	}
	user.Source = "local"
	user.IsActive = 1
	user.UpdateTime = nowTime
	user.CreateTime = nowTime
	user.Password = util.CreatePassword(user.Password)
	if err := dalf().Create(&user); err != nil {
		return model.NewError(err.Error())
	}
	return nil
}

func (s *Service) UpdateUser(user *model.User) *model.Error {
	nowTime := time.Now().Format("2006-01-02 15:04:05")
	user.UpdateTime = nowTime
	if len(user.Password) > 0 {
		if !util.IsPassword(user.Password) {
			return model.NewError("密码不符合要求", true)
		}
		user.Password = util.CreatePassword(user.Password)
	}
	if err := dalf().Update(&user); err != nil {
		return model.NewError(err.Error())
	}
	return nil
}

func (s *Service) DeleteUser(user *model.User) *model.Error {
	if err := dalf().Delete(&user); err != nil {
		return model.NewError(err.Error())
	}
	return nil
}
