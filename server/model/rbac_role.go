package model

import "time"

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
