package service

import (
	"context"
	"encoding/json"
	"errors"
	"time"
	"wukong/pkg/config"
	"wukong/pkg/db"
	"wukong/pkg/util"
	"wukong/server/model"
)

func CreateOrSetToken(token model.Token) (tokenStr string, err error) {
	var ctx = context.Background()
	token.Timestamp = time.Now().Unix()
	if token.TokenType != "session" && token.TokenType != "api" {
		return tokenStr, errors.New("tokenType must session or api")
	}
	tokenJson, _ := json.Marshal(token)
	tokenClaims := util.TokenClaims{
		Body: tokenJson,
	}
	tokenStr, err = util.CreateToken(tokenClaims, config.Conf.Secret.TokenKey)
	if err != nil {
		return tokenStr, err
	}
	if err = db.RDB.Set(ctx, model.TokenKeyName(token.Username, token.TokenType), tokenStr, time.Duration(token.TTL*1000000000)).Err(); err != nil {
		return tokenStr, err
	}
	return tokenStr, err
}
