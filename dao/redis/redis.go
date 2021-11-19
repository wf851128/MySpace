package redis

import (
	"MySpace/settings"
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
)

var rdb *redis.Client

func Init() (err error) {
	rdb = redis.NewClient(&redis.Options{
		Addr: fmt.Sprintf("%s:%d",
			settings.Conf.RedisConfig.Host,
			settings.Conf.RedisConfig.Port,
		),
		Password: settings.Conf.RedisConfig.Password,
		DB:       settings.Conf.RedisConfig.DB,
		PoolSize: settings.Conf.RedisConfig.PoolSize,
	})
	_, err = rdb.Ping(context.Background()).Result()
	return err
}

func Close() {
	_ = rdb.Close()
}
