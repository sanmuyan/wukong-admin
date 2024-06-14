package config

type Database struct {
	Mysql string `mapstructure:"mysql"`
	Redis string `mapstructure:"redis"`
}

type Secret struct {
	TokenKey string `mapstructure:"token_key"`
}

type AttributeMap struct {
	DisplayName string `mapstructure:"display_name"`
	Email       string `mapstructure:"email"`
	Mobile      string `mapstructure:"mobile"`
}

type LDAP struct {
	URL               string       `mapstructure:"url"`
	SearchBase        string       `mapstructure:"search_base"`
	AdminDN           string       `mapstructure:"admin_dn"`
	AdminPassword     string       `mapstructure:"admin_password"`
	UsernameAttribute string       `mapstructure:"username_attribute"`
	AttributeMap      AttributeMap `mapstructure:"attribute_map"`
	SearchFilter      string       `mapstructure:"search_filter"`
}

type OauthProvider struct {
	ClientID     string   `mapstructure:"client_id"`
	ClientSecret string   `mapstructure:"client_secret"`
	RedirectURL  string   `mapstructure:"redirect_url"`
	AuthURL      string   `mapstructure:"auth_url"`
	TokenURL     string   `mapstructure:"token_url"`
	Scopes       []string `mapstructure:"scopes"`
	UserInfoURL  string   `mapstructure:"user_info_url"`
	Provider     string   `mapstructure:"provider"`
}

type App struct {
	OauthLoginRedirectURL string `mapstructure:"oauth_login_redirect_url"`
}

type WebAuthn struct {
	RPID      string   `mapstructure:"rp_id"`
	RPOrigins []string `mapstructure:"rp_origins"`
}

type Config struct {
	Database                 Database        `mapstructure:"database"`
	TokenTTL                 int             `mapstructure:"token_ttl"`
	Secret                   Secret          `mapstructure:"secret" `
	LogLevel                 int             `mapstructure:"log_level"`
	ServerBind               string          `mapstructure:"server_bind"`
	ConfigSecretKey          string          `mapstructure:"config_secret_key"`
	DataStore                string          `mapstructure:"data_store"`
	DisableVerifyServerToken bool            `mapstructure:"disable_verify_server_token"`
	LDAP                     LDAP            `mapstructure:"ldap"`
	OauthProviders           []OauthProvider `mapstructure:"oauth_providers"`
	App                      App             `mapstructure:"app"`
	AppName                  string          `mapstructure:"app_name"`
	WebAuthn                 WebAuthn        `mapstructure:"web_authn"`
}

var Conf Config

var OauthProviders = make(map[string]OauthProvider)
