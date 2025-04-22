package mocks

import (
	"github.com/Sahil2k07/Golang-Unit-Tests/models"
	"github.com/stretchr/testify/mock"
)

type PostRepoMock struct {
	mock.Mock
}

func (m *PostRepoMock) GetUserPosts(email string) ([]models.Post, error) {
	args := m.Called(email)
	return args.Get(0).([]models.Post), args.Error(1)
}
