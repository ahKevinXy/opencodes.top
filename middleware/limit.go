package middleware

import (
	"net/http"
	"opencodes/common/consant"
	"opencodes/pkg/redis"
	"opencodes/tools"
	"time"

	"github.com/gin-gonic/gin"
)

func LimitRequest() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 获取 判断是否是post 请求
		if c.Request.Method == "POST" {

			// 限定接口
		}
		// 获取限定次数
		if getApiLimit(c.Request.URL.Path) {
			c.JSON(http.StatusOK, gin.H{
				"code": -1,
				"msg":  "请求超过限制",
			})
			c.Abort()
			return
		}

		//redis.IncrBy("limit_____"+c.Request.URL.Path, 1)
	}
}

// 获取接口限定
// true 有限制 false 无限制
func getApiLimit(url string) bool {
	limits, err := redis.HGet(consant.APILIMT, url)
	// 如果获取为 为空直接 返回false
	if err != nil && limits == "" {
		return false
	}
	// 获取当前限制次数
	key := consant.APILIMT + "_" + url
	res, err := redis.IncrBy(key, 1)
	if err != nil {
		return false
	}
	limit, _ := tools.StringToInt64(limits)
	if res > limit {
		return true
	}
	// 如果第一次设置过期时间
	if res == 1 {
		err := redis.SetExpire(key, time.Second*10)
		if err != nil {
			return false
		}
	}
	return false
}
