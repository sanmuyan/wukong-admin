package configpost

import (
	"context"
	"encoding/json"
	"github.com/sanmuyan/xpkg/xcrypto"
	"github.com/sanmuyan/xpkg/xutil"
	"github.com/sirupsen/logrus"
	"time"
	"wukong/pkg/config"
	"wukong/pkg/datastore"
	"wukong/pkg/db"
	"wukong/pkg/oauthlogin"
	"wukong/pkg/passkey"
	"wukong/pkg/security"
	"wukong/pkg/usersource"
	"wukong/server/controller"
	"wukong/server/model"
)

// PostInit 配置初始化后要执行的操作
func PostInit(ctx context.Context) {
	// 解密加密配置项
	err := xcrypto.DecryptCFBToStruct(&config.Conf.Secret, config.Conf.ConfigSecretKey)
	if err != nil {
		logrus.Fatalf("config secret decrypt error: %s", err.Error())
	}
	// 初始化数据库
	db.InitMysql()
	if config.Conf.Database.Redis != "" {
		db.InitRedis()
	}
	datastore.InitStore(ctx)
	// 第一次从数据库加载系统配置
	err = loadDBConfig()
	if err != nil {
		logrus.Fatalf("config unmarshal error: %s", err.Error())
	}
	logrus.Debugf("config %+v", config.Conf)
	// 动态加载配置
	go loadDBConfigTask(ctx)
	// 启动 WEB 服务
	controller.RunServer(ctx, config.Conf.ServerBind)
}

// initDynamicConfig 初始化功能模块
func initDynamicConfig() {
	oauthlogin.InitOauthProviderConfig()
	usersource.InitUserSource()
	passkey.InitWebAuthnConfig()
	security.InitPrivateKeys()
}

var lastConfigUpdateAt int64

func loadDBConfig() error {
	dbConfig := &model.Config{
		Name:    "config",
		Content: string(xutil.RemoveError(json.Marshal(config.Conf))),
	}
	db.DB.Where("name = ?", "config").FirstOrCreate(dbConfig)
	if dbConfig.UpdatedAt.Unix() > lastConfigUpdateAt {
		logrus.Debugf("load config for db")
		err := json.Unmarshal([]byte(dbConfig.Content), &config.Conf)
		if err != nil {
			return err
		}
		initDynamicConfig()
	}
	lastConfigUpdateAt = dbConfig.UpdatedAt.Unix()
	return nil
}

func loadDBConfigTask(ctx context.Context) {
	logrus.Info("start load config task")
	ticker := time.NewTicker(5 * time.Second)
	for {
		select {
		case <-ctx.Done():
			ticker.Stop()
		case <-ticker.C:
			err := loadDBConfig()
			if err != nil {
				logrus.Errorf("load config error: %s", err.Error())
				continue
			}
		}
	}
}
