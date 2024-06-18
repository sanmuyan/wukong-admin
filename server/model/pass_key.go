package model

import (
	"github.com/go-webauthn/webauthn/protocol"
	"github.com/go-webauthn/webauthn/webauthn"
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
	SessionData webauthn.SessionData
}

type PassKeyBeginRegistrationResponse struct {
	SessionID string                       `json:"session_id"`
	Options   *protocol.CredentialCreation `json:"options"`
}

type PassKeyFinishRegistrationRequest struct {
	SessionID string `form:"session_id" binding:"required"`
}

type PassKeyLoginSession struct {
	SessionData webauthn.SessionData
}

type PassKeyBeginLoginRequest struct {
	Username string `json:"username" binding:"required"`
}

type PassKeyBeginLoginResponse struct {
	SessionID string                        `json:"session_id"`
	Options   *protocol.CredentialAssertion `json:"options"`
}

type PassKeyFinishLoginRequest struct {
	SessionID string `form:"session_id" binding:"required"`
}
