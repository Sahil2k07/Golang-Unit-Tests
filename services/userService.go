package services

import (
	"net/http"

	"github.com/Sahil2k07/Golang-Unit-Tests/errors"
	"github.com/Sahil2k07/Golang-Unit-Tests/interfaces"
	"github.com/Sahil2k07/Golang-Unit-Tests/models"
	"github.com/labstack/echo"
)

type userService struct {
	ctx  echo.Context
	repo interfaces.IUserRepository
}

func NewUserService(c echo.Context, repo interfaces.IUserRepository) interfaces.IUserService {
	return &userService{
		ctx:  c,
		repo: repo,
	}
}

func (u *userService) CreateNewUser() error {
	var dto models.UserDTO

	err := u.ctx.Bind(&dto)
	if err != nil {
		return err
	}

	existingUser, err := u.repo.GetUserData(dto.Email)
	if err != nil {
		return err
	}

	if existingUser.Email != "" {
		return errors.BadRequestError("user already exists")
	}

	return u.repo.AddNewUser(dto)
}

func (u *userService) GetUserData() error {
	email := u.ctx.QueryParam("email")

	user, err := u.repo.GetUserData(email)
	if err != nil {
		return err
	}

	if user.Email == "" {
		return errors.NotFoundError("user not found")
	}

	return u.ctx.JSON(http.StatusOK, user)
}
