package service

import (
	"errors"
	"github.com/sanmuyan/xpkg/xresponse"
	"github.com/sirupsen/logrus"
	"wukong/pkg/db"
	"wukong/pkg/userauth"
	"wukong/pkg/util"
	"wukong/server/model"
)

func (s *Service) Login(login model.Login) (string, util.RespError) {
	var user model.User
	user.Username = login.Username
	if err := db.DB.Where(&model.User{Username: user.Username}).First(&user).Error; err != nil {
		return "", util.NewRespError(err)
	}

	if user.ID == 0 || user.IsActive != 1 {
		return "", util.NewRespError(errors.New("用户名或密码错误"), true).WithCode(xresponse.HttpUnauthorized)
	}

	as, ok := userauth.AuthSources[user.Source]
	if !ok {
		return "", util.NewRespError(errors.New("未知的用户来源"), false)
	}
	if !as.Login(login.Username, login.Password) {
		return "", util.NewRespError(errors.New("用户名或密码错误"), true).WithCode(xresponse.HttpUnauthorized)
	}

	token := model.Token{
		Username:    user.Username,
		AccessLevel: s.GetMaxAccessLevel(s.GetUserRoles(user.ID)),
		TokenType:   model.SessionToken,
	}
	tokenStr, err := s.CreateOrSetToken(&token)
	if err != nil {
		return "", util.NewRespError(err)
	}
	logrus.Infof("用户登陆成功: %s", user.Username)
	return tokenStr, nil
}

func (s *Service) Logout(token *model.Token) util.RespError {
	if err := s.DeleteToken(token); err != nil {
		return util.NewRespError(err)
	}
	return nil
}
