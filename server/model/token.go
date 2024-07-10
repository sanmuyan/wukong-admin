package model

import (
	"errors"
	"fmt"
	"github.com/sanmuyan/xpkg/xutil"
	"wukong/pkg/config"
)

// Token 用户令牌
type Token struct {
	// userID 用户 ID 仅供系统内部使用
	userID int
	// TokenID 唯一标识，使用 UUID 生成
	TokenID string `json:"token_id"`
	// Username 用户名
	Username string `json:"username" binding:"required"`
	// AccessLevel 访问级别
	AccessLevel int `json:"access_level,omitempty"`
	// TokenType Token 类型
	TokenType string `json:"token_type" binding:"required"`
	// ExpiresAt 过期时间
	ExpiresAt *int64 `json:"exp,omitempty"`
	// Scope OAuth 应用授权范围
	Scope string `json:"scope,omitempty"`
	// ClientID OAuth 应用授权客户端
	ClientID string `json:"client_id,omitempty"`
	// IssuedAt Token 发行时间
	IssuedAt int64 `json:"iat"`
	// Issuer Token 发行者
	Issuer string `json:"iss"`
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
	// TokenTypeSession 用户登录会话
	TokenTypeSession = "session_token"
	// TokenTypeApi 用户调用 API
	TokenTypeApi = "api_token"
	// TokenTypeOauthAccess OAuth 2.0 访问令牌
	TokenTypeOauthAccess = "oauth_access_token"
	// TokenTypeOauthRefresh OAuth 2.0 刷新令牌
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
