package model

import "time"

type RoleBind struct {
	ID         int       `json:"id"`
	ResourceID int       `json:"resource_id"`
	RoleID     int       `json:"role_id"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}
