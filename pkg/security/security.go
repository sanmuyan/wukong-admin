package security

import (
	"crypto/rsa"
	"encoding/base64"
	"github.com/sanmuyan/xpkg/xcrypto"
	"github.com/sirupsen/logrus"
	"wukong/pkg/db"
	"wukong/server/model"
)

var ClientEncryptPrivateKey *rsa.PrivateKey

func InitSecurity() {
	var cert model.Cert
	if err := db.DB.Select("private_key").Where("purpose = ?", model.RSAPurposeClientEncrypt).First(&cert).Error; err != nil {
		logrus.Fatalf("get client encrypt private key error: %s", err.Error())
	}
	privateKeyText, err := base64.StdEncoding.DecodeString(cert.PrivateKey)
	if err != nil {
		logrus.Fatalf("decode client encrypt private key error: %s", err.Error())
	}
	privateKey, err := xcrypto.TextToRSAPrivateKey(privateKeyText)
	if err != nil {
		logrus.Fatalf("decode client encrypt private key error: %s", err.Error())
	}
	ClientEncryptPrivateKey = privateKey
}
