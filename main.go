package main

import (
	"cld/cos"
	"cld/dao/mysql"
	"cld/dao/redis"
	"cld/logger"
	"cld/pkg/snowflake"
	"cld/routes"
	"cld/settings"
	"fmt"

	"go.uber.org/zap"
)

// @title sylu项目接口文档
// @version 1.1
// @description 致力于为同学们提供校园服务(忽略状态码，所有响应都是200)
// @contact.name hakcerxiao
// @contact.url http://www.hackerxiao.online
// @BasePath /api/v1
func main() {
	//1.加载配置文件
	if err := settings.Init(); err != nil {
		fmt.Println("init setting failed, err :" + err.Error())
		return
	}

	//2.初始化日志
	if err := logger.Init(&settings.Conf.Log); err != nil {
		fmt.Println("init logger falied, err :" + err.Error())
		return
	}
	defer zap.L().Sync()
	zap.L().Debug("logger init success...")

	//3.初始化Mysql
	if err := mysql.Init(&settings.Conf.MySQL); err != nil {
		fmt.Println("init mysql failed, err :" + err.Error())
		return
	}
	defer mysql.Close()

	//4.初始化Redis
	if err := redis.Init(&settings.Conf.Redis); err != nil {
		fmt.Println("init redis failed, err :" + err.Error())
		return
	}
	defer redis.Close()

	//雪花算法初始化
	if err := snowflake.Init(settings.Conf.Snowflake.StartTime, settings.Conf.Snowflake.MachineID); err != nil {
		fmt.Println("snowflake redis failed, err :" + err.Error())
		return
	}

	//COS初始化
	if err := cos.Init(&settings.Conf.Cos); err != nil {
		fmt.Println("cos init err:" + err.Error())
		return
	}

	//5.注册路由
	routes.Setup(&settings.Conf.App)

	//6.启动服务
	routes.StartServer(settings.Conf.App.Port)
}
