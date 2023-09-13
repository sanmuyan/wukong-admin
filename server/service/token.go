package service

import (
	"github.com/sanmuyan/xpkg/xjwt"
	"time"
	"wukong/pkg/config"
	"wukong/pkg/db"
	"wukong/server/model"
)

func (s *Service) CreateOrSetToken(token model.Token) (string, error) {
	ttl := time.Duration(config.Conf.TokenTTL) * time.Second
	if token.TokenType == "session" {
		token.ExpiresTime = time.Now().UTC().Add(ttl).Unix()
	}
	err := token.Valid()
	if err != nil {
		return "", err
	}
	tokenStr, err := xjwt.CreateToken(token, config.Conf.Secret.TokenKey)
	if err != nil {
		return "", err
	}
	if err = db.RDB.Set(s.ctx, model.TokenKeyName(token.Username, token.TokenType), tokenStr, ttl).Err(); err != nil {
		return "", err
	}
	return tokenStr, err
}
