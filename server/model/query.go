package model

import (
	"wukong/pkg/page"
)

type Query struct {
	QueryLikeValue string
	QueryLikeKeys  string
	QueryMustMap   map[string]any
	Page           *page.Page
	QueryType      int
}

type List interface {
	GetData() any
	GetPage() *page.Page
}

type Users struct {
	Users []User `json:"users"`
	page.Page
}

func (r *Users) GetData() any {
	return &r.Users
}

func (r *Users) GetPage() *page.Page {
	return &r.Page
}

type UserBinds struct {
	UserBinds []UserBind `json:"user_binds"`
	page.Page
}

func (r *UserBinds) GetData() any {
	return &r.UserBinds
}

func (r *UserBinds) GetPage() *page.Page {
	return &r.Page
}

type Roles struct {
	Roles []Role `json:"roles"`
	page.Page
}

func (r *Roles) GetData() any {
	return &r.Roles
}

func (r *Roles) GetPage() *page.Page {
	return &r.Page
}

type RoleBinds struct {
	RoleBinds []RoleBind `json:"role_binds"`
	page.Page
}

func (r *RoleBinds) GetData() any {
	return &r.RoleBinds
}

func (r *RoleBinds) GetPage() *page.Page {
	return &r.Page
}

type Resources struct {
	Resources []Resource `json:"resources"`
	page.Page
}

func (r *Resources) GetData() any {
	return &r.Resources
}

func (r *Resources) GetPage() *page.Page {
	return &r.Page
}

type OauthAPPS struct {
	OauthAPPS []OauthAPP `json:"oauth_apps"`
	page.Page
}

func (r *OauthAPPS) GetData() any {
	return &r.OauthAPPS
}

func (r *OauthAPPS) GetPage() *page.Page {
	return &r.Page
}
