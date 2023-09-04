package redis

import (
	"cld/settings"
	"context"
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
}
