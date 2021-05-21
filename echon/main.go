package main

import (
	"net/http"

	"echon/model"
	"echon/utils"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!12")
	})

	e.GET("/:param", func(c echo.Context) error {
		param := c.Param("param")
		return c.String(http.StatusOK, param)
	})

	e.GET("/api/v1/messages", func(c echo.Context) error {
		message := model.NewMessage("type", "message")
		return c.JSON(200, utils.ConventionalMarshaller{Value: message})
	})

	e.Logger.Fatal(e.Start(":" + "8080"))
}
