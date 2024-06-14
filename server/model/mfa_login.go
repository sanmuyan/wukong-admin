package model

import "time"

const (
	MFAProviderMFAApp = "mfa_app"
)

type MFALoginSession struct {
	ID          int
	UserID      int
	Username    string
	SessionID   string
	MFAProvider string
	ExpireAt    time.Time
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type MFAFinishLoginRequest struct {
	SessionID   string `json:"session_id" binding:"required"`
	MFAProvider string `json:"mfa_provider" binding:"required"`
	Username    string `json:"username" binding:"required"`
	Code        string `json:"code" binding:"required"`
}

type MFABeginLoginResponse struct {
	Username    string `json:"username"`
	SessionID   string `json:"session_id"`
	MFAProvider string `json:"mfa_provider"`
}
