package routes

import (
	"github.com/gin-gonic/gin"
)

type Ping struct{}

func (Ping) setup(rg *gin.RouterGroup) {
	rg.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong!",
		})
	})
}
