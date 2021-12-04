package main

import (
	"github.com/alexflint/go-arg"
	"github.com/tahsinature/future-proof-gin/pkg/application"
	"github.com/tahsinature/future-proof-gin/pkg/config"
	"github.com/tahsinature/future-proof-gin/pkg/db/seeds"
)

func main() {
	arg.MustParse(&config.EntryArgs)

	if config.EntryArgs.Run {
		application := new(application.Application)
		engine := application.Setup()
		application.Listen(engine)
	}

	if config.EntryArgs.Seed {
		seeds.Execute()
	}
}
