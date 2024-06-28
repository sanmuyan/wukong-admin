package datastore

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/sanmuyan/xpkg/xutil"
	"time"
	"wukong/pkg/config"
	"wukong/pkg/db"
	"wukong/server/model"
)

type RDBStore struct {
	ctx context.Context
}

func NewRDBStore() *RDBStore {
	return &RDBStore{ctx: context.Background()}
}

func (c *RDBStore) StoreSession(s *model.Session, username ...string) error {
	var expires time.Duration
	if s.ExpiresAt != nil {
		expires = s.ExpiresAt.Sub(time.Now().UTC())
	}
	return db.RDB.Set(c.ctx, c.generateSessionPath(s.SessionType, s.SessionID, username...), xutil.RemoveError(json.Marshal(s)), expires).Err()
}

func (c *RDBStore) LoadSession(sessionID, sessionType string, sessionRaw any, username ...string) (*model.Session, bool) {
	res, _ := db.RDB.Exists(c.ctx, c.generateSessionPath(sessionType, sessionID, username...)).Result()
	if res == 0 {
		return nil, false
	}
	var session model.Session
	sessionStr, _ := db.RDB.Get(c.ctx, c.generateSessionPath(sessionType, sessionID, username...)).Result()
	err := json.Unmarshal([]byte(sessionStr), &session)
	if err != nil {
		return nil, false
	}
	if session.IsExpired() {
		return nil, false
	}
	if sessionRaw == nil {
		return nil, false
	}
	err = json.Unmarshal([]byte(session.SessionRaw), &sessionRaw)
	if err != nil {
		return nil, false
	}
	return &session, true
}

func (c *RDBStore) DeleteSession(sessionID, sessionType string, username ...string) error {
	return db.RDB.Del(c.ctx, c.generateSessionPath(sessionID, sessionType, username...)).Err()
}

func (c *RDBStore) DeleteSessions(sessionType, username string) error {
	var cursor uint64
	var n int64
	for {
		keys, nextCursor, err := db.RDB.Scan(c.ctx, cursor, c.generateSessionPath(sessionType, fmt.Sprintf("%s:*", username)), n).Result()
		if err != nil {
			return err
		}
		if len(keys) > 0 {
			err = db.RDB.Del(c.ctx, keys...).Err()
			if err != nil {
				return err
			}
		}
		cursor = nextCursor
		if cursor == 0 {
			break
		}
	}
	return nil
}
func (c *RDBStore) generateSessionPath(sessionType, sessionID string, username ...string) string {
	if len(username) > 0 {
		return fmt.Sprintf("%s:sessions:%s:%s:%s", config.Conf.Basic.AppName, sessionType, username[0], sessionID)
	}
	return fmt.Sprintf("%s:sessions:%s:%s", config.Conf.Basic.AppName, sessionType, sessionID)
}
