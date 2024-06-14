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
	ID         int       `json:"id"`
	UserID     int       `json:"user_id"`
	TOTPSecret string    `json:"totp_secret"`
	ExpireAt   time.Time `json:"expire_at"`
	CreatedAt  time.Time `json:"created_at,omitempty"`
	UpdatedAt  time.Time `json:"updated_at,omitempty"`
}

type MFAAppBindRequest struct {
	TOTPSecret string `json:"totp_secret" binding:"required"`
	TOTPCode   string `json:"totp_code"  binding:"required"`
}

type MFAAppBindResponse struct {
	TOTPSecret string `json:"totp_secret"`
	QRCodeURI  string `json:"qr_code_uri"`
	TimeoutMin int    `json:"timeout_min"`
}
