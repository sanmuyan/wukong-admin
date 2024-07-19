package model

import (
	"time"
)

const (
	UserGitlabIDField   = "gitlab_id"
	UserWecomIDField    = "wecom_id"
	UserDingtalkIDField = "dingtalk_id"
)

// User 数据库对象，用于存储用户
type User struct {
	// ID 唯一标识
	ID int `json:"id"`
	// Username 用户名，不允许重复
	Username string `json:"username" gorm:"<-:create"`
	// DisplayName 用户显示名
	DisplayName string `json:"display_name,omitempty"`
	// Password 用户密码
	Password string `json:"password,omitempty"`
	// Mobile 用户手机号
	Mobile string `json:"mobile,omitempty"`
	// Email 用户邮箱
	Email string `json:"email,omitempty"`
	// Source 用户来源
	Source string `json:"source,omitempty"`
	// IsActive 用户是否激活，1 为激活，-1 为未激活
	IsActive int `json:"is_active"`
	// GitlabID GitLab 绑定
	GitlabID *string `json:"gitlab_id,omitempty"`
	// WechatID 企业微信绑定
	WecomID *string `json:"wecom_id,omitempty"`
	// DingtalkID 钉钉绑定
	DingtalkID *string   `json:"dingtalk_id,omitempty"`
	CreatedAt  time.Time `json:"created_at,omitempty" gorm:"<-:create"`
	UpdatedAt  time.Time `json:"updated_at,omitempty"`
}
