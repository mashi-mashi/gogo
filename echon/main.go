package main

import (
	"net/http"

	custom "echon/api"
	"echon/lib"
	"echon/model"
	"echon/utils"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())

	apiV1 := e.Group("/api/v1")

	apiV1.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	}, custom.Auth())

	apiV1.GET("/:param", func(c echo.Context) error {
		param := c.Param("param")
		addData := lib.AddDoc("test", map[string]interface{}{
			"name":  param,
			"param": param,
		})
		return c.JSON(http.StatusOK, utils.ConventionalMarshaller{Value: addData})
	})

	apiV1.GET("/messages", func(c echo.Context) error {
		message := model.NewMessage("type", "message")
		return c.JSON(200, utils.ConventionalMarshaller{Value: message})
	})

	e.Logger.Fatal(e.Start(":" + "8080"))
}
