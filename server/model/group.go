package model

import "time"

// Group 数据库对象，用于存储用户
type Group struct {
	// ID 唯一标识
	ID int `json:"id"`
	// GroupName  组名，不允许重复
	GroupName string `json:"group_name"`
	// DisplayName  组描述
	DisplayName string    `json:"display_name"`
	CreatedAt   time.Time `json:"created_at,omitempty" gorm:"<-:create"`
	UpdatedAt   time.Time `json:"updated_at,omitempty"`
}

// UserGroupBind 数据库对象，用于存储用户组关系
type UserGroupBind struct {
	// ID 唯一标识
	ID int `json:"id"`
	// UserID  用户 ID
	UserID int `json:"user_id" gorm:"<-:create"`
	// GroupID  组 ID
	GroupID   int       `json:"group_id" gorm:"<-:create"`
	CreatedAt time.Time `json:"created_at,omitempty" gorm:"<-:create"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
}
