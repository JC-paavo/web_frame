package main

import (
	"web_frame/logger"
	"web_frame/router"
	"web_frame/setting"
)

func main() {
	setting.Init()
	err := logger.Init(setting.Conf.LogConfig)
	if err != nil {
		panic(err.Error())
	}
	router.Init()

}
