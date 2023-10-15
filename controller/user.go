package controller

import (
	"cld/dao/mysql"
	"cld/dao/redis"
	"cld/logic"
	"cld/models"
	"cld/settings"
	"errors"
	"fmt"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// 登录请求处理函数
// @Summary 登录接口
// @Tags auth相关接口
// @Accept application/json
// @Produce application/json
// @Param object body models.ParamLogin true "登录参数,必填"
// @Success 1000 {object} _ResponseLoginData "code=1000,msg="success",成功返回token,若绑定教务信息，也会包含教务学生信息"
// @Failure 1001 {object} ResponseData "请求错误参数,code=1000+，msg里面是错误信息"
// @Router /auth/login [post]
func LoginHandler(c *gin.Context) {
	loginReq := new(models.ParamLogin)
	if err := c.ShouldBindJSON(loginReq); err != nil {
		zap.L().Error("LoginHandler ShouldBindJSON Error", zap.Error(err))
		ResponseBindError(c, err)
		return
	}

	resLogin, err := logic.Login(loginReq)
	if err != nil {
		zap.L().Error("LoginHandler logic.Login Error", zap.Error(err))
		ResponseError(c, CodeInvalidPassword)
		return
	}

	ResponseSuccess(c, resLogin)
}

// 注册请求处理函数
// @Summary 注册接口
// @Tags auth相关接口
// @Accept application/json
// @Produce application/json
// @Param object body models.ParamSignUp true "注册参数,必填"
// @Success 1000 {object} ResponseData "code=1000,msg="success",data=null"
// @Failure 1001 {object} ResponseData "code=1000+，msg里面是错误信息,data=null"
// @Router /auth/signup [post]
func SignUpHandler(c *gin.Context) {
	siginReq := new(models.ParamSignUp)
	if err := c.ShouldBindJSON(siginReq); err != nil {
		zap.L().Error("SignupHandler ShouldBindJSON Error", zap.Error(err))
		ResponseBindError(c, err)
		return
	}

	err := logic.SignUp(siginReq)
	if err != nil {
		zap.L().Error("SignUpHandler logic.SignUp Error", zap.Error(err))
		if errors.Is(err, redis.ErrorNotExists) {
			ResponseError(c, CodeCaptchaNotExistOrTimeOut)
			return
		}
		if errors.Is(err, logic.ErrorFailedCaptcha) {
			ResponseError(c, CodeInvalidCaptcha)
			return
		}
		ResponseError(c, CodeServerBusy)
		return
	}

	ResponseSuccess(c, nil)
}

// 发送邮件请求处理函数
// @Summary 邮箱验证接口
// @Tags auth相关接口
// @Param object query models.QuerySendEmail true "邮件参数,模式->注册：sign ->找回密码->recoverpass"
// @Success 1000 {object} ResponseData "code=1000,msg="success",data=null"
// @Failure 1005 {object} ResponseData "code=1000+，msg里面是错误信息,data=null"
// @Router /auth/sendemail [get]
func SendEmailHandler(c *gin.Context) {
	fmt.Println(c.Request.Body)
	queryEmail := new(models.QuerySendEmail)
	if err := c.ShouldBindQuery(queryEmail); err != nil {
		ResponseBindError(c, err)
		return
	}

	if err := logic.SendEmail(*queryEmail); err != nil {
		zap.L().Error("logic.SendEmail Error", zap.Error(err))
		if errors.Is(err, logic.ErrorEmailExist) {
			ResponseError(c, CodeEmailExist)
			return
		}
		if errors.Is(err, logic.ErrorEmailNotExist) {
			ResponseError(c, CodeEmailNotExist)
			return
		}
		ResponseError(c, CodeServerBusy)
		return
	}

	ResponseSuccess(c, nil)
}

// 重置密码处理函数
// @Summary 重置密码验证接口
// @Tags auth相关接口
// @Param Authorization header string true "Bearer JWT"
// @Param object body models.ParamReSet true "使用旧密码新密码重置"
// @Success 1000 {object} ResponseData "code=1000,msg="success",data=null"
// @Failure 1005 {object} ResponseData "code=1000+，msg里面是错误信息,data=null"
// @Router /auth/resetpass [post]
func ReSetPassHandler(c *gin.Context) {
	paramReSet := new(models.ParamReSet)
	if err := c.ShouldBindJSON(paramReSet); err != nil {
		ResponseBindError(c, err)
		return
	}
	userID, err := getCurrentUser(c)
	if err != nil {
		ResponseError(c, CodeNeedLogin)
		return
	}

	if err := logic.ReSetPass(userID, paramReSet); err != nil {
		zap.L().Error("logic.ReSetPass Error", zap.Error(err))
		if errors.Is(err, mysql.ErrorInvalidPassword) {
			ResponseError(c, CodeInvalidPassword)
			return
		}
		ResponseErrorWithMsg(c, CodeServerBusy, err.Error())
		return
	}

	ResponseSuccess(c, nil)
}

// 找回密码处理函数
// @Summary 找回密码验证接口
// @Tags auth相关接口
// @Param object body models.ParamReCover true "使用邮箱验证码新密码重置"
// @Success 1000 {object} ResponseData "code=1000,msg="success",data=null"
// @Failure 1005 {object} ResponseData "code=1000+，msg里面是错误信息,data=null"
// @Router /auth/recoverpass [post]
func ReCoverPassHandler(c *gin.Context) {
	parmReCover := new(models.ParamReCover)
	if err := c.ShouldBindJSON(parmReCover); err != nil {
		ResponseBindError(c, err)
		return
	}

	if err := logic.ReCoverPass(parmReCover); err != nil {
		zap.L().Error("logic.ReCoverPass", zap.Error(err))
		if errors.Is(err, redis.ErrorNotExists) {
			ResponseError(c, CodeCaptchaNotExistOrTimeOut)
			return
		}
		if errors.Is(err, logic.ErrorFailedCaptcha) {
			ResponseError(c, CodeInvalidCaptcha)
			return
		}
		ResponseErrorWithMsg(c, CodeServerBusy, err.Error())
		return
	}

	ResponseSuccess(c, nil)
}

// 获取COS临时密钥处理函数
// @Summary 获取COS临时密钥接口
// @Tags auth相关接口
// @Param Authorization header string true "Bearer JWT"
// @Success 1000 {object} ResponseData "code=1000,msg="success",data里面是cos临时密钥数据"
// @Failure 1005 {object} ResponseData "code=1000+，msg里面是错误信息,data=null"
// @Router /auth/coskey [get]
func GetCosKeyHandler(c *gin.Context) {
	resKey, err := logic.GetCosKey(&settings.Conf.Cos)
	if err != nil {
		ResponseError(c, CodeServerBusy)
		return
	}

	ResponseSuccess(c, resKey)
}
