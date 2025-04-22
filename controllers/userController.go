package controllers

import (
	"github.com/Sahil2k07/Golang-Unit-Tests/repositories"
	"github.com/Sahil2k07/Golang-Unit-Tests/services"
	"github.com/labstack/echo"
)

func userController(e *echo.Echo) {
	e.GET("/api/user", func(c echo.Context) error {
		us := services.NewUserService(c, repositories.NewUserRepository())

		return us.GetUserData()
	})

	e.POST("/api/user", func(c echo.Context) error {
		us := services.NewUserService(c, repositories.NewUserRepository())

		return us.CreateNewUser()
	})

	e.GET("/api/user/post", func(c echo.Context) error {
		us := services.NewUserService(c, repositories.NewUserRepository())

		return us.GetUserWithPosts()
	})
}
