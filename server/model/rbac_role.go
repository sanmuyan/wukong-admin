package model

import "time"

type Role struct {
	ID          int       `json:"id"`
	RoleName    string    `json:"role_name"`
	AccessLevel int       `json:"access_level"`
	UserMenus   string    `json:"user_menus"`
	Comment     string    `json:"comment"`
	CreatedAt   time.Time `json:"created_at,omitempty" gorm:"<-:create"`
	UpdatedAt   time.Time `json:"updated_at,omitempty"`
}
