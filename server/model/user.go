package model

import (
	"time"
)

// User 数据库对象，用于存储用户
type User struct {
	// ID 唯一标识
	ID int `json:"id"`
	// Username 用户名，不允许重复
	Username string `json:"username" gorm:"<-:create"`
	// 用户显示名
	DisplayName string `json:"display_name,omitempty"`
	// 用户密码
	Password string `json:"password,omitempty"`
	// 用户手机号
	Mobile string `json:"mobile,omitempty"`
	// 用户邮箱
	Email string `json:"email,omitempty"`
	// 用户来源
	Source string `json:"source,omitempty"`
	// 用户是否激活，1 为激活，-1 为未激活
	IsActive  int       `json:"is_active"`
	CreatedAt time.Time `json:"created_at,omitempty" gorm:"<-:create"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
}
