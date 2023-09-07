package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/juju/ratelimit"
)

func RateLimitMiddleware(fillInterval float64, cap int64) func(c *gin.Context) {
	//速率/容量
	bucket := ratelimit.NewBucketWithRate(fillInterval, cap)
	return func(c *gin.Context) {
		// 如果取不到令牌就中断本次请求返回 rate limit...
		if bucket.TakeAvailable(1) < 1 {
			c.String(http.StatusBadGateway, "rate limit...")
			c.Abort()
			return
		}
		c.Next()
	}
}
