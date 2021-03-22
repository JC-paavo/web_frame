package redis

import (
	"strings"
	"web_frame/setting"

	"github.com/go-redis/redis"
	"go.uber.org/zap"
)

var client interface{}

func InitCluster(conf *setting.RedisConfig) {

	addr := strings.Split(conf.Host, ",")
	client = redis.NewClusterClient(&redis.ClusterOptions{
		Addrs:    addr,
		Password: conf.Password,
		PoolSize: conf.PoolSize,
	})
	_, err := client.(*redis.ClusterClient).Ping().Result()
	if err != nil {
		zap.L().Fatal("redis connection error...",
			zap.String("error", err.Error()),
		)
	}
}

func Init(conf *setting.RedisConfig) {
	client = redis.NewClient(&redis.Options{
		Addr:     conf.Host,
		Password: conf.Password,
		DB:       conf.DB,
		PoolSize: conf.PoolSize,
	})

	_, err := client.(*redis.Client).Ping().Result()
	if err != nil {
		zap.L().Fatal("redis connection error...",
			zap.String("error", err.Error()),
		)
	}
}
