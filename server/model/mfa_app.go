package model

import "time"

type MFAApp struct {
	ID         int
	UserID     int       `gorm:"<-:create"`
	TOTPSecret string    `gorm:"<-:create"`
	CreatedAt  time.Time `gorm:"<-:create"`
	UpdatedAt  time.Time
}

type MFAAppBindSession struct {
	TOTPSecret string `json:"totp_secret"`
}

type MFAAppBindRequest struct {
	SessionID  string `json:"session_id" binding:"required"`
	TOTPSecret string `json:"totp_secret" binding:"required"`
	TOTPCode   string `json:"totp_code"  binding:"required"`
}

type MFAAppBindResponse struct {
	SessionID  string `json:"session_id"`
	TOTPSecret string `json:"totp_secret"`
	QRCodeURI  string `json:"qr_code_uri"`
	TimeoutMin int    `json:"timeout_min"`
}
