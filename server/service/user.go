package service

import (
	"errors"
	"github.com/sanmuyan/xpkg/xbcrypt"
	"github.com/sanmuyan/xpkg/xresponse"
	"github.com/sanmuyan/xpkg/xutil"
	"wukong/pkg/config"
	"wukong/pkg/db"
	"wukong/pkg/util"
	"wukong/server/model"
)

func (s *Service) GetUsers(query *model.Query) (*model.Users, util.RespError) {
	var users model.Users
	err := queryData(query, &users)
	if err != nil {
		return nil, util.NewRespError(err)
	}
	for i := range users.Users {
		users.Users[i].Password = ""
	}
	return &users, nil
}

func (s *Service) CreateUser(user *model.User) util.RespError {
	if !xutil.IsUsername(user.Username) {
		return util.NewRespError(errors.New("用户名不符合要求"), true).WithCode(xresponse.HttpBadRequest)
	}
	if !xbcrypt.IsPasswordComplexity(user.Password, config.Conf.Security.PasswordMinLength, config.Conf.Security.PasswordComplexity) {
		return util.NewRespError(errors.New("密码不符合要求"), true).WithCode(xresponse.HttpBadRequest)
	}
	tx := db.DB.Select("id").Where(&model.User{Username: user.Username}).First(&model.User{})
	if tx.RowsAffected > 0 {
		return util.NewRespError(errors.New("用户名已存在"), true).WithCode(xresponse.HttpBadRequest)
	}
	user.Source = "local"
	user.IsActive = 1
	user.Password = xbcrypt.CreatePassword(user.Password)
	if err := db.DB.Create(&user).Error; err != nil {
		return util.NewRespError(err)
	}
	return nil
}

func (s *Service) UpdateUser(user *model.User) util.RespError {
	_user := &model.User{
		ID:          user.ID,
		Email:       user.Email,
		Mobile:      user.Mobile,
		DisplayName: user.DisplayName,
	}
	if err := db.DB.Updates(_user).Error; err != nil {
		return util.NewRespError(err)
	}
	return nil
}

func (s *Service) DeleteUser(user *model.User) util.RespError {
	if err := db.DB.Delete(&model.User{}, user.ID).Error; err != nil {
		return util.NewRespError(err)
	}
	return nil
}
