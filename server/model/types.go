package model

import (
	"wukong/pkg/page"
)

type Users struct {
	Users []User `json:"users"`
	page.Page
}

type UserBinds struct {
	UserBinds []UserBind `json:"user_binds"`
	page.Page
}

type Roles struct {
	Roles []Role `json:"roles"`
	page.Page
}

type RoleBinds struct {
	RoleBinds []RoleBind `json:"role_binds"`
	page.Page
}

type Resources struct {
	Resources []Resource `json:"resources"`
	page.Page
}

type Query struct {
	QueryLikeValue string
	QueryLikeKeys  string
	QueryMustMap   map[string]any
	PageNumber     int
	PageSize       int
}
