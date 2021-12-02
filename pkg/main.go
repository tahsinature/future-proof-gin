package main

import (
	"github.com/gin-gonic/gin"
	"github.com/tahsinature/future-proof-gin/pkg/config"
	"github.com/tahsinature/future-proof-gin/pkg/db"
	"github.com/tahsinature/future-proof-gin/pkg/routes"
)

func main() {
	config.Validate()

	if config.App.IsProduction {
		gin.SetMode(gin.ReleaseMode)
	}

	r := routes.Setup()

	db.Init()
	db.SyncForce()
	db.InitRedis(1)

	r.Run(":" + config.App.Port)
}
