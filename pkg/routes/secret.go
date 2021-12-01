package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/tahsinature/future-proof-gin/pkg/middlewares"
)

func AddSecretRoutes(rg *gin.RouterGroup) {
	rg.Use(middlewares.Auth)
	rg.GET("/", secretController.Get)
}
