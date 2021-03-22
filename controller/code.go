package controller

type ResCode int64

const (
	CodeSuccess ResCode = 1000 + iota
	CodeInvalidParam
	CodeUserExist
	CodeUserNotExist
	CodeInvalidPassword
	CodeServerBusy
	CodeAuthFailed
)

var codeMsgMap = map[ResCode]string{
	CodeSuccess:         "Successful!",
	CodeInvalidParam:    "无效的请求参数!",
	CodeUserExist:       "用户已存在!",
	CodeUserNotExist:    "用户或密码错误!",
	CodeInvalidPassword: "用户或密码错误!!",
	CodeServerBusy:      "服务繁忙!",
	CodeAuthFailed:      "token 验证失败!",
}

func (c ResCode) GetMsg() string {
	if msg, ok := codeMsgMap[c]; ok {
		return msg
	} else {
		return codeMsgMap[CodeServerBusy]
	}
}
