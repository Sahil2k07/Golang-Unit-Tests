package middlewares

import (
	"net/http"

	"github.com/Sahil2k07/Golang-Unit-Tests/errors"
	"github.com/labstack/echo"
)

func ErrorHandlerMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) (err error) {
		defer func() {
			if r := recover(); r != nil {
				c.Logger().Error("Recovered from panic:", r)
				err = c.JSON(http.StatusInternalServerError, map[string]any{
					"success": false,
					"message": "Internal server error",
				})
			}
		}()

		err = next(c)
		if err == nil {
			return nil
		}

		if customErr, ok := err.(*errors.CustomError); ok {
			return c.JSON(customErr.StatusCode, customErr)
		}

		c.Logger().Error("Middleware caught unexpected error:", err)

		return err
	}
}
