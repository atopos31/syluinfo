package emailsm

import (
	"cld/settings"
	"errors"
	"fmt"
	"time"

	"go.uber.org/zap"
	"gopkg.in/gomail.v2"
)

var (
	ErrorSendFalied = errors.New("验证码发送失败！")
)

type myEmailDia struct {
	address string
	*gomail.Dialer
}

func NewEmailDialer() *myEmailDia {
	cfg := settings.Conf.Email
	max := len(cfg.EmailPass)
	index := randomInt(max)

	dialer := gomail.NewDialer(
		cfg.Host,
		cfg.Port,
		cfg.EmailPass[index].Username,
		cfg.EmailPass[index].Password,
	)

	return &myEmailDia{
		address: cfg.EmailPass[index].Username,
		Dialer:  dialer,
	}
}

// 发送验证码
func (myED *myEmailDia) SendEmail(mode string, email string) (code string, err error) {
	m := gomail.NewMessage()
	m.SetAddressHeader("From", myED.address, "hackerxiao")
	m.SetHeader("To", email)
	m.SetHeader("Subject", "[we校园小助手]"+mode)

	code = getCode()
	m.SetBody("text/html", getBody(email, code))

	if err := myED.DialAndSend(m); err != nil {
		zap.L().Error("emailDia.SendEmail(TitleSign, email)", zap.Error(err))
		return "", ErrorSendFalied
	}

	return code, nil
}

// 获取邮件主体
func getBody(email string, code string) (body string) {
	now := time.Now()
	time := fmt.Sprintf("%02d-%02d-%02d %02d:%02d:%02d", now.Year(), now.Month(), now.Day(), now.Hour(), now.Minute(), now.Second())
	body = fmt.Sprintf(`<div>
		<div>
			尊敬的%s，您好！
		</div>
		<div style="padding: 8px 40px 8px 50px;">
			<p>您于 %s 提交的邮箱验证，本次验证码为 %s，为了保证账号安全，验证码有效期为5分钟。请确认为本人操作，切勿向他人泄露，感谢您的理解与使用。</p>
		</div>
		<div>
			<p>此邮箱为系统邮箱，请勿回复。</p>
		</div>	
	</div>`, email, time, code)

	return
}
