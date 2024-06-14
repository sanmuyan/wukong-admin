package service

import (
	"fmt"
	"github.com/pkg/errors"
	"github.com/sanmuyan/xpkg/xmfa"
	"github.com/sanmuyan/xpkg/xresponse"
	"github.com/sanmuyan/xpkg/xutil"
	"time"
	"wukong/pkg/config"
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
	db.DB.Where("user_id = ?", token.GetUserID()).Delete(&model.MFAAppBindSession{})
	totpSecret := xutil.RemoveError(xmfa.GenerateTOTPSecret(config.TOTPSecretLength))
	mfaBindSession := model.MFAAppBindSession{
		UserID:     token.GetUserID(),
		TOTPSecret: totpSecret,
		ExpireAt:   time.Now().UTC().Add(config.MFABindTimeoutMin * time.Minute),
	}
	err := db.DB.Create(&mfaBindSession).Error
	if err != nil {
		return nil, util.NewRespError(err)
	}
	return &model.MFAAppBindResponse{
		TOTPSecret: totpSecret,
		QRCodeURI:  fmt.Sprintf("otpauth://totp/%s:%s?secret=%s", config.Conf.AppName, token.Username, totpSecret),
		TimeoutMin: config.MFABindTimeoutMin,
	}, nil
}

func (s *Service) MFAAppFinishBind(req *model.MFAAppBindRequest, token *model.Token) util.RespError {
	var mfaBindSession model.MFAAppBindSession
	tx := db.DB.Where("user_id = ? AND totp_secret = ?", token.GetUserID(), req.TOTPSecret).First(&mfaBindSession)
	if tx.RowsAffected == 0 {
		return util.NewRespError(errors.New("绑定错误"), true).WithCode(xresponse.HttpBadRequest)
	}
	if mfaBindSession.ExpireAt.Before(time.Now().UTC()) {
		return util.NewRespError(errors.New("绑定超时"), true).WithCode(xresponse.HttpBadRequest)
	}
	if req.TOTPCode != xutil.RemoveError(xmfa.GetTOTPToken(mfaBindSession.TOTPSecret, config.TOTPTokenInterval)) {
		return util.NewRespError(errors.New("验证码错误"), true).WithCode(xresponse.HttpBadRequest)
	}
	mfa := model.MFAApp{
		UserID:     token.GetUserID(),
		TOTPSecret: mfaBindSession.TOTPSecret,
	}
	if err := db.DB.Create(&mfa).Error; err != nil {
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
