package controllers

import (
	"github.com/Sahil2k07/Golang-Unit-Tests/errors"
	"github.com/labstack/echo"
)

func postController(e *echo.Echo) {
	e.PUT("/api/test", func(c echo.Context) error {
		// data := struct {
		// 	Success bool
		// 	Message string
		// }{
		// 	Success: false,
		// 	Message: "Test message",
		// }
		// return c.JSON(http.StatusOK, data)

		return errors.BadRequestError("testing bad request")
	})
}
