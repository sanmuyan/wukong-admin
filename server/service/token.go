package service

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/sanmuyan/xpkg/xjwt"
	"github.com/sanmuyan/xpkg/xutil"
	"time"
	"wukong/pkg/config"
	"wukong/pkg/datastore"
	"wukong/pkg/db"
	"wukong/server/model"
)

func (s *Service) CreateOrSetToken(token *model.Token, expiresAt int) (string, error) {
	token.Issuer = config.Conf.Basic.AppName
	token.TokenID = uuid.NewString()
	ttl := time.Duration(expiresAt) * time.Second
	if ttl > 0 {
		token.ExpiresAt = xutil.PtrTo[int64](time.Now().Add(ttl).Unix())
	}
	token.IssuedAt = time.Now().Unix()
	err := token.Valid()
	if err != nil {
		return "", err
	}
	tokenStr, err := xjwt.CreateToken(token, config.Conf.Secret.TokenKey)
	if err != nil {
		return "", err
	}
	st := model.NewTokenSession(token).WithTokenStr(tokenStr)
	if token.GetUserID() == 0 {
		var user model.User
		db.DB.Select("id").Where("username", token.Username).First(&user)
		token.SetUserID(user.ID)
	}
	session := model.NewSession(token.TokenID, token.TokenType, token.GetUserID(), token.Username, st)
	if expiresAt > 0 {
		session = session.SetTimeout(time.Duration(expiresAt) * time.Second)
	}
	err = datastore.DS.StoreSession(session, token.Username)
	if err != nil {
		return "", err
	}
	return tokenStr, err
}

func (s *Service) DeleteTokenSession(token *model.Token) error {
	return datastore.DS.DeleteSession(token.TokenID, token.TokenType, fmt.Sprintf("%s:%s:%s", token.TokenType, token.Username, token.TokenID))
}
