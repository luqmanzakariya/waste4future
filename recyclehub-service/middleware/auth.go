package middleware

import (
	"net/http"
	"reyclehub-service/utils"
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

		id, err := utils.ValidateJWT(token)
		if err != nil {
			return echo.NewHTTPError(http.StatusUnauthorized, "Unauthorized")
		}

		c.Set("id", id)

		// # Call the next handler if authentication is successful
		return next(c)
	}
}
