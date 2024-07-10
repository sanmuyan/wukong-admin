package datastore

import (
	"context"
	"github.com/sirupsen/logrus"
	"wukong/pkg/config"
	"wukong/server/model"
)

// DataStore 数据存储类型
type DataStore interface {
	// StoreSession 存储 session
	StoreSession(s *model.Session, username ...string) error
	// LoadSession 加载 session
	LoadSession(sessionID, sessionType string, sessionRaw any, username ...string) (*model.Session, bool)
	// DeleteSession 删除 session
	DeleteSession(sessionID, sessionType string, username ...string) error
	// DeleteSessions 删除用户的所有 session
	DeleteSessions(sessionType, username string) error
	// Clean 清理过期数据
	Clean(ctx context.Context) DataStore
}

var DS DataStore

func InitStore(ctx context.Context) {
	dataStores := map[string]DataStore{
		"redis": NewRDBStore().Clean(ctx),
		"mysql": NewMySQLStore().Clean(ctx),
	}
	if _, ok := dataStores[config.Conf.DataStore]; !ok {
		logrus.Fatalf("data store %s not supported", config.Conf.DataStore)
	}
	DS = dataStores[config.Conf.DataStore]
}
