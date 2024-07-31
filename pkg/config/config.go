package config

// Database 数据库配置
type Database struct {
	// MySQL 系统基础数据库
	MySQL string `mapstructure:"mysql"`
	// Redis 用于保存临时数据
	Redis string `mapstructure:"redis"`
}

// Secret 成员值应使用 CFB 加密
type Secret struct {
	// Token JWT 密钥
	TokenKey string `mapstructure:"token_key"`
	// SessionKey Session hmac 密钥
	SessionKey string `mapstructure:"session_key"`
}

// AttributeMap 用于映射 LDAP 属性
type AttributeMap struct {
	// DisplayName 用户显示名
	DisplayName string `mapstructure:"display_name" json:"display_name,omitempty"`
	// Email 用户邮箱
	Email string `mapstructure:"email" json:"email,omitempty"`
	// Mobile 用户手机号
	Mobile string `mapstructure:"mobile" json:"mobile,omitempty"`
}

// LDAP 登录配置
type LDAP struct {
	// Enable 是否启用 LDAP 登录
	Enable bool `mapstructure:"enable" json:"enable" binding:"required"`
	// URL 服务器地址
	URL string `mapstructure:"url" json:"url,omitempty" binding:"required"`
	// SearchBase 用户搜索基点
	SearchBase string `mapstructure:"search_base" json:"search_base,omitempty" binding:"required"`
	// AdminDN 管理员 DN
	AdminDN string `mapstructure:"admin_dn" json:"admin_dn,omitempty"`
	// AdminPassword 管理员密码
	AdminPassword string `mapstructure:"admin_password" json:"admin_password,omitempty"`
	// UsernameAttribute 用户名属性
	UsernameAttribute string `mapstructure:"username_attribute" json:"username_attribute,omitempty" binding:"required"`
	// AttributeMap 用户属性映射
	AttributeMap AttributeMap `mapstructure:"attribute_map" json:"attribute_map" binding:"required"`
	// SearchFilter 用户过滤
	SearchFilter string `mapstructure:"search_filter" json:"search_filter,omitempty" binding:"required"`
}

// OauthProvider 第三方 OAuth 配置
type OauthProvider struct {
	// Enable 是否启用
	Enable bool `mapstructure:"enable" json:"enable"`
	// CorpID 企业微信 CorpID
	CorpID string `mapstructure:"corp_id" json:"corp_id,omitempty"`
	// ClientID 客户端 ID
	ClientID string `mapstructure:"client_id" json:"client_id,omitempty" binding:"required"`
	// ClientSecret 客户端密钥
	ClientSecret string `mapstructure:"client_secret" json:"client_secret,omitempty" binding:"required"`
	// RedirectURL 回调地址
	RedirectURL string `mapstructure:"redirect_url" json:"redirect_url,omitempty" binding:"required"`
	// AuthURL 授权接口
	AuthURL string `mapstructure:"auth_url" json:"auth_url,omitempty" binding:"required"`
	// TokenURL 获取 Token 接口
	TokenURL string `mapstructure:"token_url" json:"token_url,omitempty" binding:"required"`
	// Scopes 权限范围
	Scopes string `mapstructure:"scopes" json:"scopes,omitempty"`
	// UserInfoURL 获取用户信息接口
	UserInfoURL string `mapstructure:"user_info_url" json:"user_info_url,omitempty" binding:"required"`
}

// OauthProviders 第三方登录配置
type OauthProviders struct {
	Gitlab   *OauthProvider `mapstructure:"gitlab" json:"gitlab,omitempty"`
	Wecom    *OauthProvider `mapstructure:"we_com" json:"we_com,omitempty"`
	Dingtalk *OauthProvider `mapstructure:"ding_talk" json:"ding_talk,omitempty"`
}

// Security 安全配置
type Security struct {
	// TokenTTL Token 过期时间，单位秒
	TokenTTL int `mapstructure:"token_ttl" json:"token_ttl,omitempty"  binding:"required"`
	// VerifyTokenSession 验证令牌会话是否有效
	VerifyTokenSession bool `mapstructure:"verify_token_session" json:"verify_token_session"`
	// LoginMaxFails 用户登录失败多少次后锁定
	LoginMaxFails int `mapstructure:"login_max_fails" json:"login_max_fails,omitempty" binding:"required"`
	// LoginLockTime 登录失败阈值锁定时间
	LoginLockTime int `mapstructure:"login_lock_time" json:"login_lock_time,omitempty" binding:"required"`
	// PasswordMinLength 密码最小长度
	PasswordMinLength int `mapstructure:"password_min_length" json:"password_min_length,omitempty" binding:"required"`
	// PasswordComplexity 密码复杂度
	PasswordComplexity int `mapstructure:"password_complexity" json:"password_complexity,omitempty" binding:"required"`
	// PassKeyLogin 是否启用通行密钥登录
	PassKeyLogin bool `mapstructure:"pass_key_login" json:"pass_key_login,omitempty"`
	// RequireMFA 是否要求用户开启多因素认证
	RequireMFA bool `mapstructure:"require_mfa" json:"require_mfa,omitempty"`
}

// Basic 系统基础设置
type Basic struct {
	// AppName 应用名称
	AppName string `mapstructure:"app_name" json:"app_name,omitempty" binding:"required"`
	// SiteURL 网站 URL
	SiteURL string `mapstructure:"site_url" json:"site_url,omitempty" binding:"required"`
	// SiteHost 网站域名
	SiteHost string `mapstructure:"site_host" json:"site_host,omitempty" binding:"required"`
}

// Config 全局配置
type Config struct {
	// Database 数据库配置
	Database Database `mapstructure:"database" json:"-"`
	// Secret 安全配置
	Secret Secret `mapstructure:"secret" json:"-"`
	// LogLevel 日志配置，0-6
	LogLevel int `mapstructure:"log_level" json:"-"`
	// ServerBind 服务绑定地址
	ServerBind string `mapstructure:"server_bind" json:"-"`
	// ConfigSecretKey Secret 成员 CFB 解密密钥
	ConfigSecretKey string `mapstructure:"config_secret_key" json:"-"`
	// DataStore 数据存储类型，一般用于存储临时数据
	DataStore string `mapstructure:"data_store" json:"-"`
	// Security 安全配置
	Security Security `mapstructure:"security" json:"security,omitempty"`
	// LDAP 登录配置
	LDAP LDAP `mapstructure:"ldap" json:"ldap,omitempty"`
	// OauthProviders 第三方登录配置
	OauthProviders OauthProviders `mapstructure:"oauth_providers" json:"oauth_providers,omitempty"`
	// Basic 系统基础设置
	Basic Basic `mapstructure:"basic" json:"basic,omitempty"`
}

var Conf Config
