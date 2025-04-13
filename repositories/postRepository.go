package repositories

import (
	"github.com/Sahil2k07/Golang-Unit-Tests/configs"
	"github.com/Sahil2k07/Golang-Unit-Tests/interfaces"
	"gorm.io/gorm"
)

type postRepository struct {
	db *gorm.DB
}

func NewPostRepository() interfaces.IPostRepository {
	return &postRepository{db: configs.DB}
}
