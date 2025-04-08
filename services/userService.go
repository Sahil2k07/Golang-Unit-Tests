package services

import (
	"net/http"

	"github.com/Sahil2k07/Golang-Unit-Tests/interfaces"
	"github.com/Sahil2k07/Golang-Unit-Tests/models"
	"github.com/Sahil2k07/Golang-Unit-Tests/repositories"
	"github.com/labstack/echo"
)

type userService struct {
	ctx  echo.Context
	repo interfaces.IUserRepository
}

func NewUserService(c echo.Context) interfaces.IUserService {
	return &userService{
		ctx:  c,
		repo: repositories.NewUserRepository(),
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
		return u.ctx.String(http.StatusBadRequest, "User already exists")
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
		return u.ctx.JSON(http.StatusNotFound, "no user found")
	}

	return u.ctx.JSON(http.StatusOK, user)
}
