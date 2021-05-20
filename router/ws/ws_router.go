package ws

import (
	"opencodes/pkg/ws"

	"github.com/gin-gonic/gin"
)

var (
	baseUserUrl = "ws"
)

func InitWsRouter(r *gin.Engine, version string) {
	// 无需认证路由
	NoCheckRoleRouter(r, version)

}
func NoCheckRoleRouter(r *gin.Engine, version string) {

	router := r.Group(baseUserUrl)

	{
		router.GET("/:channel", ws.WebsocketManager.WsClient)
	}
}
