package interfaces

import "github.com/Sahil2k07/Golang-Unit-Tests/models"

type (
	IPostRepository interface {
		GetUserPosts(email string) ([]models.Post, error)
	}

	IPostService interface {
		GetUserPosts() error
	}
)
