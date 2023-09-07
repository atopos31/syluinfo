package emailsm

import (
	"math/rand"
	"strconv"
	"time"
)

// 随机验证码生成
func getCode() string {
	rand.Seed(time.Now().UnixNano())
	code := rand.Intn(899999) + 100000
	res := strconv.Itoa(code) //转字符串返回
	return res
}

func randomInt(max int) int {
	if max <= 0 {
		return 0
	}
	rand.Seed(time.Now().UnixNano()) // 使用当前时间作为随机种子
	return rand.Intn(max)            // 生成0到max-1之间的随机整数
}
