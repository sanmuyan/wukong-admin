package model

import (
	"github.com/go-webauthn/webauthn/protocol"
	"time"
)

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

type PassKeyRegisterSession struct {
	ID         int
	UserID     int
	SessionRaw string
	ExpireAt   time.Time
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

type PassKeyLoginSession struct {
	ID         int
	UserID     int
	Username   string
	SessionID  string
	SessionRaw string
	ExpireAt   time.Time
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

type PassKeyBeginLoginRequest struct {
	Username string `json:"username" binding:"required"`
}

type PassKeyBeginLoginResponse struct {
	Username  string                        `json:"username"`
	SessionID string                        `json:"session_id"`
	Options   *protocol.CredentialAssertion `json:"options"`
}

type PassKeyFinishLoginRequest struct {
	Username  string `form:"username" binding:"required"`
	SessionID string `form:"session_id" binding:"required"`
}
