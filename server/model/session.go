package model

import (
	"encoding/json"
	"github.com/sanmuyan/xpkg/xutil"
	"time"
)

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

type Session struct {
	SessionID   string     `json:"session_id"`
	SessionType string     `json:"session_type"`
	UserID      int        `json:"user_id"`
	Username    string     `json:"username"`
	SessionRaw  string     `json:"session_raw"`
	ExpiresAt   *time.Time `json:"expires_at"`
	CreatedAt   time.Time  `json:"created_at"`
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
