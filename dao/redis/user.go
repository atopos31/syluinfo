package redis

import (
	"cld/models"
	"cld/settings"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"go.uber.org/zap"
)

var (
	ErrorFrequentSend = errors.New("请勿频繁操作")
	ErrorNotExists    = errors.New("验证码不存在！")
	ErrorGrade        = errors.New("成绩不存在")
)

func ExitCodeTimeOut(keyMode string, email string) (err error) {
	if rdb.Exists(ctx, keyMode+email+KeyTimeout).Val() > 0 {
		return ErrorFrequentSend
	}
	return nil
}

func SetTimeOutAndCode(keyMode, email string, code string) (err error) {

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	pipe := rdb.Pipeline()
	expire := time.Duration(settings.Conf.Email.Expires) * time.Second

	key1 := keyMode + email + KeyTimeout
	pipe.Set(ctx, key1, KeyCheck, time.Minute)
	pipe.Expire(ctx, key1, time.Minute)

	key2 := keyMode + email
	pipe.Set(ctx, key2, code, expire)
	pipe.Expire(ctx, key2, expire)

	_, err = pipe.Exec(ctx)
	if err != nil {
		return err
	}

	return nil
}

func GetCaptchaByEmail(keyMode, email string) (Ocaptcha string, err error) {
	Ocaptcha, err = rdb.Get(ctx, keyMode+email).Result()
	if err != nil {
		return "", ErrorNotExists
	}
	return
}

func DelCodeByEmail(KeyMode, email string) {
	if err := rdb.Del(ctx, KeyMode+email).Err(); err != nil {
		zap.L().Error(" rdb.Del(ctx,KeyEmail+email).Err()", zap.Error(err), zap.String("email", email))
	}

	return
}

func GetGradeDetail(uuid int64, classId string) (gradeDetail []*models.ResGradeDetail, err error) {
	key := fmt.Sprintf("%s%s:%d", KeyGradeDetail, classId, uuid)
	value, err := rdb.Get(ctx, key).Result()
	if err != nil {
		return nil, ErrorNotExists
	}

	gradeDetail = *new([]*models.ResGradeDetail)
	json.Unmarshal([]byte(value), &gradeDetail)

	return
}

func SetGradeDetail(uuid int64, classId string, gradeDetail []*models.ResGradeDetail) error {
	key := fmt.Sprintf("%s%s:%d", KeyGradeDetail, classId, uuid)
	gradeByte, err := json.Marshal(gradeDetail)
	if err != nil {
		return err
	}
	value := string(gradeByte)
	if err := rdb.Set(ctx, key, value, 7*24*time.Hour).Err(); err != nil {
		return err
	}
	return nil
}
