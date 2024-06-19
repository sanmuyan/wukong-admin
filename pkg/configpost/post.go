package configpost

import (
	"wukong/pkg/config"
	"wukong/pkg/datastore"
	"wukong/pkg/db"
	"wukong/pkg/passkey"
	"wukong/pkg/security"
	"wukong/pkg/usersource"
	"wukong/server/controller"
)

func PostInit() {
	if config.Conf.Database.Redis != "" {
		db.InitRedis()
	}
	db.InitMysql()
	initProviders()
	controller.RunServer(config.Conf.ServerBind)
}

func initProviders() {
	for _, oauthProvider := range config.Conf.OauthProviders {
		config.OauthProviders[oauthProvider.Provider] = oauthProvider
	}
	datastore.InitStore()
	usersource.InitUserSource()
	passkey.InitWebAuthnConfig(config.Conf.AppName, config.Conf.WebAuthn.RPID, config.Conf.WebAuthn.RPOrigins)
	security.InitSecurity()
}
