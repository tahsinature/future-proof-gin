package seeds

import (
	"github.com/tahsinature/future-proof-gin/pkg/db"
	"github.com/tahsinature/future-proof-gin/pkg/models"
)

type UserSeeder struct{}

func (UserSeeder) CreateOne() *models.User {
	_, db := db.GetDB()
	userModel := models.User{
		Email:    "john@mail.com",
		Name:     "John Doe",
		Password: "password",
	}
	err := db.Create(&userModel).Error
	if err != nil {
		panic(err)
	}

	return &userModel
}

func (u UserSeeder) CreateMany(count int) []models.User {
	users := make([]models.User, count)

	for i := 0; i < count; i++ {
		users[i] = *u.CreateOne()
	}

	return users
}
