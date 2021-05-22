package middleware

import (
	"context"
	"strings"

	"github.com/labstack/echo/v4"

	firebase "firebase.google.com/go"
)

func Auth() echo.MiddlewareFunc {
	return auth
}

func auth(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		app, err := firebase.NewApp(context.Background(), nil)
		if err != nil {
			return err
		}

		client, err := app.Auth(context.Background())
		if err != nil {
			return err
		}

		auth := c.Request().Header.Get("Authorization")
		idToken := strings.Replace(auth, "Bearer ", "", 1)
		token, err := client.VerifyIDToken(context.Background(), idToken)
		if err != nil {
			return err
		}

		c.Set("token", token)
		return next(c)
	}
}
