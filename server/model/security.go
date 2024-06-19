package model

import "time"

const (
	RSAPurposeClientEncrypt = "client_encrypt"
)

type ModifyPasswordRequest struct {
	NewPassword string `json:"new_password" binding:"required"`
}

type LoginSecurity struct {
	ID             int        `json:"id"`
	UserID         int        `json:"user_id"`
	LastLoginAt    *time.Time `json:"last_login_at"`
	LoginFailCount *int       `json:"login_fail"`
	LockAt         *time.Time `json:"lock_at"`
	CreatedAt      time.Time  `json:"created_at"`
	UpdatedAt      time.Time  `json:"updated_at"`
}

type Cert struct {
	ID         int
	PrivateKey string
	PublicKey  string
	Purpose    string
	CreatedAt  time.Time
	UpdatedAt  time.Time
}
