package repositories

import (
	"errors"

	"github.com/tahsinature/future-proof-gin/pkg/db"
	"github.com/tahsinature/future-proof-gin/pkg/forms"
	"github.com/tahsinature/future-proof-gin/pkg/models"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type UserRepository struct{}

func (m UserRepository) GetModel() models.User {
	return models.User{}
}

func (m UserRepository) CheckUserExistsByEmail(email string) (bool, error) {
	var user models.User
	db := db.GetDB()

	err := db.First(&user, "email = ?", email).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return false, nil
	}

	if err != nil {
		return false, err
	}

	return true, nil
}

func (m UserRepository) GetUserByEmail(email string) (models.User, error) {
	var user models.User
	db := db.GetDB()

	err := db.First(&user, "email = ?", email).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return user, errors.New("user not found")
	}

	if err != nil {
		return user, err
	}

	return user, nil
}

func (m UserRepository) Register(form forms.Register) (user models.User, err error) {
	db := db.GetDB()
	if exists, err := m.CheckUserExistsByEmail(form.Email); err != nil {
		return user, err
	} else if exists {
		return user, errors.New("user already exists")
	}

	bytePassword := []byte(form.Password)
	hashedPassword, err := bcrypt.GenerateFromPassword(bytePassword, bcrypt.DefaultCost)
	if err != nil {
		return user, errors.New("something went wrong, please try again later")
	}

	user.Password = string(hashedPassword)
	user.Name = form.Name
	user.Email = form.Email
	err = db.Create(&user).Error

	if err != nil {
		return user, errors.New("something went wrong, please try again later")
	}

	return user, err
}
