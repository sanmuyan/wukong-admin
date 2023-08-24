package service

import (
	"errors"
	tokenutil "github.com/sanmuyan/dao/token"
	"time"
	"wukong/pkg/config"
	"wukong/pkg/db"
	"wukong/server/model"
)

func (s *Service) CreateOrSetToken(token model.Token) (tokenStr string, err error) {
	if token.TokenType != "session" && token.TokenType != "api" {
		return tokenStr, errors.New("tokenType must session or api")
	}
	ttl := time.Duration(config.Conf.TokenTTL) * time.Second
	if token.TokenType == "session" {
		token.ExpiresTime = time.Now().UTC().Add(ttl).Unix()
	}
	tokenStr, err = tokenutil.CreateToken(token, config.Conf.Secret.TokenKey)
	if err != nil {
		return tokenStr, err
	}
	if err = db.RDB.Set(s.ctx, model.TokenKeyName(token.Username, token.TokenType), tokenStr, ttl).Err(); err != nil {
		return tokenStr, err
	}
	return tokenStr, err
}
