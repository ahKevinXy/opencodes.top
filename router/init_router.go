package router

import (
	"opencodes/middleware"
	"opencodes/router/app"
	"opencodes/router/swagger"
	"opencodes/router/ws"

	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/pprof"
	"github.com/gin-contrib/requestid"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	r := gin.New()
	// 初始化 gin-swagger

	// 注册中间件 todo
	swagger.InitRouter(r)
	r.Use(requestid.New())
	r.Use(middleware.JWTAuth())
	r.Use(middleware.LimitRequest())
	r.Use(middleware.RepeatReqLimit())
	r.Use(middleware.APiCounts())
	r.Use(middleware.Session())
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://google.com"}
	pprof.Register(r)

	r.Use(cors.New(config))
	app.InitAppRouter(r, "")

	ws.InitWsRouter(r, "")

	return r
}
