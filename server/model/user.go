package model

import (
	"time"
)

type Login struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type User struct {
	ID          int       `json:"id"`
	Username    string    `json:"username" gorm:"<-:create"`
	DisplayName string    `json:"display_name,omitempty"`
	Password    string    `json:"password,omitempty"`
	Mobile      string    `json:"mobile,omitempty"`
	Email       string    `json:"email,omitempty"`
	Source      string    `json:"source,omitempty"`
	IsActive    int       `json:"is_active"`
	CreatedAt   time.Time `json:"created_at,omitempty" gorm:"<-:create"`
	UpdatedAt   time.Time `json:"updated_at,omitempty"`
}
