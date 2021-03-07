package router

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
	"web_frame/logger"
	"web_frame/setting"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func Init() {
	engine := gin.New()
	engine.Use(logger.GinLogger(), logger.GinRecovery(true))

	engine.GET("/hello", func(context *gin.Context) {
		context.String(http.StatusOK, "hello world!")
		return
	})
	//engine.Run(fmt.Sprintf(":%d", setting.Conf.MainConfig.Port))
	srv := &http.Server{
		Addr:    fmt.Sprintf(":%d", setting.Conf.MainConfig.Port),
		Handler: engine,
	}
	go func() {
		err := srv.ListenAndServe()
		if err != nil {
			panic(err.Error())
		}
	}()

	sig := make(chan os.Signal)

	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)

	timeout, cancelFunc := context.WithTimeout(context.Background(), time.Second*5)
	defer cancelFunc()
	for {
		<-sig
		err := srv.Shutdown(timeout)
		if err != nil {
			zap.L().Error("shutdown error...",
				zap.String("err", err.Error()),
			)
		}
	}
}
