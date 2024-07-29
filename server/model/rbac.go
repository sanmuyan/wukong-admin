package model

import "time"

// Resource 数据库对象，用于存储 API 鉴权资源
type Resource struct {
	// ID 唯一标识
	ID int `json:"id"`
	// ResourcePath 资源路径
	ResourcePath string `json:"resource_path"`
	// IsAuth 是否需要鉴权
	IsAuth int `json:"is_auth"`
	// Comment 备注
	Comment   string    `json:"comment,omitempty"`
	CreatedAt time.Time `json:"created_at,omitempty" gorm:"<-:create"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
}

// Role 数据库对象，用于存储权限角色
type Role struct {
	// ID 唯一标识
	ID int `json:"id"`
	// RoleName 角色名称，不能重复
	RoleName string `json:"role_name"`
	// AccessLevel 权限等级，0-100
	AccessLevel int `json:"access_level"`
	// UserMenus 用户菜单，多个菜单用逗号隔开
	UserMenus string `json:"user_menus"`
	// Comment 备注
	Comment   string    `json:"comment"`
	CreatedAt time.Time `json:"created_at,omitempty" gorm:"<-:create"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
}

// RoleResourceBind 数据库对象，用于存储角色资源绑定
type RoleResourceBind struct {
	// ID 唯一标识
	ID int `json:"id"`
	// API 资源 ID
	ResourceID int `json:"resource_id" gorm:"<-:create"`
	// 角色 ID
	RoleID    int       `json:"role_id" gorm:"<-:create"`
	CreatedAt time.Time `json:"created_at,omitempty" gorm:"<-:create"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
}

// UserRoleBind 数据库对象，用于存储用户权限绑定
type UserRoleBind struct {
	// ID 唯一标识
	ID int `json:"id"`
	// 用户 ID
	UserID int `json:"user_id" gorm:"<-:create"`
	// 角色 ID
	RoleID    int       `json:"role_id" gorm:"<-:create"`
	CreatedAt time.Time `json:"created_at,omitempty" gorm:"<-:create"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
}

// GroupRoleBind 数据库对象，用于存储用户组权限绑定
type GroupRoleBind struct {
	// ID 唯一标识
	ID int `json:"id"`
	// 用户 ID
	GroupID int `json:"group_id" gorm:"<-:create"`
	// 角色 ID
	RoleID    int       `json:"role_id" gorm:"<-:create"`
	CreatedAt time.Time `json:"created_at,omitempty" gorm:"<-:create"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
}
