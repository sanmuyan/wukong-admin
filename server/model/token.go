package model

import (
	"errors"
)

type Token struct {
	userID      int
	Username    string `json:"username" binding:"required"`
	AccessLevel int    `json:"access_level" binding:"required"`
	TokenType   string `json:"token_type" binding:"required"`
	ExpiresAt   *int64 `json:"exp,omitempty"`
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
	if t.Issuer != AppName {
		return err
	}
	if t.Username == "" {
		return err
	}
	if t.TokenType == "" {
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
	SessionToken = "session"
	ApiToken     = "api"
)

func init() {
	TokenTypes[SessionToken] = struct{}{}
	TokenTypes[ApiToken] = struct{}{}
}
