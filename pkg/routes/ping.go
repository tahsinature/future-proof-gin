package routes

import (
	"github.com/gin-gonic/gin"
)

func AddPingRoutes(rg *gin.RouterGroup) {
	rg.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong!",
		})
	})
}
