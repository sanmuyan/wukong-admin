package model

import "time"

// Resource 数据库对象，用于存储 API 鉴权资源
type Resource struct {
	// ID 唯一标识
	ID int `json:"id"`
	// ResourceName 资源名称
	ResourceName string `json:"resource_name"`
	// ResourcePath 资源路径
	ResourcePath string `json:"resource_path"`
	// IsAuth 是否需要鉴权
	IsAuth int `json:"is_auth"`
	// Comment 备注
	Comment   string    `json:"comment,omitempty"`
	CreatedAt time.Time `json:"created_at,omitempty" gorm:"<-:create"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
}
