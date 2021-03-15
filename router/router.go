package router

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
	"web_frame/controller"
	"web_frame/logger"
	"web_frame/pkg/validor_translator"
	"web_frame/setting"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func Init(cfg *setting.MainConfig) {
	gin.SetMode(cfg.Mode)
	if err := validor_translator.InitTrans("zh"); err != nil {
		zap.L().Error("初始化参数翻译故障!", zap.String("err", err.Error()))
	}

	engine := gin.New()
	engine.Use(logger.GinLogger(), logger.GinRecovery(true))

	engine.GET("/hello", func(context *gin.Context) {
		time.Sleep(5 * time.Second)
		context.String(http.StatusOK, "hello world!")
		return
	})

	engine.POST("/signup", controller.SignUpHandler)

	//engine.Run(fmt.Sprintf(":%d", setting.Conf.MainConfig.Port))
	srv := &http.Server{
		Addr:         fmt.Sprintf(":%d", cfg.Port),
		Handler:      engine,
		WriteTimeout: 10 * time.Second,
	}

	go func() {
		err := srv.ListenAndServe()
		if err != nil && err != http.ErrServerClosed {
			panic(err.Error())
		}
	}()

	//平滑关闭
	syscal(srv)
}
func syscal(srv *http.Server) {
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
