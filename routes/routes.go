package routes

import (
	"cld/controller"
	_ "cld/docs"
	"cld/logger"
	"cld/middlewares"
	"cld/settings"
	"fmt"
	"os"

	"net/http"
	"strconv"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"go.uber.org/zap"
)

var r *gin.Engine

func Setup(cfg *settings.AppConfig) {
	r = gin.Default()
	//翻译器初始化
	controller.InitTrans("zh")
	//跨域
	r.Use(cors.Default())
	//日志写入中间件
	r.Use(logger.GinLogger(), logger.GinRecovery(false))
	//接口文档UI
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	//路由组
	baseapi := r.Group("/api/v1")
	//日志对外显示接口 dev阶段
	baseapi.GET("/log", showLog)
	//auth相关
	authen := baseapi.Group("/auth")
	{
		authen.POST("/login", controller.LoginHandler)
		authen.POST("/signup", controller.SignUpHandler)
		authen.GET("/sendemail", controller.SendEmailHandler)
		authen.POST("/resetpass", controller.ReSetPassHandler)
		authen.POST("/recoverpass", controller.ReCoverPassHandler)

		authen.GET("/coskey", middlewares.JWTAuthMiddleware(), controller.GetCosKeyHandler)
	}
	//edu相关
	edu := baseapi.Group("/edu")
	//加入token中间件验证
	edu.Use(middlewares.JWTAuthMiddleware())
	{
		edu.GET("/cookie", controller.CookieHandler)

		edu.POST("/bind", controller.BingEducationalHandler)
		edu.POST("/courses", controller.CourseHandler)
		edu.POST("/grades", controller.GradesHandler)
		edu.POST("/grade/detaile", controller.GradeDetaileHandler)
		edu.POST("/gpas", controller.GpaHandler)
	}

	r.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, cfg.Name)
	})
}

func StartServer(p int) {
	r.Run(fmt.Sprintf(":%s", strconv.Itoa(p)))
}

func showLog(c *gin.Context) {
	data, err := os.ReadFile("./web_app.log")
	if err != nil {
		zap.L().Error("ioutil.ReadFile Error :", zap.Error(err))
		c.String(http.StatusOK, "日志加载失败")
	}
	c.String(http.StatusOK, string(data))
}
