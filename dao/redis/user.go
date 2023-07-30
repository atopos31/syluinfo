package redis

import (
	"cld/settings"
	"errors"
	"time"

	"go.uber.org/zap"
)

var (
	ErrorFrequentSend = errors.New("请勿频繁操作")
	ErrorNotExists    = errors.New("验证码不存在！")
)

func ExitCodeTimeOut(keyMode string, email string) (err error) {
	if rdb.Exists(ctx, keyMode+email+KeyTimeout).Val() > 0 {
		return ErrorFrequentSend
	}
	return nil
}

func SetTimeOutAndCode(keyMode, email string, code string) (err error) {

	if err := rdb.Set(ctx, keyMode+email+KeyTimeout, KeyCheck, time.Minute).Err(); err != nil {
		return err
	}
	if err := rdb.Set(ctx, keyMode+email, code, time.Duration(settings.Conf.Email.Expires)*time.Second).Err(); err != nil {
		return err
	}

	return
}

func GetCaptchaByEmail(keyMode, email string) (Ocaptcha string, err error) {
	Ocaptcha, err = rdb.Get(ctx, keyMode+email).Result()
	if err != nil {
		return "", ErrorNotExists
	}
	return
}

func DelCodeByEmail(KeyMode, email string) {
	err := rdb.Del(ctx, KeyMode+email).Err()
	zap.L().Error(" rdb.Del(ctx,KeyEmail+email).Err()", zap.Error(err), zap.String("email", email))
}
