package routes

import "github.com/gin-gonic/gin"

type User struct{}

func (User) setup(rg *gin.RouterGroup) {
	rg.POST("/login", userController.Login)
	rg.POST("/register", userController.Register)
}
