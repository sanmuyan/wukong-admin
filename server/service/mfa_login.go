package service

import (
	"errors"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
	"wukong/pkg/config"
	"wukong/pkg/db"
	"wukong/pkg/mfalogin"
	"wukong/pkg/util"
	"wukong/server/model"
)

func (s *Service) mfaBeginLogin(user *model.User) (*model.MFABeginLoginResponse, error) {
	db.DB.Where("user_id = ?", user.ID).Delete(&model.MFALoginSession{})
	var mfaApp = model.MFAApp{}
	tx := db.DB.Where("user_id = ?", user.ID).First(&mfaApp)
	if tx.Error != nil && !errors.Is(tx.Error, gorm.ErrRecordNotFound) {
		return nil, tx.Error
	}
	if tx.RowsAffected > 0 {
		newSessionID := uuid.NewString()
		mfaAppAuthSession := model.MFALoginSession{
			UserID:      user.ID,
			Username:    user.Username,
			SessionID:   newSessionID,
			MFAProvider: model.MFAProviderMFAApp,
			ExpireAt:    time.Now().UTC().Add(config.MFALoginTimeoutMin * time.Minute),
		}
		err := db.DB.Create(&mfaAppAuthSession).Error
		if err != nil {
			return nil, err
		}
		return &model.MFABeginLoginResponse{
			Username:    user.Username,
			SessionID:   newSessionID,
			MFAProvider: model.MFAProviderMFAApp,
		}, nil
	}
	return nil, nil
}

func (s *Service) MFAFinishLogin(req *model.MFAFinishLoginRequest) (*model.LoginResponse, util.RespError) {
	var mfaAuthSession = model.MFALoginSession{}
	tx := db.DB.Where("session_id = ? AND username = ? AND mfa_provider = ?", req.SessionID, req.Username, req.MFAProvider).First(&mfaAuthSession)
	if tx.RowsAffected == 0 {
		return nil, util.NewRespError(errors.New("未知 session"))
	}
	defer func() {
		db.DB.Delete(&model.MFALoginSession{}, mfaAuthSession.ID)
	}()
	if time.Now().UTC().After(mfaAuthSession.ExpireAt) {
		return nil, util.NewRespError(errors.New("登陆超时"), true)
	}
	ap, ok := mfalogin.MFAProviders[req.MFAProvider]
	if !ok {
		return nil, util.NewRespError(errors.New("未知的 MFA 验证方式"))
	}
	ia, err := ap.IsApprove(req.Code, mfaAuthSession.UserID)
	if err != nil {
		return nil, util.NewRespError(errors.New("验证失败"), true)
	}
	if !ia {
		return nil, util.NewRespError(errors.New("验证失败"), true)
	}
	return s.createLoginToken(mfaAuthSession.UserID, mfaAuthSession.Username)
}
