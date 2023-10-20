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

// 便签表
type Record struct {
	gorm.Model
	RecordID int64  `json:"record_id" gorm:"unique"`
	UserID   int64  `json:"userID"`
	Title    string `json:"title"`
	Content  string `json:"content"`
}

// 数据库反馈信息表
type FeedBack struct {
	gorm.Model
	Uuid     int64  `json:"uuid" `
	Username string `json:"username"`
	Email    string `json:"email"`
	Title    string `json:"title"`
	Content  string `json:"content"`
}

// 校园资讯表
type SchoolNews struct {
	gorm.Model
	LogoPath string `json:"logo"`
	Title    string `json:"title"`
	Content  string `json:"content"`
}
