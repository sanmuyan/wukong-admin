package security

import (
	"encoding/base64"
	"github.com/sanmuyan/xpkg/xcrypto"
	"github.com/sirupsen/logrus"
	"sync"
	"wukong/pkg/db"
	"wukong/server/model"
)

var PrivateKeys *sync.Map

func InitPrivateKeys() {
	PrivateKeys = new(sync.Map)
	var certs []*model.Cert
	if err := db.DB.Find(&certs).Error; err != nil {
		logrus.Errorf("get client encrypt private key error: %s", err.Error())
	}
	for _, cert := range certs {
		privateKeyText, err := base64.StdEncoding.DecodeString(cert.PrivateKey)
		if err != nil {
			logrus.Fatalf("decode client encrypt private key error: %s", err.Error())
		}
		privateKey, err := xcrypto.TextToRSAPrivateKey(privateKeyText)
		if err != nil {
			logrus.Errorf("decode client encrypt private key error: %s", err.Error())
			continue
		}
		PrivateKeys.Store(cert.Purpose, privateKey)
	}
}
