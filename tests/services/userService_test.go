package services_test

import (
	"testing"

	"github.com/Sahil2k07/Golang-Unit-Tests/errors"
	"github.com/Sahil2k07/Golang-Unit-Tests/models"
	"github.com/Sahil2k07/Golang-Unit-Tests/services"
	"github.com/Sahil2k07/Golang-Unit-Tests/tests"
	"github.com/Sahil2k07/Golang-Unit-Tests/tests/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestCreateUser_Success(t *testing.T) {
	mockRepo := new(mocks.UserRepoMock)

	reqBody := `{"name":"Shahil","email":"test@example.com"}`
	ctx := tests.SetupContext("POST", "/api/user", reqBody)

	userDTO := models.UserDTO{
		Name:  "Shahil",
		Email: "test@example.com",
	}

	mockRepo.On("GetUserData", "test@example.com").Return(models.User{}, nil)
	mockRepo.On("AddNewUser", userDTO).Return(nil)

	service := services.NewUserService(ctx, mockRepo)
	err := service.CreateNewUser()

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestCreateUser_UserExists(t *testing.T) {
	mockRepo := new(mocks.UserRepoMock)

	reqBody := `{"name":"Shahil","email":"test@example.com"}`
	ctx := tests.SetupContext("POST", "/api/user", reqBody)

	existingUser := models.User{
		ID:    1,
		Name:  "Shahil",
		Email: "test@example.com",
	}

	mockRepo.On("GetUserData", "test@example.com").Return(existingUser, nil)

	service := services.NewUserService(ctx, mockRepo)
	err := service.CreateNewUser()

	assert.EqualError(t, err, "user already exists")
	assert.IsType(t, &errors.CustomError{}, err)

	mockRepo.AssertNotCalled(t, "AddNewUser", mock.Anything)
	mockRepo.AssertExpectations(t)
}
