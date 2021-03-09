package main

import (
	"web_frame/dao/mysql"
	"web_frame/dao/redis"
	"web_frame/logger"
	"web_frame/router"
	"web_frame/setting"
)

func main() {
	//初始化配置
	setting.Init()

	//初始化日志
	err := logger.Init(setting.Conf.LogConfig)
	if err != nil {
		panic(err.Error())
	}

	//初始化redis
	if setting.Conf.Cluster {
		redis.InitCluster(setting.Conf.RedisConfig)
	} else {
		redis.Init(setting.Conf.RedisConfig)

	}

	//初始化数据
	mysql.Init(setting.Conf.MysqlConfig)

	router.Init()

}
