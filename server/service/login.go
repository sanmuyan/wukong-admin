package service

import (
	"errors"
	"github.com/sanmuyan/xpkg/xresponse"
	"github.com/sirupsen/logrus"
	"wukong/pkg/config"
	"wukong/pkg/db"
	"wukong/pkg/usersource"
	"wukong/pkg/util"
	"wukong/server/model"
)

func (s *Service) Login(login model.LoginRequest) (res *model.LoginResponse, ue util.RespError) {
	var user model.User
	user.Username = login.Username
	if err := db.DB.Where(&model.User{Username: user.Username}).First(&user).Error; err != nil {
		return nil, util.NewRespError(err)
	}

	if user.ID == 0 || user.IsActive != 1 {
		return nil, util.NewRespError(errors.New("用户名或密码错误"), true).WithCode(xresponse.HttpUnauthorized)
	}

	us, ok := usersource.UserSources[user.Source]
	if !ok {
		return nil, util.NewRespError(errors.New("未知的用户来源"), false)
	}
	if !us.Login(login.Username, login.Password) {
		return nil, util.NewRespError(errors.New("用户名或密码错误"), true).WithCode(xresponse.HttpUnauthorized)
	}
	return s.login(&user)
}

func (s *Service) login(user *model.User) (*model.LoginResponse, util.RespError) {
	passkey, re := s.BeginPassKeyLogin(&model.PassKeyBeginLoginRequest{Username: user.Username})
	if re != nil {
		if re.Err.Error() != "Found no credentials for user" {
			return nil, re
		}
	}
	if passkey != nil {
		return &model.LoginResponse{
			PassKeyBeginLogin: passkey,
		}, nil
	}
	mfaLogin, err := s.mfaBeginLogin(user)
	if err != nil {
		return nil, util.NewRespError(err)
	}
	if mfaLogin != nil {
		return &model.LoginResponse{
			MFABeginLogin: mfaLogin,
		}, nil
	}
	return s.createLoginToken(user.ID, user.Username)
}

func (s *Service) createLoginToken(userID int, username string) (res *model.LoginResponse, ue util.RespError) {
	token := model.Token{
		Username:    username,
		AccessLevel: s.GetMaxAccessLevel(s.GetUserRoles(userID)),
		TokenType:   model.SessionToken,
	}
	tokenStr, err := s.CreateOrSetToken(&token, config.Conf.TokenTTL)
	if err != nil {
		return nil, util.NewRespError(err)
	}
	logrus.Infof("用户登陆成功: %s", username)
	return &model.LoginResponse{Token: tokenStr}, nil

}

func (s *Service) Logout(token *model.Token) util.RespError {
	if err := s.DeleteToken(token); err != nil {
		return util.NewRespError(err)
	}
	return nil
}
