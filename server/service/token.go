package service

import (
	"github.com/sanmuyan/xpkg/xjwt"
	"github.com/sanmuyan/xpkg/xutil"
	"time"
	"wukong/pkg/config"
	"wukong/pkg/tokenclient"
	"wukong/server/model"
)

func (s *Service) CreateOrSetToken(token *model.Token) (string, error) {
	ttl := time.Duration(config.Conf.TokenTTL) * time.Second
	token.Issuer = model.AppName
	if token.TokenType == model.ApiToken {
		ttl = 0
	} else {
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
	if err = tokenclient.TC.SetToken(token.Username, token.TokenType, tokenStr, ttl); err != nil {
		return "", err
	}
	return tokenStr, err
}

func (s *Service) DeleteToken(token *model.Token) error {
	return tokenclient.TC.DeleteToken(token.Username, token.TokenType)
}
