package config

type Database struct {
	Mysql string `mapstructure:"mysql"`
	Redis string `mapstructure:"redis"`
}

type Secret struct {
	TokenKey string `mapstructure:"token_key"`
}

type Config struct {
	Database        Database `mapstructure:"database"`
	TokenTTL        int      `mapstructure:"token_ttl"`
	Secret          Secret   `mapstructure:"secret" `
	LogLevel        int      `mapstructure:"log_level"`
	ServerBind      string   `mapstructure:"server_bind"`
	ConfigSecretKey string   `mapstructure:"config_secret_key"`
}

var Conf Config
