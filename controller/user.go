package controller

import (
	"web_frame/logic"
	"web_frame/modules"
	"web_frame/pkg/validor_translator"

	"go.uber.org/zap"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func SignUpHandler(c *gin.Context) {
	//获取业务参数
	var p modules.ParamSignUp
	err := c.ShouldBind(&p)
	if err != nil {
		zap.L().Error("SignUpHandler wirh invalid param", zap.Error(err))
		if errs, ok := err.(validator.ValidationErrors); ok {
			ResponseError(c, CodeInvalidParam, removeTopStruct(errs.Translate(validor_translator.Trans)))
			//c.JSON(http.StatusOK, gin.H{
			//	"msg": removeTopStruct(errs.Translate(validor_translator.Trans)),
			//})
			return
		} else {
			ResponseError(c, CodeServerBusy, err.Error())
			//c.JSON(http.StatusOK, gin.H{
			//	"msg": err.Error(),
			//})
			return
		}
	}
	//业务处理
	err = logic.SignUp(&p)

	if err != nil {
		zap.L().Error("注册失败", zap.Error(err))
		ResponseError(c, CodeUserExist, nil)
		//c.JSON(http.StatusOK, "Failed!")
		return
	} else {
		ResponseSuccess(c, nil)
		//c.JSON(http.StatusOK, "Successful!")
		return
	}
	//返回响应
}

func SignInHandler(c *gin.Context) {
	//获取业务参数
	var p modules.ParamSignIn
	err := c.ShouldBind(&p)
	if err != nil {
		zap.L().Error("SignInHandler wirh invalid param", zap.Error(err))
		if errs, ok := err.(validator.ValidationErrors); ok {
			ResponseError(c, CodeInvalidParam, removeTopStruct(errs.Translate(validor_translator.Trans)))
			//c.JSON(http.StatusInternalServerError, gin.H{
			//	"msg": removeTopStruct(errs.Translate(validor_translator.Trans)),
			//})
			return
		} else {
			ResponseError(c, CodeServerBusy, err.Error())
			//c.JSON(http.StatusInternalServerError, gin.H{
			//	"msg": err.Error(),
			//})
			return
		}
	}
	//业务处理
	err = logic.SignIn(&p)

	if err != nil {
		zap.L().Error("login failed", zap.Error(err))
		ResponseError(c, CodeInvalidPassword, nil)
		//c.JSON(http.StatusInternalServerError, err.Error())
		return
	} else {
		ResponseSuccess(c, nil)
		//c.JSON(http.StatusOK, "Successful!")
		return
	}
	//返回响应
}
