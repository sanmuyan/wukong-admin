package service

import (
	"github.com/sanmuyan/xpkg/xjwt"
	"github.com/sanmuyan/xpkg/xutil"
	"time"
	"wukong/pkg/config"
	"wukong/pkg/datastorage"
	"wukong/server/model"
)

func (s *Service) CreateOrSetToken(token *model.Token, tokenID string, expiresAt int) (string, error) {
	token.Issuer = model.AppName
	ttl := time.Duration(expiresAt) * time.Second
	if ttl > 0 {
		token.ExpiresAt = xutil.PtrTo[int64](time.Now().UTC().Add(ttl).Unix())
	}
	token.IssuedAt = time.Now().UTC().Unix()
	err := token.Valid()
	if err != nil {
		return "", err
	}
	tokenStr, err := xjwt.CreateToken(token, config.Conf.Secret.TokenID)
	if err != nil {
		return "", err
	}
	if err = datastorage.DS.StoreToken(tokenID, token.TokenType, tokenStr, ttl); err != nil {
		return "", err
	}
	return tokenStr, err
}

func (s *Service) DeleteToken(token *model.Token) error {
	return datastorage.DS.DeleteToken(token.Username, token.TokenType)
}
