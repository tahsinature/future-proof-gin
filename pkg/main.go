package main

import (
	"fmt"

	"github.com/alexflint/go-arg"
	"github.com/gin-gonic/gin"
	"github.com/tahsinature/future-proof-gin/pkg/config"
	"github.com/tahsinature/future-proof-gin/pkg/db"
	"github.com/tahsinature/future-proof-gin/pkg/db/seeds"
	"github.com/tahsinature/future-proof-gin/pkg/routes"
)

func main() {
	arg.MustParse(&config.EntryArgs)
	config.Validate()
	gin.SetMode(gin.ReleaseMode)

	db.Init()
	db.InitRedis()

	if config.EntryArgs.Run {
		engine := routes.Setup()

		fmt.Println("Server started on port:", config.App.Port)
		err := engine.Run(":" + config.App.Port)
		if err != nil {
			panic(err)
		}
	}

	if config.EntryArgs.Seed {
		seeds.Execute()
	}
}
