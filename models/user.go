package models

import (
	"gorm.io/gorm"
)

// 数据库用户表
type User struct {
	gorm.Model
	Uuid     int64  `json:"uuid"`
	Username string `json:"username"`
	Avatar   string `json:"avatar"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

// 数据库学生教务信息表
type SyluUser struct {
	gorm.Model
	Uuid       int64  `json:"uuid"`
	ReUsername string `json:"reusername"`
	StudentID  string `json:"studentID"`
	Grade      string `json:"grade"`
	College    string `json:"college"`
	Major      string `json:"major"`
}

// 数据库学生账号密码表
type SyluPass struct {
	gorm.Model
	Uuid      int64  `json:"uuid"`
	StudentID string `json:"studentID"`
	Password  string `json:"password"`
}
