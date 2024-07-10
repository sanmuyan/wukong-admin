package model

import "time"

// MFAApp 数据库对象
type MFAApp struct {
	// ID 唯一标识
	ID int
	// UserID 用户 ID
	UserID int `gorm:"<-:create"`
	// TOTPSecret TOTP 算法密钥
	TOTPSecret string    `gorm:"<-:create"`
	CreatedAt  time.Time `gorm:"<-:create"`
	UpdatedAt  time.Time
}

// MFAAppBindSession 会话数据
type MFAAppBindSession struct {
	TOTPSecret string `json:"totp_secret"`
}

// MFAAppBindRequest 请求对象
type MFAAppBindRequest struct {
	SessionID  string `json:"session_id" binding:"required"`
	TOTPSecret string `json:"totp_secret" binding:"required"`
	TOTPCode   string `json:"totp_code"  binding:"required"`
}

// MFAAppBindResponse 返回对象
type MFAAppBindResponse struct {
	SessionID  string `json:"session_id"`
	TOTPSecret string `json:"totp_secret"`
	QRCodeURI  string `json:"qr_code_uri"`
	TimeoutMin int    `json:"timeout_min"`
}
