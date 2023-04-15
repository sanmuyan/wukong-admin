package service

import (
	"github.com/sirupsen/logrus"
	"wukong/pkg/config"
	"wukong/pkg/db"
	"wukong/pkg/util"
	"wukong/server/model"
)

func (s *Service) Login(login model.Login) (string, *model.Error) {
	var user model.User
	dal := newDal(&options{})
	user.Username = login.Username
	if err := dal.Get(&user); err != nil {
		return "", model.NewError(err.Error())
	}
	if user.Id == 0 || !util.ComparePassword(user.Password, login.Password) || user.IsActive == 0 {
		return "", model.NewError("用户名或密码错误", true)
	}

	token := model.Token{
		UserId:      user.Id,
		Username:    user.Username,
		AccessLevel: GetMaxAccessLevel(user.Id),
		TokenType:   "session",
		TTL:         config.Conf.TokenTTL,
	}
	tokenStr, err := CreateOrSetToken(token)
	if err != nil {
		return "", model.NewError(err.Error())
	}
	logrus.Infof("用户登陆成功: %s", user.Username)
	return tokenStr, nil
}

func (s *Service) Logout(token *model.Token) *model.Error {
	if token == nil {
		return model.NewError("token is nil")
	}
	if err := db.RDB.Del(s.ctx, model.TokenKeyName(token.Username, token.TokenType)).Err(); err != nil {
		return model.NewError(err.Error())
	}
	return nil
}
