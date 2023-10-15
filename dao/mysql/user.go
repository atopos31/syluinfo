package mysql

import (
	"cld/models"
	"errors"

	"go.uber.org/zap"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

var (
	ErrorUserExist       = errors.New("用户已存在")
	ErrorUserNoExist     = errors.New("用户不存在")
	ErrorInvalidPassword = errors.New("密码错误")
	ErrorUnbound         = errors.New("未绑定教务账号")
)

func Login(loginReq *models.ParamLogin) (user *models.User, err error) {
	err = db.Model(&models.User{}).
		Where("email = ?", loginReq.Email).
		First(&user).Error
	//邮箱未注册
	if err == gorm.ErrRecordNotFound {
		return nil, ErrorUserNoExist
	}
	//查询出错
	if err != nil {
		zap.L().Error("Mysql Login search Error", zap.Error(err))
		return nil, err
	}
	//密码错误
	enPassword := encryptPassword(loginReq.Password)
	if user.Password != enPassword {
		return nil, ErrorInvalidPassword
	}

	return
}

func CreateUser(siginUserInfo *models.User) (err error) {
	//密码加密
	siginUserInfo.Password = encryptPassword(siginUserInfo.Password)
	return db.Model(&models.User{}).Create(siginUserInfo).Error
}

func CheckEmailExist(email string) (bool, error) {
	var count int64

	if err := db.Model(&models.User{}).Where("email = ?", email).Count(&count).Error; err != nil {
		return false, err
	}
	return count > 0, nil
}

func GetSyluInfoByUUID(uuid int64) (resyluInfo *models.ReqSyluInfo, err error) {
	syluInfo := new(models.SyluUser)
	err = db.Model(&models.SyluUser{}).Where("uuid = ?", uuid).First(&syluInfo).Error

	//教务信息不存在，未绑定，不返回错误
	if err == gorm.ErrRecordNotFound {
		return nil, nil
	}

	//查询出错才返回错误
	if err != nil {
		zap.L().Error("Mysql GetSyluInfoByUUID search Error", zap.Error(err))
		return nil, err
	}

	resyluInfo = &models.ReqSyluInfo{
		ReUsername: syluInfo.ReUsername,
		StudentID:  syluInfo.StudentID,
		Grade:      syluInfo.Grade,
		College:    syluInfo.College,
		Major:      syluInfo.Major,
	}

	return
}

func UpDatePassByEmail(email string, newPassword string) (err error) {
	newPassword = encryptPassword(newPassword)
	return db.Model(&models.User{}).Where("email = ?", email).Update("password", newPassword).Error
}

func CreateOrUpdateSyluPass(userInfo *models.ParamBind, userID int64) (err error) {
	syluPass := &models.SyluPass{
		Uuid:      userID,
		StudentID: userInfo.StudentID,
		Password:  userInfo.Password,
	}
	// 先查询是否存在
	//return db.Where("uuid = ?", syluPass.Uuid).Save(syluPass).Error
	if err := db.Where("uuid = ?", userID).First(&models.SyluPass{}).Error; errors.Is(err, gorm.ErrRecordNotFound) {
		// 记录不存在,执行创建
		return db.Create(syluPass).Error
	}

	// 记录存在,执行更新
	return db.Model(&models.SyluPass{}).Where("uuid = ?", userID).Updates(&syluPass).Error
}

func CreateOrUpdateSyluUser(syluUer *models.SyluUser) (err error) {
	if err := db.Where("uuid = ?", syluUer.Uuid).First(&models.SyluUser{}).Error; errors.Is(err, gorm.ErrRecordNotFound) {
		// 记录不存在,执行创建
		return db.Create(syluUer).Error
	}
	return db.Model(&models.SyluUser{}).Where("uuid = ?", syluUer.Uuid).Updates(&syluUer).Error
}

func GetSyluPassByUuid(uuid int64) (userPass *models.SyluPass, err error) {
	err = db.Model(&models.SyluPass{}).Where("uuid = ?", uuid).First(&userPass).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		err = ErrorUnbound
		return
	}
	return
}

func GetUserInfoByUuid(uuid int64) (userInfo *models.User, err error) {
	userInfo = new(models.User)
	err = db.Model(&models.User{}).Where("uuid = ?", uuid).First(&userInfo).Error
	if err != nil {
		return nil, err
	}

	return
}

func CreatFeed(feedInfo *models.FeedBack) error {
	return db.Create(feedInfo).Error
}

func CreatOrUpdateRecord(record *models.Record) error {
	return db.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "record_id"}},
		UpdateAll: true,
	}).Create(record).Error
}

func DelRecordByID(id int64) error {
	return db.Where("record_id = ?", id).Delete(&models.Record{}).Error
}

func GetRedordsByUserID(userID int64) ([]models.Record, error) {
	var records []models.Record
	if err := db.Model(&models.Record{}).Where("user_id = ?", userID).Order("created_at DESC").Find(&records).Error; err != nil {
		return nil, err
	}

	return records, nil
}
