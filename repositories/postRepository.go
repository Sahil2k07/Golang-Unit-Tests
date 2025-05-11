package repositories

import (
	"github.com/Sahil2k07/Golang-Unit-Tests/configs"
	"github.com/Sahil2k07/Golang-Unit-Tests/interfaces"
	"github.com/Sahil2k07/Golang-Unit-Tests/models"
	"gorm.io/gorm"
)

type postRepository struct {
	db *gorm.DB
}

func NewPostRepository() interfaces.IPostRepository {
	return &postRepository{db: configs.DB}
}

func (pr *postRepository) GetUserPosts(email string) ([]models.Post, error) {
	var user models.User

	err := pr.db.Preload("Posts").Where("email = ?", email).First(&user).Error
	if err != nil {
		return nil, err
	}

	return user.Posts, nil
}
