package app

import (
	"net/http"
	"opencodes/api/app/v1/order"

	"github.com/gin-gonic/gin"
)

var (
	baseOrderUrl = "/app/order/"
)

func InitOrderRouter(r *gin.Engine, version string) {
	// 无需认证路由
	OrderNoCheckRoleRouter(r, version)

}
func OrderNoCheckRoleRouter(r *gin.Engine, version string) {

	router := r.Group(version + baseOrderUrl)

	{
		router.GET("info", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{"ok": 1})
		})
		router.POST("log", order.Create)
	}
}
