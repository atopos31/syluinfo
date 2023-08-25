package jwt

import (
	"cld/settings"
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type MyCustomClaims struct {
	UserID int64  `json:"user_id"`
	Email  string `json:"email"`
	jwt.RegisteredClaims
}

// 创建Token
func GenToken(userID int64, email string) (string, error) {
	//创建一个声明的Token
	c := MyCustomClaims{
		UserID: userID,
		Email:  email,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * time.Duration(settings.Conf.Jwt.Timeout))),
			Issuer:    settings.Conf.Jwt.Issuer,
		},
	}
	//使用指定的签名方法
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, c)
	//使用指定的secret签名获取完整的编码后的字符串Token
	return token.SignedString([]byte(settings.Conf.Jwt.Secret))
}

// 解析Token
func ParseToken(tokenstring string) (*MyCustomClaims, error) {
	var mc = new(MyCustomClaims)
	token, err := jwt.ParseWithClaims(tokenstring, mc, func(t *jwt.Token) (interface{}, error) {
		return []byte(settings.Conf.Jwt.Secret), nil
	})

	if err != nil {
		return nil, err
	}
	//验证Token
	if token.Valid {
		return mc, nil
	}

	return nil, errors.New("invalid token")
}
