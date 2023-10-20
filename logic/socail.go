package logic

import (
	"cld/dao/mysql"
	"cld/models"
	"cld/pkg/snowflake"
	"strconv"
	"time"
)

func Record(req *models.ParamRecord, userID int64) (id string, err error) {

	var int64id int64
	if req.RecordID == "" {
		int64id = snowflake.GenID()
	} else {
		int64id, _ = strconv.ParseInt(req.RecordID, 10, 64)
	}

	if req.Type == 1 {
		err = pushOrChangeRecord(req, int64id, userID)
	} else if req.Type == 2 {
		err = mysql.DelRecordByID(int64id)
	}

	return strconv.FormatInt(int64id, 10), err
}

func GetRedords(userID int64) ([]models.ResRcords, error) {
	records, err := mysql.GetRedordsByUserID(userID)
	if err != nil {
		return nil, err
	}

	var resRecords []models.ResRcords
	var resRecord models.ResRcords
	for _, v := range records {
		resRecord.Title = v.Title
		resRecord.Content = v.Content
		resRecord.Time = v.CreatedAt.UnixMilli()
		resRecord.ID = strconv.FormatInt(v.RecordID, 10)

		resRecords = append(resRecords, resRecord)
	}

	return resRecords, nil
}

func pushOrChangeRecord(req *models.ParamRecord, id int64, userID int64) error {

	rerord := &models.Record{
		RecordID: id,
		UserID:   userID,
		Title:    req.Title,
		Content:  req.Content,
	}

	if err := mysql.CreatOrUpdateRecord(rerord); err != nil {
		return err
	}
	return nil
}

func GetNews() ([]models.ResNews, error) {
	news, err := mysql.GetNews()
	if err != nil {
		return nil, err
	}

	var resNews []models.ResNews
	var resNew models.ResNews

	for _, v := range news {
		resNew.ID = strconv.Itoa(int(v.ID))
		resNew.Title = v.Title
		resNew.Content = v.Content
		resNew.Path = v.LogoPath
		resNew.Time = v.CreatedAt.Format(time.DateTime)
		resNews = append(resNews, resNew)
	}

	return resNews, nil
}

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

	if err = mysql.CreatFeed(feedInfo); err != nil {
		return err
	}

	return nil
}
