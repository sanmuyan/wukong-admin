package datastorage

import (
	"github.com/sirupsen/logrus"
	"time"
	"wukong/pkg/config"
	"wukong/server/model"
)

type DataStorage interface {
	IsTokenExist(string, string, string) bool
	StoreToken(string, string, string, time.Duration) error
	DeleteToken(string, string) error
	StoreCode(code *model.OauthCode) error
	LoadCode(string, string) (*model.OauthCode, error)
	DeleteCode(string, string) error
}

var DS DataStorage

func InitTokenStorage() {
	tokenStorages := map[string]DataStorage{
		"redis": NewRDBStorage(),
		"mysql": NewMySQLStorage(),
	}
	if _, ok := tokenStorages[config.Conf.DataStorage]; !ok {
		logrus.Fatalf("data storage %s not supported", config.Conf.DataStorage)
	}
	DS = tokenStorages[config.Conf.DataStorage]
}
