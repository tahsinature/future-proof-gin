package main

import (
	"github.com/alexflint/go-arg"
	"github.com/gin-gonic/gin"
	"github.com/tahsinature/future-proof-gin/pkg/config"
	"github.com/tahsinature/future-proof-gin/pkg/db"
	"github.com/tahsinature/future-proof-gin/pkg/db/seeds"
	"github.com/tahsinature/future-proof-gin/pkg/routes"
)

var args struct {
	Run  bool `arg:"-r,help:Run the server"`
	Seed bool `arg:"-s,help:Seed the database"`
}

func main() {
	arg.MustParse(&args)
	config.Validate()
	gin.SetMode(gin.ReleaseMode)

	db.Init()
	db.SyncForce()
	db.InitRedis()

	if args.Run {
		engine := routes.Setup()

		engine.Run(":" + config.App.Port)
	}

	if args.Seed {
		seeds.Execute()
	}
}
