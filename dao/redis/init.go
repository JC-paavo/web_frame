package redis

import (
	"github.com/go-redis/redis"
	"go.uber.org/zap"
	"strings"
	"web_frame/setting"
)

var Client interface{}

func InitCluster(conf *setting.RedisConfig) {

	addr := strings.Split(conf.Host, ",")
	Client = redis.NewClusterClient(&redis.ClusterOptions{
		Addrs:    addr,
		Password: conf.Password,
		PoolSize: conf.PoolSize,
	})
	_, err := Client.(*redis.ClusterClient).Ping().Result()
	if err != nil {
		zap.L().Fatal("redis connection error...",
			zap.String("error", err.Error()),
		)
	}
}

func Init(conf *setting.RedisConfig) {
	Client = redis.NewClient(&redis.Options{
		Addr:     conf.Host,
		Password: conf.Password,
		DB:       conf.DB,
		PoolSize: conf.PoolSize,
	})

	_, err := Client.(*redis.Client).Ping().Result()
	if err != nil {
		zap.L().Fatal("redis connection error...",
			zap.String("error", err.Error()),
		)
	}
}
