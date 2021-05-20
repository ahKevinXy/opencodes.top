package app

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-contrib/requestid"

	"github.com/gin-gonic/gin"
)

var (
	baseUserUrl = "/app/user/"
)

func InitUserRouter(r *gin.Engine, version string) {
	// 无需认证路由
	UserNoCheckRoleRouter(r, version)

}
func UserNoCheckRoleRouter(r *gin.Engine, version string) {

	router := r.Group(version + baseUserUrl)

	{
		router.GET("info", func(c *gin.Context) {

			fmt.Println(requestid.Get(c), "------")
			c.SetCookie("gin_cookie", time.Now().Format("2006-01-02 15-04-05"), 3600, "/", "localhost", false, true)
			c.JSON(http.StatusOK, gin.H{"ok": 1})
		})

	}
}
