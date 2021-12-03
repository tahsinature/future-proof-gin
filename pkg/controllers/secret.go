package controllers

import "github.com/gin-gonic/gin"

type Secret struct{}

func (Secret) Get(c *gin.Context) {
	c.JSON(200, map[string]interface{}{
		"message": "Secret Page",
	})
}
