package datastore

import (
	"context"
	"encoding/json"
	"fmt"
	"time"
	"wukong/pkg/db"
	"wukong/server/model"
)

type RDBStore struct {
	ctx context.Context
}

func NewRDBStore() *RDBStore {
	return &RDBStore{ctx: context.Background()}
}

func (c *RDBStore) StoreToken(ts *model.StoreToken) error {
	var expires time.Duration
	if ts.ExpiresAt != nil {
		expires = ts.ExpiresAt.Sub(time.Now().UTC())
	}
	return db.RDB.Set(c.ctx, c.generateTokenKey(ts), ts.TokenStr, expires).Err()
}

func (c *RDBStore) DeleteToken(ts *model.StoreToken) error {
	return db.RDB.Del(c.ctx, c.generateTokenKey(ts)).Err()
}

func (c *RDBStore) IsTokenExist(ts *model.StoreToken) bool {
	res, err := db.RDB.Exists(c.ctx, c.generateTokenKey(ts)).Result()
	if err != nil || res == 0 {
		return false
	}
	return false
}

func (c *RDBStore) generateTokenKey(ts *model.StoreToken) string {
	return fmt.Sprintf("%s:store_tokens:%s:%s:%s", model.AppName, ts.TokenType, ts.Username, ts.UUID)
}

func (c *RDBStore) StoreCode(code *model.OauthCode) error {
	codeStr, _ := json.Marshal(code)
	return db.RDB.Set(context.Background(), c.generateCodeKey(code.Code, code.ClientID), codeStr, code.ExpiresAt.Sub(time.Now().UTC())).Err()
}

func (c *RDBStore) LoadCode(code string, clientID string) (*model.OauthCode, error) {
	var codeModel model.OauthCode
	codeStr, err := db.RDB.Get(context.Background(), c.generateCodeKey(code, clientID)).Result()
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal([]byte(codeStr), &codeModel)
	return &codeModel, err
}

func (c *RDBStore) DeleteCode(code string, clientID string) error {
	return db.RDB.Del(context.Background(), c.generateCodeKey(code, clientID)).Err()
}

func (c *RDBStore) generateCodeKey(code string, clientID string) string {
	return fmt.Sprintf("%s:codes:%s:%s", model.AppName, clientID, code)
}
