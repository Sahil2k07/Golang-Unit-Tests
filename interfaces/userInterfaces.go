package interfaces

import "github.com/Sahil2k07/Golang-Unit-Tests/models"

type (
	IUserService interface {
		CreateNewUser() error
		GetUserData() error
		GetUserWithPosts() error
	}

	IUserRepository interface {
		AddNewUser(data models.UserDTO) error
		GetUserData(email string) (models.User, error)
		GetUserWithPosts(email string) (models.User, error)
	}
)
