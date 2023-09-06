package routes

import (
	"cld/controller"
	_ "cld/docs"
	"cld/logger"
	"cld/middlewares"
	"cld/settings"
	"fmt"

	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

var r *gin.Engine

func Setup(cfg *settings.AppConfig) {
	r = gin.Default()
	//翻译器初始化
	controller.InitTrans("zh")
	//跨域
	r.Use(middlewares.AllAlowCors())
	//日志写入中间件
	r.Use(logger.GinLogger(), logger.GinRecovery(false))
	//接口文档UI
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	//路由组
	baseapi := r.Group("/api/v1")
	//关于页面，可选，自行创建对应目录以及md文件
	baseapi.StaticFile("/about", "./about/about.md")
	//反馈
	baseapi.POST("/feedback", middlewares.JWTAuthMiddleware(), controller.FeedBackHandler)
	//auth相关
	authen := baseapi.Group("/auth")
	{
		authen.POST("/login", controller.LoginHandler)
		authen.POST("/signup", controller.SignUpHandler)
		authen.GET("/sendemail", controller.SendEmailHandler)
		authen.POST("/resetpass", middlewares.JWTAuthMiddleware(), controller.ReSetPassHandler)
		authen.POST("/recoverpass", controller.ReCoverPassHandler)

		authen.GET("/coskey", middlewares.JWTAuthMiddleware(), controller.GetCosKeyHandler)
	}
	//edu相关
	edu := baseapi.Group("/edu")
	//加入token中间件验证
	edu.Use(middlewares.JWTAuthMiddleware())
	{
		edu.GET("/cookie", controller.CookieHandler)
		edu.GET("/semester", controller.SemesterHandler)

		edu.POST("/bind", controller.BingEducationalHandler)
		edu.GET("/courses/auto", controller.AutoCourseHandler)
		edu.POST("/courses", controller.CourseHandler)
		edu.POST("/grades", controller.GradesHandler)
		edu.POST("/grade/detaile", controller.GradeDetaileHandler)
		edu.POST("/gpas", controller.GpaHandler)
		edu.GET("/cale", controller.CaleHandler)
	}

	r.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, cfg.Name)
	})

}

func StartServer(p int) {
	r.Run(fmt.Sprintf(":%s", strconv.Itoa(p)))
}
