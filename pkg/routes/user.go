package routes

import "github.com/gin-gonic/gin"

func AddUserRoutes(rg *gin.RouterGroup) {
	rg.POST("/login", userController.Login)
	rg.POST("/register", userController.Register)
}
