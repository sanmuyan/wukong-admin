package model

import (
	"errors"
	"fmt"
	"github.com/sanmuyan/xpkg/xutil"
	"wukong/pkg/config"
)

type Token struct {
	userID      int
	TokenID     string `json:"token_id"`
	Username    string `json:"username" binding:"required"`
	AccessLevel int    `json:"access_level,omitempty"`
	TokenType   string `json:"token_type" binding:"required"`
	ExpiresAt   *int64 `json:"exp,omitempty"`
	Scope       string `json:"scope,omitempty"`
	ClientID    string `json:"client_id,omitempty"`
	IssuedAt    int64  `json:"iat"`
	Issuer      string `json:"iss"`
}

func (t *Token) SetUserID(userID int) {
	t.userID = userID
}

func (t *Token) GetUserID() int {
	return t.userID
}

func (t *Token) Valid() error {
	if t.Issuer != config.Conf.Basic.AppName {
		return errors.New("token is not required, issuer is not " + config.Conf.Basic.AppName)
	}
	if xutil.IsZero(t.Username, t.TokenType) {
		return errors.New("token is not required, username or token_type is empty")
	}
	if _, ok := TokenTypes[t.TokenType]; !ok {
		return errors.New("invalid token type")
	}
	if t.TokenType != TokenTypeApi && t.ExpiresAt == nil {
		return errors.New("token is not required, expires_at is empty")
	}
	return nil
}

var TokenTypes = map[string]struct{}{}

const (
	TokenTypeSession      = "session_token"
	TokenTypeApi          = "api_token"
	TokenTypeOauthAccess  = "oauth_access_token"
	TokenTypeOauthRefresh = "oauth_refresh_token"
)

func init() {
	TokenTypes[TokenTypeSession] = struct{}{}
	TokenTypes[TokenTypeApi] = struct{}{}
	TokenTypes[TokenTypeOauthAccess] = struct{}{}
	TokenTypes[TokenTypeOauthRefresh] = struct{}{}
}

type TokenSession struct {
	TokenType string `json:"token_type"`
	TokenStr  string `json:"token_str"`
}

func NewTokenSession(token *Token) *TokenSession {
	return &TokenSession{TokenType: token.TokenType}
}

func (c *TokenSession) WithTokenStr(tokenStr string) *TokenSession {
	c.TokenStr = tokenStr
	return c
}

func (c *TokenSession) GenerateCustomIndex(token *Token) string {
	return fmt.Sprintf("%s:%s", token.Username, token.TokenType)
}
