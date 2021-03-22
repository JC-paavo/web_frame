package redis

import (
	"fmt"
	"time"
	"web_frame/setting"

	"go.uber.org/zap"

	"github.com/go-redis/redis"
)

const SYSTEMPREFIX = "#system_"

func SaveSingIn(id int64) (err error) {
	if setting.Conf.Cluster {
		set := client.(*redis.ClusterClient).Set(fmt.Sprintf("%s%d", SYSTEMPREFIX, id), 1, time.Hour*2)
		if err = set.Err(); err != nil {
			zap.L().Error("cache set error", zap.Error(err))
		}

	} else {
		set := client.(*redis.Client).Set(fmt.Sprintf("%s%d", SYSTEMPREFIX, id), 1, time.Hour*2)
		if err = set.Err(); err != nil {
			zap.L().Error("cache set error", zap.Error(err))
		}
	}
	return
}

func GetKey(key string) (err error) {
	if setting.Conf.Cluster {
		get := client.(*redis.ClusterClient).Get(key)

		if err = get.Err(); err != nil && err != redis.Nil {
			zap.L().Error("cache get error", zap.Error(err))
		}

	} else {
		get := client.(*redis.Client).Get(key)
		if err = get.Err(); err != nil && err != redis.Nil {
			zap.L().Error("cache get error", zap.Error(err))
		}
	}
	return
}
