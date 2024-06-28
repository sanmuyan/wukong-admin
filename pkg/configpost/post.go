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

func PostInit(ctx context.Context) {
	err := xcrypto.DecryptCFBToStruct(&config.Conf.Secret, config.Conf.ConfigSecretKey)
	if err != nil {
		logrus.Fatalf("config secret decrypt error: %s", err.Error())
	}
	db.InitMysql()
	datastore.InitStore()
	err = loadDBConfig()
	if err != nil {
		logrus.Fatalf("config unmarshal error: %s", err.Error())
	}
	logrus.Debugf("config %+v", config.Conf)
	if config.Conf.Database.Redis != "" {
		db.InitRedis()
	}
	go loadDBConfigTimer(ctx)
	controller.RunServer(ctx, config.Conf.ServerBind)
}

func initDynamicConfig() {
	oauthlogin.InitOauthProviders()
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

func loadDBConfigTimer(ctx context.Context) {
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
