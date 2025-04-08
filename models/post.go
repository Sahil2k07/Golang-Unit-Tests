package models

import "gorm.io/gorm"

type (
	Post struct {
		gorm.Model
		ID     uint   `gorm:"primaryKey" json:"id"`
		Title  string `json:"title"`
		Body   string `json:"body"`
		UserID uint   `json:"userId"`
		User   User   `json:"user"`
	}

	PostDTO struct {
		Title string `json:"title"`
		Body  string `json:"body"`
	}
)
