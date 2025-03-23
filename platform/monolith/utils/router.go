package utils

import (
	"github.com/gin-gonic/gin"
	"github.com/go-kit/log"
	"net/http"
)




func NewRouter() *gin.Engine {
	r := gin.New()

	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status": "ok",
		})
	})

	return r
}
