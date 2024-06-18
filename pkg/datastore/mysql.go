package datastore

import (
	"wukong/pkg/db"
	"wukong/server/model"
)

type MySQLStore struct {
}

func NewMySQLStore() *MySQLStore {
	return &MySQLStore{}
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
	return db.DB.Where("session_type = ? AND session_id = ?", username, sessionType).Delete(&model.Session{}).Error
}
