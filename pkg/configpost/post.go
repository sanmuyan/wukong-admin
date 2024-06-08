package configpost

import (
	"wukong/pkg/config"
	"wukong/pkg/datastore"
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
	datastore.InitStore()
	usersource.InitUserSource()
	controller.RunServer(config.Conf.ServerBind)
}
