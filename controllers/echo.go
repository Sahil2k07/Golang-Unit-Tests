package controllers

import "github.com/labstack/echo"

func InitControllers(e *echo.Echo) {
	// All the user controllers
	userController(e)

	// All the post controllers
	postController(e)
}
