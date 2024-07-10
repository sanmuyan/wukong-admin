package mfalogin

import "wukong/server/model"

// MFAProvider MFA 验证方式
type MFAProvider interface {
	// IsApprove 验证用户提交的验证码
	IsApprove(string, int) (bool, error)
}

var MFAProviders = make(map[string]MFAProvider)

func init() {
	MFAProviders[model.MFAProviderMFAApp] = NewMFAAppLogin()
}
