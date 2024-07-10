package model

import (
	"wukong/pkg/page"
)

// Query 用于便捷查询数据
type Query struct {
	// QueryLikeValue 模糊查询的值
	QueryLikeValue string
	// QueryLikeKeys 允许模糊查询的字段，多个字段用逗号隔开
	QueryLikeKeys string
	// QueryMustMap 精确查询的值
	QueryMustMap map[string]any
	// Page 查询分页器
	Page *page.Page
	// QueryType 查询类型
	// 0 精确查询
	// 1 模糊查询
	QueryType int
}

// List 数据查询列表结构
type List interface {
	// GetData 获取数据
	GetData() any
	// GetPage 获取分页器
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

type OauthApps struct {
	OauthApps []OauthApp `json:"oauth_apps"`
	page.Page
}

func (r *OauthApps) GetData() any {
	return &r.OauthApps
}

func (r *OauthApps) GetPage() *page.Page {
	return &r.Page
}
