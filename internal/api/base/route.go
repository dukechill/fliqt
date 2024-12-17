package base

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Route(router *gin.Engine) {

	router.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "It works")
	})
}
