package tool

import (
	"strconv"
	"time"
)

// 获取毫秒级时间戳
func NowTime() string {
	timestamp := time.Now().UnixNano() / 1000000
	return strconv.FormatInt(timestamp, 10)
}
