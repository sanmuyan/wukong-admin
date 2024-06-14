package mfalogin

import (
	"errors"
	"github.com/sanmuyan/xpkg/xmfa"
	"wukong/pkg/config"
	"wukong/pkg/db"
	"wukong/server/model"
)

type MFAAppLogin struct {
}

func NewMFAAppLogin() *MFAAppLogin {
	return &MFAAppLogin{}
}

func (m *MFAAppLogin) IsApprove(code string, userID int) (bool, error) {
	var mfaApp = model.MFAApp{}
	tx := db.DB.Select("totp_secret").Where("user_id = ?", userID).First(&mfaApp)
	if tx.RowsAffected > 0 {
		tc, _ := xmfa.GetTOTPToken(mfaApp.TOTPSecret, config.TOTPTokenInterval)
		if tc == code {
			return true, nil
		}
	}
	return false, errors.New("验证码错误")
}
