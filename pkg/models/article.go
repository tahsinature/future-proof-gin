package models

import "gorm.io/gorm"

type Article struct {
	gorm.Model
	UserID  int64
	Title   string
	Content string
}

func (Article) TableName() string {
	return "articles"
}
