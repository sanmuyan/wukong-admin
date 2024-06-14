package model

import (
	"errors"
	"github.com/sanmuyan/xpkg/xutil"
	"time"
	"wukong/pkg/config"
)

type Token struct {
	userID      int
	UUID        string `json:"uuid"`
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
	err := errors.New("token is not required")
	if t.Issuer != config.Conf.AppName {
		return err
	}
	if xutil.IsZero(t.Username, t.TokenType) {
		return err
	}
	if _, ok := TokenTypes[t.TokenType]; !ok {
		return err
	}
	if t.TokenType != ApiToken && t.ExpiresAt == nil {
		return err
	}
	return nil
}

var TokenTypes = map[string]struct{}{}

const (
	SessionToken      = "session"
	ApiToken          = "api"
	OauthAccessToken  = "oauth_access_token"
	OauthRefreshToken = "oauth_refresh_token"
)

func init() {
	TokenTypes[SessionToken] = struct{}{}
	TokenTypes[ApiToken] = struct{}{}
	TokenTypes[OauthAccessToken] = struct{}{}
	TokenTypes[OauthRefreshToken] = struct{}{}
}

type StoreToken struct {
	UUID      string     `gorm:"<-:create"`
	Username  string     `gorm:"<-:create"`
	TokenType string     `gorm:"<-:create"`
	TokenStr  string     `gorm:"<-:create"`
	ExpiresAt *time.Time `gorm:"<-:create"`
	CreatedAt time.Time  `gorm:"<-:create"`
	UpdatedAt time.Time
}

func NewTokenStore(token *Token) *StoreToken {
	return &StoreToken{UUID: token.UUID, TokenType: token.TokenType, Username: token.Username}
}

func (c *StoreToken) WithTokenStr(tokenStr string) *StoreToken {
	c.TokenStr = tokenStr
	return c
}

func (c *StoreToken) WithExpiresAt(expiresAt *int64) *StoreToken {
	if expiresAt == nil {
		return c
	}
	c.ExpiresAt = xutil.PtrTo[time.Time](time.Unix(*expiresAt, 0).UTC())
	return c
}
