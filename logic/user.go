package logic

import (
	"cld/cos"
	"cld/dao/mysql"
	"cld/dao/redis"
	"cld/models"
	"cld/pkg/jwt"
	"cld/pkg/sendemail"
	"cld/pkg/snowflake"
	"cld/settings"
	"errors"

	"go.uber.org/zap"
)

var (
	ModeSign    = "sign"
	ModeReset   = "resetpass"
	ModeRecover = "recoverpass"
)

var (
	TitleSign            = "邮箱注册"
	TitleRecoverPassword = "找回密码"
)

var (
	ErrorFailedCaptcha = errors.New("验证码错误！")
	ErrorEmailExist    = errors.New("该邮箱已被注册！")
	ErrorEmailNotExist = errors.New("该邮箱尚未注册！")
)

func Login(loginReq *models.ParamLogin) (data *models.ReqLogin, err error) {
	user, err := mysql.Login(loginReq)
	if err != nil {
		return nil, err
	}
	token, err := jwt.GenToken(user.Uuid, user.Email)
	if err != nil {
		return nil, err
	}
	//此时已经登陆成功了
	data = &models.ReqLogin{
		Username: user.Username,
		Email:    user.Email,
		Token:    token,
	}
	//获取教务信息
	syluinfo, err := mysql.GetSyluInfoByUUID(user.Uuid)
	if err != nil {
		return
	}

	data.SyluInfo = syluinfo
	return
}

func SignUp(signReq *models.ParamSignUp) (err error) {
	Ocaptcha, err := redis.GetCaptchaByEmail(redis.KeyModeSignUp, signReq.Email)
	if err != nil {
		return err
	}

	if Ocaptcha != signReq.Captcha {
		return ErrorFailedCaptcha
	}
	//获取基于雪花算法生成的uuid
	uuid := snowflake.GenID()

	userInfo := &models.User{
		Uuid:     uuid,
		Username: signReq.Username,
		Password: signReq.Password,
		Email:    signReq.Email,
	}

	if err := mysql.CreateUser(userInfo); err != nil {
		return err
	}
	//删除验证码防止多次注册！
	redis.DelCodeByEmail(redis.KeyModeSignUp, signReq.Email)

	return nil
}

func SendEmail(queryEmail models.QuerySendEmail) (err error) {
	if queryEmail.Mode == ModeSign {
		err = SignUpSendEmail(queryEmail.Email)
	} else if queryEmail.Mode == ModeRecover {
		err = RecoverSendEmail(queryEmail.Email)
	}

	return
}

func SignUpSendEmail(email string) (err error) {
	ok, err := mysql.CheckEmailExist(email)
	if err != nil {
		zap.L().Error("mysql.CheckEmailExist(email)", zap.Error(err), zap.String("email", email))
		return err
	}

	if ok {
		return ErrorEmailExist
	}

	if err := redis.ExitCodeTimeOut(redis.KeyModeSignUp, email); err != nil {
		return err
	}
	//获取6位随机验证码
	code := sendemail.GetCode()
	//发送验证码
	if err := sendemail.SendEmail(TitleSign, email, code); err != nil {
		return err
	}
	//设置超时和检验缓存
	if err := redis.SetTimeOutAndCode(redis.KeyModeSignUp, email, code); err != nil {
		return err
	}

	return
}

func RecoverSendEmail(email string) (err error) {
	ok, err := mysql.CheckEmailExist(email)
	if err != nil {
		zap.L().Error("mysql.CheckEmailExist(email)", zap.Error(err), zap.String("email", email))
		return err
	}

	if !ok {
		return ErrorEmailNotExist
	}

	if err := redis.ExitCodeTimeOut(redis.KeyModeRecover, email); err != nil {
		return err
	}

	//获取6位随机验证码
	code := sendemail.GetCode()
	//发送验证码
	if err := sendemail.SendEmail(TitleRecoverPassword, email, code); err != nil {
		return err
	}
	//设置超时和检验缓存
	if err := redis.SetTimeOutAndCode(redis.KeyModeRecover, email, code); err != nil {
		return err
	}

	return
}

func ReSetPass(parmReSet *models.ParamReSet) (err error) {
	loginTest := &models.ParamLogin{
		Email:    parmReSet.Email,
		Password: parmReSet.Password,
	}

	if _, err := mysql.Login(loginTest); err != nil {
		return err
	}

	if err := mysql.UpDatePassByEmail(parmReSet.Email, parmReSet.NewPassword); err != nil {
		return err
	}

	return
}

func ReCoverPass(parmReCover *models.ParamReCover) (err error) {
	Ocaptcha, err := redis.GetCaptchaByEmail(redis.KeyModeRecover, parmReCover.Email)
	if err != nil {
		return err
	}

	if Ocaptcha != parmReCover.Captcha {
		return ErrorFailedCaptcha
	}

	if err := mysql.UpDatePassByEmail(parmReCover.Email, parmReCover.NewPassword); err != nil {
		return err
	}

	redis.DelCodeByEmail(redis.KeyModeRecover, parmReCover.Email)

	return
}

func GetCosKey(cfg *settings.CosConfig) (resKey *models.ResCosKey, err error) {
	res, err := cos.GetKey()
	if err != nil {
		return nil, err
	}

	resKey = &models.ResCosKey{
		Bucket:        cfg.Resource.Bucket,
		Region:        cfg.Resource.Region,
		AllowsPath:    cfg.Resource.AllowPath,
		TmpSecretId:   res.Credentials.TmpSecretID,
		TmpSecretKey:  res.Credentials.TmpSecretKey,
		SecurityToken: res.Credentials.SessionToken,
		StartTime:     res.StartTime,
		ExpiredTime:   res.ExpiredTime,
	}

	return
}
