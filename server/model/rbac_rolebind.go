package model

import "time"

// RoleBind 数据库对象，用于存储角色资源绑定
type RoleBind struct {
	// ID 唯一标识
	ID int `json:"id"`
	// API 资源 ID
	ResourceID int `json:"resource_id" gorm:"<-:create"`
	// 角色 ID
	RoleID    int       `json:"role_id" gorm:"<-:create"`
	CreatedAt time.Time `json:"created_at,omitempty" gorm:"<-:create"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
}
