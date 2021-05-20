package order

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type Name struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

//retries int
//	types   string
//	status  int
//	params  interface{}
// Deprecated
func Create(c *gin.Context) {

	time.Sleep(10 * time.Second)

	c.JSON(http.StatusOK, 1)
}
