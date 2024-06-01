package configpost

import (
	"wukong/pkg/config"
	"wukong/pkg/db"
	"wukong/pkg/tokenclient"
	"wukong/pkg/userauth"
	"wukong/server/controller"
)

func PostInit() {
	if config.Conf.Database.Redis != "" {
		db.InitRedis()
	}
	db.InitMysql()
	tokenclient.InitTokenClient()
	userauth.InitUserAuth()
	controller.RunServer(config.Conf.ServerBind)
}
