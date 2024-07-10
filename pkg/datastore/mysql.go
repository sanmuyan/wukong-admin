package datastore

import (
	"context"
	"github.com/sirupsen/logrus"
	"time"
	"wukong/pkg/db"
	"wukong/server/model"
)

type MySQLStore struct {
}

func NewMySQLStore() *MySQLStore {
	return &MySQLStore{}
}

func (c *MySQLStore) Clean(ctx context.Context) DataStore {
	go c.cleanTask(ctx)
	return c
}

// cleanTask 使用 mysql 存储临时数据，需要自行清理过期数据
func (c *MySQLStore) cleanTask(ctx context.Context) {
	logrus.Info("start mysql data store clean task")
	ticker := time.NewTicker(5 * time.Minute)
	for {
		select {
		case <-ctx.Done():
			ticker.Stop()
		case <-ticker.C:
			db.DB.Where("expires_at < ?", time.Now().UTC().Format("2006-01-02 15:04:05")).Delete(&model.Session{})
		}
	}
}

func (c *MySQLStore) StoreSession(s *model.Session, username ...string) error {
	return db.DB.Create(s).Error
}

func (c *MySQLStore) LoadSession(sessionID, sessionType string, sessionRaw any, username ...string) (*model.Session, bool) {
	var session model.Session
	tx := db.DB.Where("session_id = ? AND session_type = ?", sessionID, sessionType).First(&session)
	if tx.RowsAffected == 0 {
		return nil, false
	}
	if session.IsExpired() {
		return nil, false
	}
	if sessionRaw == nil {
		return nil, true
	}
	err := session.UnmarshalSessionRaw(sessionRaw)
	if err != nil {
		return nil, false
	}
	return &session, true
}

func (c *MySQLStore) DeleteSession(sessionID, sessionType string, username ...string) error {
	return db.DB.Where("session_id = ? AND session_type = ?", sessionID, sessionType).Delete(&model.Session{}).Error
}

func (c *MySQLStore) DeleteSessions(sessionType, username string) error {
	return db.DB.Where("session_type = ? AND username = ?", sessionType, username).Delete(&model.Session{}).Error
}
