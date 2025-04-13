package main

import (
	"github.com/Sahil2k07/Golang-Unit-Tests/configs"
	"github.com/Sahil2k07/Golang-Unit-Tests/controllers"
	"github.com/Sahil2k07/Golang-Unit-Tests/middlewares"
	"github.com/labstack/echo"
)

func init() {
	configs.InitDatabase()
}

func main() {
	e := echo.New()

	e.Use(middlewares.ErrorHandlerMiddleware)

	controllers.InitControllers(e)

	e.Logger.Fatal(e.Start(":1324"))
}
