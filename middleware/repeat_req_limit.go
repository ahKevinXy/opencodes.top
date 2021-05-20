package middleware

import (
	"fmt"
	"opencodes/pkg/jwt"
	"opencodes/pkg/redis"
	"opencodes/tools"
	"time"

	"github.com/gin-gonic/gin"
)

// 限制重复提交
func RepeatReqLimit() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 限定 POST 接口
		if c.Request.Method != "POST" {
			return
		}
		// 获取uid
		uid := jwt.GetUid(c)
		if uid == 0 {
			return
		}
		// 判断是否存在
		//redis_key = uid + request uri
		uri := c.Request.URL.Path
		redisKey := tools.IntToString(uid) + "_" + uri
		// 判断键是否存在
		res, err := redis.Get(redisKey)
		fmt.Println(res, "-----", err, redisKey, time.Now().Format("2006-01-02 15:04:05"))
		if err != nil {
			return
		}
		// 重复提交
		if res != "" {
			c.JSON(400, gin.H{
				"name": "重复提交小伙子",
			})

			c.Abort()
			return
		}
		// 设置请求时间 3秒 内 不能重复提交
		redis.Set(redisKey, 1, 30*time.Second)

		c.Next()
		fmt.Printf("结束:%+v", time.Now().Format("2006-01-02 15:04:05"))
		redis.Del(redisKey)
		return
	}
}
