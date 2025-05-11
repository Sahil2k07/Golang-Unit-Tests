package controllers

import (
	"github.com/Sahil2k07/Golang-Unit-Tests/repositories"
	"github.com/Sahil2k07/Golang-Unit-Tests/services"
	"github.com/labstack/echo"
)

func postController(e *echo.Echo) {
	e.GET("/api/post/user", func(c echo.Context) error {
		ps := services.NewPostService(c, repositories.NewPostRepository())

		return ps.GetUserPosts()
	})
}
