package controllers

import "github.com/gin-gonic/gin"

type SecretController struct{}

func (SecretController) Get(c *gin.Context) {
	c.JSON(200, map[string]interface{}{
		"message": "Secret Page",
	})
}
