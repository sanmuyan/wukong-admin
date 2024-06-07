package configpost

import (
	"wukong/pkg/config"
	"wukong/pkg/datastorage"
	"wukong/pkg/db"
	"wukong/pkg/usersource"
	"wukong/server/controller"
)

func PostInit() {
	if config.Conf.Database.Redis != "" {
		db.InitRedis()
	}
	for _, oauthProvider := range config.Conf.OauthProviders {
		config.OauthProviders[oauthProvider.Provider] = oauthProvider
	}
	db.InitMysql()
	datastorage.InitTokenStorage()
	usersource.InitUserSource()
	controller.RunServer(config.Conf.ServerBind)
}
