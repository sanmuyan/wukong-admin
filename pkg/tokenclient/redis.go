package tokenclient

import (
	"context"
	"time"
	"wukong/pkg/db"
	"wukong/server/model"
)

type RDBTokenClient struct {
	ctx context.Context
}

func NewRDBTokenClient() *RDBTokenClient {
	return &RDBTokenClient{ctx: context.Background()}
}

func (c *RDBTokenClient) SetToken(tokenKey string, tokenType string, tokenStr string, e time.Duration) error {
	return db.RDB.Set(c.ctx, c.generateKeyName(tokenKey, tokenType), tokenStr, e).Err()
}

func (c *RDBTokenClient) DeleteToken(tokenKey string, tokenType string) error {
	return db.RDB.Del(c.ctx, c.generateKeyName(tokenKey, tokenType)).Err()
}

func (c *RDBTokenClient) IsTokenExist(tokenKey string, tokenType string, tokenStr string) bool {
	res, err := db.RDB.Get(c.ctx, c.generateKeyName(tokenKey, tokenType)).Result()
	if err != nil {
		return false
	}
	if res == tokenStr {
		return true
	}
	return false
}

func (c *RDBTokenClient) generateKeyName(username string, tokenType string) string {
	return model.AppName + ":token:" + tokenType + ":" + username
}
