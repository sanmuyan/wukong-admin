package model

import (
	"time"
)

// OauthApp 数据库对象，用于存储 OAuth 应用
type OauthApp struct {
	// ID 唯一标识
	ID int `json:"id"`
	// AppName 应用名称
	AppName string `json:"app_name"`
	// ClientID 客户端 ID，使用 UUID 生成
	ClientID string `json:"client_id" gorm:"<-:create"`
	// ClientSecret 客户端密钥，使用64位随机字符串生成
	ClientSecret string `json:"client_secret" gorm:"<-:create"`
	// 应用允许的权限范围，多个权限用逗号隔开
	Scope string `json:"scope,omitempty"`
	// 应用允许回调的地址，多个地址用逗号隔开
	RedirectURI string `json:"redirect_uri,omitempty"`
	// Comment 备注
	Comment   string    `json:"comment,omitempty"`
	CreatedAt time.Time `json:"created_at,omitempty" gorm:"<-:create"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
}
