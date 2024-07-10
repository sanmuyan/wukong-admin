package model

import (
	"time"
)

// Config 数据库对象，用户存储系统配置
type Config struct {
	// ID 唯一标识
	ID int `json:"id"`
	// 配置名称
	Name string `json:"name"`
	// 配置内容，应为 JSON 字符串
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at,omitempty" gorm:"<-:create"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
}
