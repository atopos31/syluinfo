package mysql

import (
	"cld/models"
	"cld/settings"
	"fmt"

	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

func Init(cfg *settings.MySQLConfig) (err error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		cfg.User,
		cfg.Password,
		cfg.Host,
		cfg.Port,
		cfg.DBName,
	)
	mysqlConfig := gorm.Config{
		SkipDefaultTransaction: false,
		//这里可以添加这个gorm的logger
		//Logger: ,
	}

	db, err = gorm.Open(mysql.Open(dsn), &mysqlConfig)
	if err != nil {
		zap.L().Error("mysql init errr", zap.Error(err))
		panic("warn err!")
	}

	//建表
	initMigrate()
	return
}

func initMigrate() {
	//建立用户表
	db.AutoMigrate(&models.User{})
	//建立教务学生信息表
	db.AutoMigrate(&models.SyluUser{})
	//建立教务学生登录信息表
	db.AutoMigrate(&models.SyluPass{})
	//建立反馈信息表
	db.AutoMigrate(&models.FeedBack{})
}

func Close() {
	db.Statement.ReflectValue.Close()
}
