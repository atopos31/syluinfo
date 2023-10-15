package controller

import (
	"cld/logic"
	"cld/models"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func RecordHandler(c *gin.Context) {
	bindReq := new(models.ParamRecord)
	if err := c.ShouldBindJSON(bindReq); err != nil {
		zap.L().Error("FeedBackHandler ShouldBindJSON Error", zap.Error(err))
		ResponseBindError(c, err)
		return
	}
	// 获取当前请求的用户的id
	userID, err := getCurrentUser(c)
	if err != nil {
		ResponseError(c, CodeNeedLogin)
		return
	}
	recordID, err := logic.Record(bindReq, userID)
	if err != nil {
		zap.L().Error("RecordHandler FeedBack Error", zap.Error(err))
		ResponseErrorWithMsg(c, CodeServerBusy, err.Error())
		return
	}

	ResponseSuccess(c, recordID)
}

func RecordsHandler(c *gin.Context) {
	userID, err := getCurrentUser(c)
	if err != nil {
		ResponseError(c, CodeNeedLogin)
		return
	}

	records, err := logic.GetRedords(userID)
	if err != nil {
		zap.L().Error("RecordHandler FeedBack Error", zap.Error(err))
		ResponseErrorWithMsg(c, CodeServerBusy, err.Error())
		return
	}

	ResponseSuccess(c, records)

}

// 反馈请求处理函数
// @Summary 反馈接口
// @Tags 反馈相关接口
// @Accept application/json
// @Produce application/json
// @Param object body models.ParamFeedBack true "反馈参数,必填"
// @Success 1000 {string} ResponseData "code=1000,msg="success",data="null""
// @Failure 1001 {object} ResponseData "请求错误参数,code=1000+，msg里面是错误信息"
// @Router /feedback [post]
func FeedBackHandler(c *gin.Context) {
	bindReq := new(models.ParamFeedBack)
	if err := c.ShouldBindJSON(bindReq); err != nil {
		zap.L().Error("FeedBackHandler ShouldBindJSON Error", zap.Error(err))
		ResponseBindError(c, err)
		return
	}
	// 获取当前请求的用户的id
	userID, err := getCurrentUser(c)
	if err != nil {
		ResponseError(c, CodeNeedLogin)
		return
	}
	if err := logic.FeedBack(bindReq, userID); err != nil {
		zap.L().Error("FeedBackHandler FeedBack Error", zap.Error(err))
		ResponseErrorWithMsg(c, CodeServerBusy, err.Error())
		return
	}

	ResponseSuccess(c, nil)
}
