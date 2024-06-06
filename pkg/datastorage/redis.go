package datastorage

import (
	"context"
	"encoding/json"
	"fmt"
	"time"
	"wukong/pkg/db"
	"wukong/server/model"
)

type RDBStorage struct {
	ctx context.Context
}

func NewRDBStorage() *RDBStorage {
	return &RDBStorage{ctx: context.Background()}
}

func (c *RDBStorage) StoreToken(tokenID string, tokenType string, tokenStr string, e time.Duration) error {
	return db.RDB.Set(c.ctx, c.generateTokenKey(tokenID, tokenType), tokenStr, e).Err()
}

func (c *RDBStorage) DeleteToken(tokenID string, tokenType string) error {
	return db.RDB.Del(c.ctx, c.generateTokenKey(tokenID, tokenType)).Err()
}

func (c *RDBStorage) IsTokenExist(tokenID string, tokenType string, tokenStr string) bool {
	res, err := db.RDB.Get(c.ctx, c.generateTokenKey(tokenID, tokenType)).Result()
	if err != nil {
		return false
	}
	if res == tokenStr {
		return true
	}
	return false
}

func (c *RDBStorage) generateTokenKey(tokenID string, tokenType string) string {
	return fmt.Sprintf("%s:tokens:%s:%s", model.AppName, tokenType, tokenID)
}

func (c *RDBStorage) StoreCode(code *model.OauthCode) error {
	codeStr, _ := json.Marshal(code)
	return db.RDB.Set(context.Background(), c.generateCodeKey(code.Code, code.ClientID), codeStr, code.ExpiresAt.Sub(time.Now().UTC())).Err()
}

func (c *RDBStorage) LoadCode(code string, clientID string) (*model.OauthCode, error) {
	var codeModel model.OauthCode
	codeStr, err := db.RDB.Get(context.Background(), c.generateCodeKey(code, clientID)).Result()
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal([]byte(codeStr), &codeModel)
	return &codeModel, err
}

func (c *RDBStorage) DeleteCode(code string, clientID string) error {
	return db.RDB.Del(context.Background(), c.generateCodeKey(code, clientID)).Err()
}

func (c *RDBStorage) generateCodeKey(code string, clientID string) string {
	return fmt.Sprintf("%s:codes:%s:%s", model.AppName, clientID, code)
}
