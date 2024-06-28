package config

type Database struct {
	Mysql string `mapstructure:"mysql"`
	Redis string `mapstructure:"redis"`
}

type Secret struct {
	TokenKey string `mapstructure:"token_key"`
}

type AttributeMap struct {
	DisplayName string `mapstructure:"display_name" json:"display_name,omitempty"`
	Email       string `mapstructure:"email" json:"email,omitempty"`
	Mobile      string `mapstructure:"mobile" json:"mobile,omitempty"`
}

type LDAP struct {
	Enable            bool         `mapstructure:"enable" json:"enable" binding:"required"`
	URL               string       `mapstructure:"url" json:"url,omitempty" binding:"required"`
	SearchBase        string       `mapstructure:"search_base" json:"search_base,omitempty" binding:"required"`
	AdminDN           string       `mapstructure:"admin_dn" json:"admin_dn,omitempty"`
	AdminPassword     string       `mapstructure:"admin_password" json:"admin_password,omitempty"`
	UsernameAttribute string       `mapstructure:"username_attribute" json:"username_attribute,omitempty" binding:"required"`
	AttributeMap      AttributeMap `mapstructure:"attribute_map" json:"attribute_map" binding:"required"`
	SearchFilter      string       `mapstructure:"search_filter" json:"search_filter,omitempty" binding:"required"`
}

type OauthProvider struct {
	Enable       bool     `mapstructure:"enable" json:"enable"`
	ClientID     string   `mapstructure:"client_id" json:"client_id,omitempty" binding:"required"`
	ClientSecret string   `mapstructure:"client_secret" json:"client_secret,omitempty"`
	RedirectURL  string   `mapstructure:"redirect_url" json:"redirect_url,omitempty" binding:"required"`
	AuthURL      string   `mapstructure:"auth_url" json:"auth_url,omitempty" binding:"required"`
	TokenURL     string   `mapstructure:"token_url" json:"token_url,omitempty" binding:"required"`
	Scopes       []string `mapstructure:"scopes" json:"scopes,omitempty" binding:"required"`
	UserInfoURL  string   `mapstructure:"user_info_url" json:"user_info_url,omitempty" binding:"required"`
	Provider     string   `mapstructure:"provider" json:"provider,omitempty"  binding:"required"`
}

type Security struct {
	TokenTTL                 int  `mapstructure:"token_ttl" json:"token_ttl,omitempty"  binding:"required"`
	DisableVerifyServerToken bool `mapstructure:"disable_verify_server_token" json:"disable_verify_server_token,omitempty" binding:"required"`
	LoginMaxFails            int  `mapstructure:"login_max_fails" json:"login_max_fails,omitempty" binding:"required"`
	LoginLockTime            int  `mapstructure:"login_lock_time" json:"login_lock_time,omitempty" binding:"required"`
	PasswordMinLength        int  `mapstructure:"password_min_length" json:"password_min_length,omitempty" binding:"required"`
	PasswordComplexity       int  `mapstructure:"password_complexity" json:"password_complexity,omitempty" binding:"required"`
	PassKeyLogin             bool `mapstructure:"pass_key_login" json:"pass_key_login,omitempty"`
	RequireMFA               bool `mapstructure:"require_mfa" json:"require_mfa,omitempty"`
}

type Basic struct {
	AppName  string `mapstructure:"app_name" json:"app_name,omitempty" binding:"required"`
	SiteURL  string `mapstructure:"site_url" json:"site_url,omitempty" binding:"required"`
	SiteHost string `mapstructure:"site_host" json:"site_host,omitempty" binding:"required"`
}

type Config struct {
	Database        Database        `mapstructure:"database" json:"-"`
	Secret          Secret          `mapstructure:"secret" json:"-"`
	LogLevel        int             `mapstructure:"log_level" json:"-"`
	ServerBind      string          `mapstructure:"server_bind" json:"-"`
	ConfigSecretKey string          `mapstructure:"config_secret_key" json:"-"`
	DataStore       string          `mapstructure:"data_store" json:"-"`
	Security        Security        `mapstructure:"security" json:"security,omitempty"`
	LDAP            LDAP            `mapstructure:"ldap" json:"ldap,omitempty"`
	OauthProviders  []OauthProvider `mapstructure:"oauth_providers" json:"oauth_providers,omitempty"`
	Basic           Basic           `mapstructure:"basic" json:"basic,omitempty"`
}

var Conf Config
