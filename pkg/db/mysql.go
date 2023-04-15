package db

import (
	"github.com/sirupsen/logrus"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"time"
	"wukong/pkg/config"
)

var DB *gorm.DB

func InitMysql() {
	var logLevel logger.LogLevel
	var err error
	logLevel = logger.LogLevel(config.Conf.LogLevel - 1)
	dsn := config.Conf.Database.Mysql
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logLevel),
	})
	if err != nil {
		logrus.Fatal(err)
	}
	sqlDB, _ := DB.DB()
	sqlDB.SetMaxIdleConns(100)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(time.Second * 600)
}
