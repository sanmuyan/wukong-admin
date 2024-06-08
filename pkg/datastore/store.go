package datastore

import (
	"github.com/sirupsen/logrus"
	"wukong/pkg/config"
	"wukong/server/model"
)

type DataStore interface {
	IsTokenExist(*model.StoreToken) bool
	StoreToken(*model.StoreToken) error
	DeleteToken(*model.StoreToken) error
	StoreCode(code *model.OauthCode) error
	LoadCode(string, string) (*model.OauthCode, error)
	DeleteCode(string, string) error
}

var DS DataStore

func InitStore() {
	dataStores := map[string]DataStore{
		"redis": NewRDBStore(),
		"mysql": NewMySQLStore(),
	}
	if _, ok := dataStores[config.Conf.DataStore]; !ok {
		logrus.Fatalf("data store %s not supported", config.Conf.DataStore)
	}
	DS = dataStores[config.Conf.DataStore]
}
