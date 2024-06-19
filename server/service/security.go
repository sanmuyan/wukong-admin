package service

import (
	"errors"
	"github.com/sanmuyan/xpkg/xbcrypt"
	"github.com/sanmuyan/xpkg/xcrypto"
	"github.com/sanmuyan/xpkg/xresponse"
	"wukong/pkg/config"
	"wukong/pkg/db"
	"wukong/pkg/security"
	"wukong/pkg/util"
	"wukong/server/model"
)

func (s *Service) ModifyPassword(req *model.ModifyPasswordRequest, token *model.Token) util.RespError {
	if !xbcrypt.IsPasswordComplexity(req.NewPassword, config.PasswordMinLength, config.PasswordMinIncludeCase) {
		return util.NewRespError(errors.New("密码格式不正确"), true).WithCode(xresponse.HttpBadRequest)
	}
	newPassword := xbcrypt.CreatePassword(req.NewPassword)
	if err := db.DB.Where("id = ?", token.GetUserID()).Updates(&model.User{Password: newPassword}).Error; err != nil {
		return util.NewRespError(err)
	}
	return nil
}

func (s *Service) GetClientEncryptPublicKey() (any, util.RespError) {
	var cert model.Cert
	if err := db.DB.Select("public_key").Where("purpose = ?", model.RSAPurposeClientEncrypt).First(&cert).Error; err != nil {
		return nil, util.NewRespError(err)
	}
	data := make(map[string]string)
	data["public_key"] = cert.PublicKey
	return data, nil
}

func (s *Service) DecryptClientData(ciphertext string) (string, error) {
	plaintext, err := xcrypto.DecryptPKCSRSA(ciphertext, security.ClientEncryptPrivateKey)
	if err != nil {
		return "", err
	}
	return plaintext, nil
}
