package service

import (
	"errors"
	"github.com/sanmuyan/xpkg/xresponse"
	"github.com/sanmuyan/xpkg/xutil"
	"github.com/sirupsen/logrus"
	"time"
	"wukong/pkg/config"
	"wukong/pkg/datastore"
	"wukong/pkg/db"
	"wukong/pkg/usersource"
	"wukong/pkg/util"
	"wukong/server/model"
)

func (s *Service) Login(login model.LoginRequest) (res *model.LoginResponse, ue util.RespError) {
	var err error
	login.Password, err = s.DecryptClientData(login.Password)
	if err != nil {
		return nil, util.NewRespError(err).WithCode(xresponse.HttpBadRequest)
	}
	var user model.User
	user.Username = login.Username
	db.DB.Where(&model.User{Username: user.Username}).First(&user)
	if user.ID == 0 || user.IsActive != 1 {
		return nil, util.NewRespError(errors.New("用户名或密码错误"), true).WithCode(xresponse.HttpUnauthorized)
	}
	var logSecurity model.LoginSecurity
	db.DB.Where("user_id = ?", user.ID).First(&logSecurity)
	if logSecurity.LockAt != nil {
		if time.Now().UTC().Before(*logSecurity.LockAt) {
			return nil, util.NewRespError(errors.New("用户名或密码错误"), true).WithCode(xresponse.HttpUnauthorized)
		}
	}

	us, ok := usersource.UserSources[user.Source]
	if !ok {
		return nil, util.NewRespError(errors.New("未知的用户来源"), false)
	}
	if !us.Login(login.Username, login.Password) {
		go s.updateLoginFail(&logSecurity)
		return nil, util.NewRespError(errors.New("用户名或密码错误"), true).WithCode(xresponse.HttpUnauthorized)
	}
	return s.mfaLogin(&user)
}

func (s *Service) mfaLogin(user *model.User) (*model.LoginResponse, util.RespError) {
	//passkey, re := s.BeginPassKeyLogin(&model.PassKeyBeginLoginRequest{Username: user.Username})
	//if re != nil {
	//	if re.Err.Error() != "Found no credentials for user" {
	//		return nil, re
	//	}
	//}
	//if passkey != nil {
	//	return &model.LoginResponse{
	//		PassKeyBeginLogin: passkey,
	//	}, nil
	//}
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
		TokenType:   model.TokenTypeSession,
	}
	token.SetUserID(userID)
	tokenStr, err := s.CreateOrSetToken(&token, config.Conf.TokenTTL)
	if err != nil {
		return nil, util.NewRespError(err)
	}
	logrus.Infof("用户登陆成功: %s", username)
	go s.updateLoginSuccess(userID)
	return &model.LoginResponse{Token: tokenStr}, nil

}

func (s *Service) Logout(token *model.Token) util.RespError {
	if err := s.DeleteTokenSession(token); err != nil {
		return util.NewRespError(err)
	}
	return nil
}

func (s *Service) LogoutAll(token *model.Token) util.RespError {
	err := datastore.DS.DeleteSessions(model.SessionTypeSessionToken, token.Username)
	if err != nil {
		return util.NewRespError(err)
	}
	return nil
}

func (s *Service) updateLoginSuccess(userID int) {
	var ls model.LoginSecurity
	db.DB.Where("user_id = ?", userID).First(&ls)
	ls.UserID = userID
	ls.LastLoginAt = xutil.PtrTo(time.Now().UTC())
	ls.LoginFailCount = nil
	ls.LockAt = nil
	db.DB.Save(&ls)
}

func (s *Service) updateLoginFail(ls *model.LoginSecurity) {
	if ls.LoginFailCount == nil {
		ls.LoginFailCount = xutil.PtrTo(0)
	}
	ls.LoginFailCount = xutil.PtrTo(*ls.LoginFailCount + 1)
	if *ls.LoginFailCount >= 10 {
		ls.LockAt = xutil.PtrTo(time.Now().UTC().Add(time.Minute * 1))
	}
	db.DB.Save(ls)
}
