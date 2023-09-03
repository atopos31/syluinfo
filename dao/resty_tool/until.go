package resty_tool

import (
	"crypto/rand"
	"crypto/rsa"
	"encoding/base64"
	"math/big"
	"regexp"
	"strconv"
	"strings"
)

// rsa加密
func rsaByPublicKey(password string, publicKey *PublicKey) (string, error) {
	modulusBytes, err := base64.StdEncoding.DecodeString(publicKey.Modulus)
	if err != nil {
		return "", err
	}

	exponentBytes, err := base64.StdEncoding.DecodeString(publicKey.Exponent)
	if err != nil {
		return "", err
	}

	// 解析公钥
	pubKey := &rsa.PublicKey{
		N: new(big.Int).SetBytes(modulusBytes),
		E: int(new(big.Int).SetBytes(exponentBytes).Int64()),
	}

	// 加密密码
	bypassword := []byte(password)
	encryptedBytes, err := rsa.EncryptPKCS1v15(rand.Reader, pubKey, bypassword)
	if err != nil {
		panic(err)
	}

	// Base64 编码加密后的密码
	encryptedPassword := base64.StdEncoding.EncodeToString(encryptedBytes)

	return encryptedPassword, nil
}

func isDegree(degree string) bool {
	if degree == "是" {
		return true
	} else {
		return false
	}
}

func parseWeeks(input string) []int {
	var weeks []int

	ranges := strings.Split(input, ",")
	for _, r := range ranges {
		re := regexp.MustCompile(`(\d+)`)
		bounds := re.FindAllString(r, -1)

		var start int
		var end int
		//有些是1-2周 有些是2周这种 分开看待
		if len(bounds) > 1 {
			start, _ = strconv.Atoi(bounds[0])
			end, _ = strconv.Atoi(bounds[1])
		} else {
			start, _ = strconv.Atoi(bounds[0])
			end = start
		}

		for i := start; i <= end; i++ {
			weeks = append(weeks, i)
		}
	}

	return weeks
}

func timeToInt(time string) (section int, sectionCount int) {
	if len(time) <= 1 {
		var err error
		section, err = strconv.Atoi(time)
		if err != nil {
			section = 0
		}
		sectionCount = 0
		return
	}

	sections := strings.Split(time, "-")
	section, _ = strconv.Atoi(sections[0])
	lastTime, _ := strconv.Atoi(sections[1])
	sectionCount = lastTime - section + 1

	//我觉得应该不会有15-16节吧
	return
}
