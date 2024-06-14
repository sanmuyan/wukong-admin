package mfalogin

import "wukong/server/model"

type MFAProvider interface {
	IsApprove(string, int) (bool, error)
}

var MFAProviders = make(map[string]MFAProvider)

func init() {
	MFAProviders[model.MFAProviderMFAApp] = NewMFAAppLogin()
}
