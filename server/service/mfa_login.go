package service

import (
	"errors"
	"time"
	"wukong/pkg/config"
	"wukong/pkg/datastore"
	"wukong/pkg/db"
	"wukong/pkg/mfalogin"
	"wukong/pkg/security"
	"wukong/pkg/util"
	"wukong/server/model"
)

func (s *Service) mfaBeginLogin(user *model.User) (*model.MFABeginLoginResponse, error) {
	var mfaApp = model.MFAApp{}
	tx := db.DB.Where("user_id = ?", user.ID).Find(&mfaApp).Limit(1)
	if tx.Error != nil {
		return nil, tx.Error
	}
	if tx.RowsAffected > 0 {
		sessionID := security.GetSessionID()
		err := datastore.DS.StoreSession(model.NewSession(sessionID, model.SessionTypeMFALogin, user.ID, user.Username, model.MFALoginSession{
			MFAProvider: model.MFAProviderMFAApp,
		}).SetTimeout(config.MFALoginTimeoutMin * time.Minute))
		if err != nil {
			return nil, err
		}
		return &model.MFABeginLoginResponse{
			SessionID:   sessionID,
			Username:    user.Username,
			MFAProvider: model.MFAProviderMFAApp,
		}, nil
	}
	return nil, nil
}

func (s *Service) MFAFinishLogin(req *model.MFAFinishLoginRequest) (*model.LoginResponse, util.RespError) {
	var mfaLoginSession = model.MFALoginSession{}
	session, ok := datastore.DS.LoadSession(req.SessionID, model.SessionTypeMFALogin, &mfaLoginSession)
	if !ok {
		return nil, util.NewRespError(errors.New("登陆超时"))
	}
	defer func() {
		// session 只能使用一次
		_ = datastore.DS.DeleteSession(req.SessionID, model.SessionTypeMFALogin)
	}()
	ap, ok := mfalogin.MFAProviders[req.MFAProvider]
	if !ok {
		return nil, util.NewRespError(errors.New("未知的 MFA 验证方式"))
	}
	ia, err := ap.IsApprove(req.Code, session.UserID)
	if err != nil {
		return nil, util.NewRespError(errors.New("验证失败"), true)
	}
	if !ia {
		return nil, util.NewRespError(errors.New("验证失败"), true)
	}
	return s.createLoginToken(session.UserID, session.Username)
}
