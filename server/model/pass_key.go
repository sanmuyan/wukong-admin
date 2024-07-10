package model

import (
	"github.com/go-webauthn/webauthn/protocol"
	"github.com/go-webauthn/webauthn/webauthn"
	"time"
)

// PassKey 数据库对象，用于存储通行密钥
type PassKey struct {
	ID            int        `json:"id"`
	UserID        int        `json:"user_id" gorm:"<-:create"`
	CredentialID  string     `json:"credential_id,omitempty" gorm:"<-:create"`
	CredentialRaw string     `json:"credential_raw,omitempty" gorm:"<-:create"`
	DisplayName   string     `json:"display_name"`
	LastUsedAt    *time.Time `json:"last_used_at,omitempty"`
	CreatedAt     time.Time  `json:"created_at,omitempty" gorm:"<-:create"`
	UpdatedAt     time.Time  `json:"updated_at,omitempty"`
}

// PassKeyRegisterSession 会话数据
type PassKeyRegisterSession struct {
	SessionData webauthn.SessionData
}

// PassKeyBeginRegistrationResponse 返回对象
type PassKeyBeginRegistrationResponse struct {
	SessionID string `json:"session_id"`
	// Options 注意 Base64url 编码问题，浏览器默认是 Base64 编码
	Options *protocol.CredentialCreation `json:"options"`
}

// PassKeyFinishRegistrationRequest 请求对象
type PassKeyFinishRegistrationRequest struct {
	SessionID string `form:"session_id" binding:"required"`
}

// PassKeyLoginSession 登录会话数据
type PassKeyLoginSession struct {
	SessionData webauthn.SessionData
}

// PassKeyBeginLoginRequest 请求对象
type PassKeyBeginLoginRequest struct {
	Username string `json:"username" binding:"required"`
}

// PassKeyBeginLoginResponse 返回对象
type PassKeyBeginLoginResponse struct {
	SessionID string `json:"session_id"`
	// Options 注意 Base64url 编码问题，浏览器默认是 Base64 编码
	Options *protocol.CredentialAssertion `json:"options"`
}

// PassKeyFinishLoginRequest 请求对象
type PassKeyFinishLoginRequest struct {
	SessionID string `form:"session_id" binding:"required"`
}
