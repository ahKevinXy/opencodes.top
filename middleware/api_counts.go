package middleware

import (
	"opencodes/common/consant"
	"opencodes/pkg/redis"
	"time"

	"github.com/gin-gonic/gin"
)

func APiCounts() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 获取 判断是否是post 请求

		// 获取限定次数
		apiUrl := c.Request.URL.Path

		if apiUrl != "" {
			Counts(apiUrl)
		}

	}
}

func Counts(url string) {
	t := time.Now().Format("2006_01_02")
	redis.HIncrBy(consant.APICOUNTS+t, url, 1)
}
