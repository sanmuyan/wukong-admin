package model

import (
	"time"
)

type Config struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at,omitempty" gorm:"<-:create"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
}
