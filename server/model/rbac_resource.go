package model

import "time"

type Resource struct {
	ID           int       `json:"id"`
	ResourcePath string    `json:"resource_path"`
	IsAuth       int       `json:"is_auth"`
	Comment      string    `json:"comment,omitempty"`
	CreatedAt    time.Time `json:"created_at,omitempty" gorm:"<-:create"`
	UpdatedAt    time.Time `json:"updated_at,omitempty"`
}
