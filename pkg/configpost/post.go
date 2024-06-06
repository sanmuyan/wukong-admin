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
	db.InitMysql()
	datastorage.InitTokenStorage()
	usersource.InitUserSource()
	controller.RunServer(config.Conf.ServerBind)
}
