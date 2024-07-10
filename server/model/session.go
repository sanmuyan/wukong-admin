package model

import (
	"encoding/json"
	"github.com/sanmuyan/xpkg/xutil"
	"time"
)

// Session 的类型
const (
	SessionTypeMFAAppBind        = "mfa_app_bind"
	SessionTypeMFALogin          = "mfa_login"
	SessionTypePassKeyRegister   = "pass_key_register"
	SessionTypePassKeyLogin      = "pass_key_login"
	SessionTypeOAuthCode         = "oauth_code"
	SessionTypeSessionToken      = TokenTypeSession
	SessionTypeApiToken          = TokenTypeApi
	SessionTypeOauthAccessToken  = TokenTypeOauthAccess
	SessionTypeOauthRefreshToken = TokenTypeOauthRefresh
)

// Session 数据库对象，用于存储各类会话
type Session struct {
	// SessionID 唯一标识，使用 UUID 生成，用于命中用户身份信息
	SessionID string `json:"session_id"`
	// SessionType 类型
	SessionType string `json:"session_type"`
	// UserID 用户 ID
	UserID int `json:"user_id"`
	// Username 用户名
	Username string `json:"username"`
	// SessionRaw Session 原始内容
	SessionRaw string `json:"session_raw"`
	// ExpiresAt 过期时间
	ExpiresAt *time.Time `json:"expires_at"`
	CreatedAt time.Time  `json:"created_at"`
}

func NewSession(sessionID string, sessionType string, userID int, username string, sessionRaw any) *Session {
	return &Session{
		SessionID:   sessionID,
		SessionType: sessionType,
		UserID:      userID,
		Username:    username,
		SessionRaw:  string(xutil.RemoveError(json.Marshal(sessionRaw))),
		CreatedAt:   time.Now().UTC()}
}

func (s *Session) SetTimeout(d time.Duration) *Session {
	s.ExpiresAt = xutil.PtrTo(time.Now().UTC().Add(d))
	return s
}

func (s *Session) IsExpired() bool {
	if s.ExpiresAt == nil {
		return false
	}
	return s.ExpiresAt.Before(time.Now().UTC())
}

func (s *Session) UnmarshalSessionRaw(sessionRaw any) error {
	return json.Unmarshal([]byte(s.SessionRaw), sessionRaw)
}
