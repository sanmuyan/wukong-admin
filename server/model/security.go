package model

import "time"

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
