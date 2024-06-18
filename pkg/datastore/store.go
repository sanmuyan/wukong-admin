package datastore

import (
	"github.com/sirupsen/logrus"
	"wukong/pkg/config"
	"wukong/server/model"
)

type DataStore interface {
	StoreSession(s *model.Session, username ...string) error
	LoadSession(sessionID, sessionType string, sessionRaw any, username ...string) (*model.Session, bool)
	DeleteSession(sessionID, sessionType string, username ...string) error
	DeleteSessions(sessionType, username string) error
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
