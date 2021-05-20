package middleware

import (
	"net/http"
	"opencodes/pkg/jwt"

	"github.com/gin-gonic/gin"
)

// JWTAuth 中间件，检查token
func JWTAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		var err error
		token := c.Request.Header.Get("token")
		if token == "" {
			c.JSON(http.StatusOK, gin.H{
				"code": -1,
				"msg":  "请求未携带token，无权限访问",
			})
			c.Abort()
			return
		}
		// parseToken 解析token包含的信息
		claims, err := jwt.ParseToken(token)

		if err != nil {
			if err == jwt.TokenExpired {
				c.JSON(http.StatusOK, gin.H{
					"code": -1,
					"msg":  "授权已过期",
				})
				c.Abort()
				return
			}
			c.JSON(http.StatusOK, gin.H{
				"code": -1,
				"msg":  "授权失败",
			})
			c.Abort()
			return
		}
		// 解析ID
		id := jwt.GetUid(c)
		if id == 0 {
			c.JSON(http.StatusOK, gin.H{
				"code": -1,
				"msg":  "授权已过期",
			})
			c.Abort()
			return
		}

		// 继续交由下一个路由处理,并将解析出的信息传递下去
		c.Set("claims", claims)

	}
}
