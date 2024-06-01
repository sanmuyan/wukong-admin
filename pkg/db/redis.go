package db

import (
	"context"
	"github.com/go-redis/redis/v8"
	"github.com/sirupsen/logrus"
	"wukong/pkg/config"
)

var RDB *redis.Client

func InitRedis() {
	var ctx = context.Background()
	RDB = redis.NewClient(&redis.Options{
		Addr:     config.Conf.Database.Redis,
		PoolSize: 100,
	})
	if ok, err := RDB.Ping(ctx).Result(); ok != "PONG" && err != nil {
		logrus.Fatalf("redis connect PING: %s error: %s", ok, err.Error())
	}

}
