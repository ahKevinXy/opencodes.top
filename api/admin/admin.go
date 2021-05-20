package admin

import "github.com/gin-gonic/gin"

func GoogleAuth(c *gin.Context) {
	id := c.Query("id")
	if id == "" {
		return
	}
}
