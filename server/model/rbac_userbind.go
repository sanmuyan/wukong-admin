package model

import "time"

// UserBind 数据库对象，用于存储用户权限绑定
type UserBind struct {
	// ID 唯一标识
	ID int `json:"id"`
	// 用户 ID
	UserID int `json:"user_id" gorm:"<-:create"`
	// 角色 ID
	RoleID    int       `json:"role_id" gorm:"<-:create"`
	CreatedAt time.Time `json:"created_at,omitempty" gorm:"<-:create"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
}
