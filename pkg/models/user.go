package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Email    string
	Password string
	Name     string
}

func (User) TableName() string {
	return "users"
}
