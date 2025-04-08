package repositories

import (
	"errors"

	"github.com/Sahil2k07/Golang-Unit-Tests/configs"
	"github.com/Sahil2k07/Golang-Unit-Tests/interfaces"
	"github.com/Sahil2k07/Golang-Unit-Tests/models"
	"gorm.io/gorm"
)

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository() interfaces.IUserRepository {
	return &userRepository{db: configs.DB}
}

func (u *userRepository) AddNewUser(data models.UserDTO) error {
	user := &models.User{
		Name:  data.Name,
		Email: data.Email,
	}

	result := u.db.Create(user)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (u *userRepository) GetUserData(email string) (models.User, error) {
	var user models.User

	result := u.db.Where("email = ?", email).Find(&user)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return models.User{}, nil
		}

		return models.User{}, result.Error
	}

	return user, nil
}
