package service

import (
	"github.com/google/uuid"
	"github.com/sanmuyan/xpkg/xjwt"
	"github.com/sanmuyan/xpkg/xutil"
	"time"
	"wukong/pkg/config"
	"wukong/pkg/datastore"
	"wukong/server/model"
)

func (s *Service) CreateOrSetToken(token *model.Token, expiresAt int) (string, error) {
	token.Issuer = model.AppName
	token.UUID = uuid.NewString()
	ttl := time.Duration(expiresAt) * time.Second
	if ttl > 0 {
		token.ExpiresAt = xutil.PtrTo[int64](time.Now().UTC().Add(ttl).Unix())
	}
	token.IssuedAt = time.Now().UTC().Unix()
	err := token.Valid()
	if err != nil {
		return "", err
	}
	tokenStr, err := xjwt.CreateToken(token, config.Conf.Secret.TokenKey)
	if err != nil {
		return "", err
	}
	st := model.NewTokenStore(token).WithTokenStr(tokenStr).WithExpiresAt(token.ExpiresAt)
	if err = datastore.DS.StoreToken(st); err != nil {
		return "", err
	}
	return tokenStr, err
}

func (s *Service) DeleteToken(token *model.Token) error {
	return datastore.DS.DeleteToken(model.NewTokenStore(token))
}
