package model

import "time"

type Resource struct {
	ID           int       `json:"id"`
	ResourcePath string    `json:"resource_path"`
	IsAuth       int       `json:"is_auth"`
	Comment      string    `json:"comment"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}
