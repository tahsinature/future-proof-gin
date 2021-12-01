package models

import "gorm.io/gorm"

type Article struct {
	gorm.Model
	UserID  int64
	Title   string
	Content string
}

func (a *Article) TableName() string {
	return "article"
}
