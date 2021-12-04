package application

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/tahsinature/future-proof-gin/pkg/config"
	"github.com/tahsinature/future-proof-gin/pkg/db"
	"github.com/tahsinature/future-proof-gin/pkg/routes"
)

type Application struct{}

func (*Application) Setup() *gin.Engine {
	config.Validate()
	gin.SetMode(gin.ReleaseMode)

	db.Init()
	db.InitRedis()

	engine := routes.Setup()

	return engine
}

func (*Application) Listen(engine *gin.Engine) {
	fmt.Println("Server started on port:", config.App.Port)
	if err := engine.Run(":" + config.App.Port); err != nil {
		panic(err)
	}
}
