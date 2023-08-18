package models

import "gorm.io/gorm"

// 数据库反馈信息表
type FeedBack struct {
	gorm.Model
	Uuid     int64  `json:"uuid"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Title    string `json:"title"`
	Content  string `json:"content"`
}
