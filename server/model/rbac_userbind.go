package model

import "time"

type UserBind struct {
	ID        int       `json:"id"`
	UserID    int       `json:"user_id" gorm:"<-:create"`
	RoleID    int       `json:"role_id" gorm:"<-:create"`
	CreatedAt time.Time `json:"created_at,omitempty" gorm:"<-:create"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
}
