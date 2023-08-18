package logic

import (
	"cld/dao/mysql"
	"cld/models"
)

func FeedBack(req *models.ParamFeedBack, userID int64) error {
	//通过userID获取用户名和邮箱
	userInfo, err := mysql.GetUserInfoByUuid(userID)
	if err != nil {
		return err
	}

	feedInfo := &models.FeedBack{
		Uuid:     userInfo.Uuid,
		Username: userInfo.Username,
		Email:    userInfo.Email,
		Title:    req.Title,
		Content:  req.Content,
	}

	err = mysql.CreatFeed(feedInfo)
	if err != nil {
		return err
	}

	return nil
}
