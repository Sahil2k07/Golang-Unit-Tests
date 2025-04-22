package mocks

import (
	"github.com/Sahil2k07/Golang-Unit-Tests/models"
	"github.com/stretchr/testify/mock"
)

type UserRepoMock struct {
	mock.Mock
}

func (m *UserRepoMock) AddNewUser(data models.UserDTO) error {
	args := m.Called(data)
	return args.Error(0)
}

func (m *UserRepoMock) GetUserData(email string) (models.User, error) {
	args := m.Called(email)
	return args.Get(0).(models.User), args.Error(1)
}

func (m *UserRepoMock) GetUserWithPosts(email string) (models.User, error) {
	args := m.Called(email)
	return args.Get(0).(models.User), args.Error(1)
}
