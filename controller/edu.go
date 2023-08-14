package controller

import (
	"cld/dao/mysql"
	"cld/dao/resty_tool"
	"cld/logic"
	"cld/models"
	"errors"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// 绑定教务请求处理函数
// @Summary 绑定接口
// @Tags sylu相关接口
// @Accept application/json
// @Produce application/json
// @Param Authorization header string true "Bearer JWT"
// @Param object body models.ParamBind true "绑定参数,必填"
// @Success 1000 {object} models.ReqBind "code=1000,msg="success","
// @Failure 1001 {object} ResponseData "请求错误参数,code=1000+，msg里面是错误信息"
// @Router /edu/bind [post]
func BingEducationalHandler(c *gin.Context) {
	bindReq := new(models.ParamBind)
	if err := c.ShouldBindJSON(bindReq); err != nil {
		zap.L().Error("BingEducational ShouldBindJSON Error", zap.Error(err))
		ResponseBindError(c, err)
		return
	}
	// 获取当前请求的用户的id
	userID, err := getCurrentUser(c)
	if err != nil {
		ResponseError(c, CodeNeedLogin)
		return
	}
	//登录测试
	reqLogin, err := logic.BintLogin(bindReq, userID)
	if err != nil {
		zap.L().Error("BingEducational logic.BintLogin Error", zap.Error(err))
		ResponseErrorWithMsg(c, CodeServerBusy, err.Error())
		return
	}

	ResponseSuccess(c, reqLogin)
}

// 获取cookie请求处理函数
// @Summary 获取cookie接口
// @Tags sylu相关接口
// @Accept application/json
// @Produce application/json
// @Param Authorization header string true "Bearer JWT"
// @Success 1000 {string} data "code=1000,msg="success","
// @Failure 1001 {object} ResponseData "请求错误参数,code=1000+，msg里面是错误信息"
// @Router /edu/cookie [get]
func CookieHandler(c *gin.Context) {
	// 获取当前请求的用户的id
	userID, err := getCurrentUser(c)
	if err != nil {
		ResponseError(c, CodeNeedLogin)
		return
	}

	cookisString, err := logic.GetCookie(userID)
	if err != nil {
		zap.L().Error("CookieHandler logic.GetCookie Error", zap.Error(err))
		if errors.Is(err, mysql.ErrorUnbound) {
			ResponseError(c, CodeUnbound)
			return
		}
		ResponseError(c, CodeServerBusy)
		return
	}

	ResponseSuccess(c, cookisString)
}

// 获取学期请求处理函数
// @Summary 获取学期接口
// @Tags sylu相关接口
// @Accept application/json
// @Produce application/json
// @Param Authorization header string true "Bearer JWT"
// @Success 1000 {object} models.ResSemeSter "code=1000,msg="success","
// @Failure 1001 {object} ResponseData "请求错误参数,code=1000+，msg里面是错误信息"
// @Router /edu/semester [get]
func SemesterHandler(c *gin.Context) {
	userID, err := getCurrentUser(c)
	if err != nil {
		ResponseError(c, CodeNeedLogin)
		return
	}

	semeSterRes, err := logic.GetSemeSter(userID)
	if err != nil {
		zap.L().Error("SemesterHandler logic.GetSemeSter Error", zap.Error(err))
		if errors.Is(err, mysql.ErrorUnbound) {
			ResponseError(c, CodeUnbound)
			return
		}
		ResponseError(c, CodeServerBusy)
		return
	}

	ResponseSuccess(c, semeSterRes)
}

// 获取课表请求处理函数
// @Summary 获取课表接口
// @Tags sylu相关接口
// @Accept application/json
// @Produce application/json
// @Param Authorization header string true "Bearer JWT"
// @Param object body models.ParamCourse true "课表参数,必填,其中semester为3或12表示某学期，例如year=2022 semester=3 表示2022-2023学年第一学期"
// @Success 1000 {object} models.ReqCourse "code=1000,msg="success","
// @Failure 1001 {object} ResponseData "请求错误参数,code=1000+，msg里面是错误信息"
// @Router /edu/courses [post]
func CourseHandler(c *gin.Context) {
	bindCourse := new(models.ParamCourse)
	if err := c.ShouldBindJSON(bindCourse); err != nil {
		zap.L().Error("CourseHandler ShouldBindJSON Error", zap.Error(err))
		ResponseBindError(c, err)
		return
	}
	courses, err := logic.GetCourse(bindCourse)
	if err != nil {
		zap.L().Error("CourseHandler logic.GetCourse Error", zap.Error(err))
		if errors.Is(err, resty_tool.ErrorLapse) {
			ResponseError(c, CodeInvalidCookie)
			return
		} else if errors.Is(err, resty_tool.ErrorCourseNoOpen) {
			ResponseError(c, CodeNotData)
			return
		}
		ResponseErrorWithMsg(c, CodeServerBusy, err.Error())
		return
	}

	ResponseSuccess(c, courses)
}

// 获取某学期全部成绩请求处理函数
// @Summary 获取成绩接口
// @Tags sylu相关接口
// @Accept application/json
// @Produce application/json
// @Param Authorization header string true "Bearer JWT"
// @Param object body models.ParamGrades true "成绩参数,必填，其中semester为3或12表示某学期，例如year=2022 semester=3 表示2022-2023学年第一学期"
// @Success 1000 {object} []models.JsonGrades "code=1000,msg="success","
// @Failure 1001 {object} ResponseData "请求错误参数,code=1000+，msg里面是错误信息"
// @Router /edu/grades [post]
func GradesHandler(c *gin.Context) {
	bindGrades := new(models.ParamGrades)
	if err := c.ShouldBindJSON(bindGrades); err != nil {
		zap.L().Error("GradesHandler ShouldBindJSON Error", zap.Error(err))
		ResponseBindError(c, err)
		return
	}

	grades, err := logic.GetGrades(bindGrades)
	if err != nil {
		zap.L().Error("GradesHandler logic.GetCourse Error", zap.Error(err))
		if errors.Is(err, resty_tool.ErrorLapse) {
			ResponseError(c, CodeInvalidCookie)
			return
		} else if errors.Is(err, resty_tool.ErrorGradesNoOpen) {
			ResponseError(c, CodeNotData)
			return
		}
		ResponseErrorWithMsg(c, CodeServerBusy, err.Error())
		return
	}

	ResponseSuccess(c, grades)
}

// 获取成绩详情请求处理函数
// @Summary 获取成绩详情接口
// @Tags sylu相关接口
// @Accept application/json
// @Produce application/json
// @Param Authorization header string true "Bearer JWT"
// @Param object body models.ParamGradeDetaile true "成绩参数,必填，其中semester为3或12表示某学期，例如year=2022 semester=3 表示2022-2023学年第一学期"
// @Success 1000 {object} []models.ResGradeDetail "code=1000,msg="success","
// @Failure 1001 {object} ResponseData "请求错误参数,code=1000+，msg里面是错误信息"
// @Router /edu/grade/detaile [post]
func GradeDetaileHandler(c *gin.Context) {
	bindGradeDetail := new(models.ParamGradeDetaile)
	if err := c.ShouldBindJSON(bindGradeDetail); err != nil {
		zap.L().Error("GradesHandler ShouldBindJSON Error", zap.Error(err))
		ResponseBindError(c, err)
		return
	}

	gradeDetail, err := logic.GetGradeDetail(bindGradeDetail)
	if err != nil {
		zap.L().Error("GradeDetaileHandler logic.GetGradeDetail Error", zap.Error(err))
		if errors.Is(err, resty_tool.ErrorLapse) {
			ResponseError(c, CodeInvalidCookie)
			return
		}
		ResponseErrorWithMsg(c, CodeServerBusy, err.Error())
		return
	}

	ResponseSuccess(c, gradeDetail)
}

// 获取绩点请求处理函数
// @Summary 获取绩点接口
// @Tags sylu相关接口
// @Accept application/json
// @Produce application/json
// @Param Authorization header string true "Bearer JWT"
// @Param object body models.ParamGpa true "此接口响应时间>=4s，暂无优化思路"
// @Success 1000 {object} models.ResGpa "code=1000,msg="success","
// @Failure 1001 {object} ResponseData "请求错误参数,code=1000+，msg里面是错误信息"
// @Router /edu/gpas [post]
func GpaHandler(c *gin.Context) {
	bindGpa := new(models.ParamGpa)
	if err := c.ShouldBindJSON(bindGpa); err != nil {
		zap.L().Error("GpaHandler ShouldBindJSON Error", zap.Error(err))
		ResponseBindError(c, err)
		return
	}

	resGpa, err := logic.GetGpas(bindGpa)
	if err != nil {
		zap.L().Error("GpaHandler logic.GetGpas Error", zap.Error(err))
		if errors.Is(err, resty_tool.ErrorLapse) {
			ResponseError(c, CodeInvalidCookie)
			return
		}
		ResponseErrorWithMsg(c, CodeServerBusy, err.Error())
		return
	}

	ResponseSuccess(c, resGpa)
}
