package controller

import "cld/models"

type _ResponseLoginData struct {
	Code ResCode          `json:"code" default:"1005"`
	Msg  string           `json:"msg" default:"success"`
	Data *models.ReqLogin `json:"data" `
}
