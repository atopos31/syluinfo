package sendemail

import (
	"cld/settings"
	"errors"
	"fmt"
	"math/rand"
	"strconv"
	"time"

	"gopkg.in/gomail.v2"
)

var (
	ErrorSendFalied = errors.New("验证码发送失败！")
)

// 随机验证码生成
func GetCode() string {
	rand.Seed(time.Now().UnixNano())
	code := rand.Intn(899999) + 100000
	res := strconv.Itoa(code) //转字符串返回
	return res
}

// 发送验证码
func SendEmail(mode string, email string, code string) (err error) {
	m := gomail.NewMessage()
	m.SetAddressHeader("From", settings.Conf.Email.Username, "hackerxiao")
	m.SetHeader("To", email)
	m.SetHeader("Subject", "[我的验证码]"+mode)
	m.SetBody("text/html", getBody(email, code))

	config := gomail.NewDialer(
		settings.Conf.Email.Host,
		settings.Conf.Email.Port,
		settings.Conf.Email.Username,
		settings.Conf.Email.Password,
	)

	if err := config.DialAndSend(m); err != nil {
		return ErrorSendFalied
	}

	return
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
