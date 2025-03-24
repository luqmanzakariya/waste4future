package middleware

import (
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
)

func AuthMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		// # Authentication logic
		token := strings.TrimPrefix(c.Request().Header.Get("Authorization"), "Bearer ")
		if token == "" {
			return echo.NewHTTPError(http.StatusUnauthorized, "unauthorized")
		}

		c.Set("token", token)

		// # Call the next handler if authentication is successful
		return next(c)
	}
}
