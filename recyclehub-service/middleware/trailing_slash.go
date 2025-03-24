package middleware

import (
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
)

// TrailingSlashMiddleware removes the trailing slash from URLs
func TrailingSlashMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		req := c.Request()
		if strings.HasSuffix(req.URL.Path, "/") && req.URL.Path != "/" {
			// Redirect to path without trailing slash
			newPath := strings.TrimSuffix(req.URL.Path, "/")
			return c.Redirect(http.StatusMovedPermanently, newPath)
		}
		return next(c)
	}
}
