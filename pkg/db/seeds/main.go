package seeds

import (
	"fmt"

	"github.com/tahsinature/future-proof-gin/pkg/config"
	"github.com/tahsinature/future-proof-gin/pkg/db"
)

func Execute() {
	db.Sync(config.EntryArgs.SyncForce)
	userSeeder := new(UserSeeder)
	userSeeder.CreateOne()

	fmt.Println("Seeding Done...")
}
