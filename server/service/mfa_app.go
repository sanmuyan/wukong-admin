package service

import (
	"fmt"
	"github.com/pkg/errors"
	"github.com/sanmuyan/xpkg/xmfa"
	"github.com/sanmuyan/xpkg/xresponse"
	"github.com/sanmuyan/xpkg/xutil"
	"time"
	"wukong/pkg/config"
	"wukong/pkg/datastore"
	"wukong/pkg/db"
	"wukong/pkg/util"
	"wukong/server/model"
)

func (s *Service) GetMFAAppStatus(token *model.Token) (any, util.RespError) {
	data := make(map[string]bool)
	data["is_bind"] = false
	tx := db.DB.Select("user_id").Where("user_id = ?", token.GetUserID()).First(&model.MFAApp{})
	if tx.RowsAffected > 0 {
		data["is_bind"] = true
	}
	return data, nil
}

func (s *Service) MFAAppBeginBind(token *model.Token) (*model.MFAAppBindResponse, util.RespError) {
	totpSecret := xutil.RemoveError(xmfa.GenerateTOTPSecret(config.TOTPSecretLength))
	sessionID := util.GetRandomID()
	err := datastore.DS.StoreSession(model.NewSession(sessionID, model.SessionTypeMFAAppBind, token.GetUserID(), token.Username, &model.MFAAppBindSession{
		TOTPSecret: totpSecret,
	}).SetTimeout(config.MFABindTimeoutMin * time.Minute))
	if err != nil {
		return nil, util.NewRespError(err)
	}
	return &model.MFAAppBindResponse{
		SessionID:  sessionID,
		TOTPSecret: totpSecret,
		QRCodeURI:  fmt.Sprintf("otpauth://totp/%s:%s?secret=%s", config.Conf.AppName, token.Username, totpSecret),
		TimeoutMin: config.MFABindTimeoutMin,
	}, nil
}

func (s *Service) MFAAppFinishBind(req *model.MFAAppBindRequest, token *model.Token) util.RespError {
	var mfaBindSession model.MFAAppBindSession
	_, ok := datastore.DS.LoadSession(req.SessionID, model.SessionTypeMFAAppBind, &mfaBindSession)
	if !ok {
		return util.NewRespError(errors.New("绑定超时"), true).WithCode(xresponse.HttpBadRequest)
	}
	defer func() {
		_ = datastore.DS.DeleteSession(req.SessionID, model.SessionTypeMFAAppBind)
	}()
	if req.TOTPCode != xutil.RemoveError(xmfa.GetTOTPToken(mfaBindSession.TOTPSecret, config.TOTPTokenInterval)) {
		return util.NewRespError(errors.New("验证码错误"), true).WithCode(xresponse.HttpBadRequest)
	}
	mfaApp := model.MFAApp{
		UserID:     token.GetUserID(),
		TOTPSecret: mfaBindSession.TOTPSecret,
	}
	if err := db.DB.Create(&mfaApp).Error; err != nil {
		return util.NewRespError(err)
	}
	return nil
}

func (s *Service) DeleteMFAApp(token *model.Token) util.RespError {
	if err := db.DB.Where("user_id = ?", token.GetUserID()).Delete(&model.MFAApp{}).Error; err != nil {
		return util.NewRespError(err)
	}
	return nil
}
