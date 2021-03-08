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
		time.Sleep(5 * time.Second)
		context.String(http.StatusOK, "hello world!")
		return
	})
	//engine.Run(fmt.Sprintf(":%d", setting.Conf.MainConfig.Port))
	srv := &http.Server{
		Addr:         fmt.Sprintf(":%d", setting.Conf.MainConfig.Port),
		Handler:      engine,
		WriteTimeout: 10 * time.Second,
	}
	go func() {
		err := srv.ListenAndServe()
		if err != nil && err != http.ErrServerClosed {
			panic(err.Error())
		}
	}()

	sig := make(chan os.Signal)

	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)

	<-sig

	ctx, cancelFunc := context.WithTimeout(context.Background(), time.Second*10)
	defer cancelFunc()

	fmt.Println("shutdown server ...")
	err := srv.Shutdown(ctx)
	if err != nil {
		zap.L().Error("server error...",
			zap.String("err", err.Error()),
		)
	}

}
