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
	DisplayName string    `json:"display_name"`
	Password    string    `json:"password"`
	Mobile      string    `json:"mobile"`
	Email       string    `json:"email"`
	Source      string    `json:"source"`
	IsActive    int       `json:"is_active"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
