package models

type QuerySendEmail struct {
	Mode  string `json:"mode" form:"mode" binding:"required,oneof=sign recoverpass"`
	Email string `json:"email" form:"email" binding:"required,email"`
}
