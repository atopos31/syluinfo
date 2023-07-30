package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

/*
{
	"code":10001,//错误码
	"msg": xx,   //提示信息
	"data": {},  //数据
}
*/

type ResponseData struct {
	Code ResCode     `json:"code"`
	Msg  interface{} `json:"msg"`
	Data interface{} `json:"data"`
}

// 参数绑定及校验错误处理
func ResponseBindError(c *gin.Context, err error) {
	// 获取validator.ValidationErrors类型的errors
	errs, ok := err.(validator.ValidationErrors)
	if !ok {
		ResponseErrorWithMsg(c, CodeInvalidParam, err.Error())
		return
	}
	ResponseErrorWithMsg(c, CodeServerBusy, removeTopStruct(errs.Translate(trans)))
}

func ResponseError(c *gin.Context, code ResCode) {
	c.JSON(http.StatusOK, ResponseData{
		Code: code,
		Msg:  code.Msg(),
		Data: nil,
	})
}

func ResponseErrorWithMsg(c *gin.Context, code ResCode, msg interface{}) {
	c.JSON(http.StatusOK, ResponseData{
		Code: code,
		Msg:  msg,
		Data: nil,
	})
}

func ResponseSuccess(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, ResponseData{
		Code: CodeSuccess,
		Msg:  CodeSuccess.Msg(),
		Data: data,
	})
}
