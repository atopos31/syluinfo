package controller

type ResCode int64

const (
	//"成功！"
	CodeSuccess ResCode = 1000 + iota
	//"请求参数错误"
	CodeInvalidParam
	//"邮箱已被注册"
	CodeEmailExist
	//"邮箱不存在"
	CodeEmailNotExist
	//"邮箱或密码错误"
	CodeInvalidPassword
	//验证码错误
	CodeInvalidCaptcha
	//验证码不存在或已过期
	CodeCaptchaNotExistOrTimeOut
	//未绑定教务！
	CodeUnbound
	//"服务繁忙"
	CodeServerBusy
	//"无效的token"
	CodeInvalidToken
	//"需要登陆"
	CodeNeedLogin
	//"cookie无效"
	CodeInvalidCookie
	//"数据不存在"
	CodeNotData
)

var codeMsgMap = map[ResCode]string{
	CodeSuccess:                  "success",
	CodeInvalidParam:             "请求参数错误",
	CodeEmailExist:               "邮箱已被注册",
	CodeEmailNotExist:            "邮箱不存在",
	CodeInvalidPassword:          "邮箱或密码错误",
	CodeInvalidCaptcha:           "验证码错误",
	CodeCaptchaNotExistOrTimeOut: "验证码不存在或已过期",
	CodeUnbound:                  "未绑定教务账号",
	CodeServerBusy:               "服务繁忙",
	CodeInvalidToken:             "无效的token",
	CodeNeedLogin:                "需要登陆",
	CodeInvalidCookie:            "cookie无效，请重新进入小程序",
	CodeNotData:                  "数据不存在",
}

func (c ResCode) Msg() string {
	msg, ok := codeMsgMap[c]

	if !ok {
		msg = codeMsgMap[CodeServerBusy]
	}

	return msg
}
