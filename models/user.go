package models

import "gorm.io/gorm"

type (
	User struct {
		gorm.Model
		Name  string `json:"name"`
		Email string `gorm:"unique" json:"email"`
		Posts []Post `json:"posts"`
	}

	UserDTO struct {
		Name  string `json:"name"`
		Email string `gorm:"unique" json:"email"`
	}
)
