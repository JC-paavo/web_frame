package controller

import (
	"strings"
	"web_frame/logic"

	"go.uber.org/zap"

	"github.com/gin-gonic/gin"
)

func AuthTokenMiddler() func(c *gin.Context) {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			ResponseError(c, CodeAuthFailed, nil)
			c.Abort()
			return
		}
		// 按空格分割
		parts := strings.SplitN(authHeader, " ", 2)
		if !(len(parts) == 2 && parts[0] == "Bearer") {
			ResponseError(c, CodeAuthFailed, nil)
			c.Abort()
			return
		}
		// parts[1]是获取到的tokenString，我们使用之前定义好的解析JWT的函数来解析它
		mc, err := logic.ParseToken(parts[1])
		if err != nil {
			ResponseError(c, CodeAuthFailed, nil)
			c.Abort()
			return
		}
		//校验用户是否存在
		err = logic.CheckUserId(mc.User_id)
		if err != nil {
			zap.L().Error("token 校验失败", zap.Error(err))
			ResponseError(c, CodeAuthFailed, nil)
			c.Abort()
			return
		}
		// 将当前请求的username信息保存到请求的上下文c上
		c.Set("user_id", mc.User_id)
		c.Next()
	}
}
