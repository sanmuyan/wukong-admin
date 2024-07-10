package model

import "time"

const (
	// RSAPurposeClientEncrypt 用于客户端传输数据加密
	RSAPurposeClientEncrypt = "client_encrypt"
)

type ModifyPasswordRequest struct {
	NewPassword string `json:"new_password" binding:"required"`
}

// LoginSecurity 数据库对象，用于用户登录安全
type LoginSecurity struct {
	// 唯一标识
	ID int `json:"id"`
	// 用户 ID
	UserID int `json:"user_id"`
	// 最后登录时间
	LastLoginAt *time.Time `json:"last_login_at"`
	// 用户登录失败阈值
	LoginFailCount *int `json:"login_fail"`
	// 登录失败锁定时间
	LockAt    *time.Time `json:"lock_at"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
}

// Cert 数据库对象，用于存储系统证书
type Cert struct {
	// ID 唯一标识
	ID int
	// PrivateKey RSA 私钥，使用 base64 编码存储
	PrivateKey string
	// PublicKey RSA 公钥，使用 base64 编码存储
	PublicKey string
	// Purpose 证书用途
	Purpose   string
	CreatedAt time.Time
	UpdatedAt time.Time
}
