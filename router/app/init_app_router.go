package app

import "github.com/gin-gonic/gin"

func InitAppRouter(r *gin.Engine, version string) {
	InitUserRouter(r, version)
	InitOrderRouter(r, version)
}
