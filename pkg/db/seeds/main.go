package seeds

import (
	"fmt"

	"github.com/tahsinature/future-proof-gin/pkg/db"
)

func Execute() {
	db.SyncForce()
	userSeeder := new(UserSeeder)
	userSeeder.CreateOne()

	fmt.Println("Seeding Done...")
}
