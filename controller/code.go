package controller

type ResCode int64

const (
	CodeSuccess ResCode = 1000 + iota
	CodeInvalidParam
	CodeUserExist
	CodeUserNotExist
	CodeInvalidPassword
	CodeServerBusy

	CodeNeedLogin
	CodeInvalidToken
)

var codeMsgMap = map[ResCode]string{
	CodeSuccess:         "成功200",
	CodeInvalidParam:    "参数错误",
	CodeUserExist:       "用户已存在",
	CodeUserNotExist:    "用户不存在",
	CodeInvalidPassword: "用户密码错误",
	CodeServerBusy:      "服务器繁忙",

	CodeNeedLogin:    "需要登陆",
	CodeInvalidToken: "无效的token",
}

func (c ResCode) Msg() string {
	msg, ok := codeMsgMap[c]
	if !ok {
		msg = codeMsgMap[CodeServerBusy]
	}

	return msg
}
