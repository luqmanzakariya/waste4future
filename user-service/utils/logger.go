package utils

import (
	"time"

	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
)

func LogEntry(c echo.Context) *logrus.Entry {
	log := logrus.NewEntry(logrus.StandardLogger())
	log.Logger.SetLevel(logrus.DebugLevel)
	if c == nil {
		return log
	}

	return log.WithFields(logrus.Fields{
		"at":         time.Now().Format("2006-01-02 15:04:05"),
		"method":     c.Request().Method,
		"body":       c.Request().Body,
		"request_id": c.Request().Header.Get("X-Request-ID"),
		"ip":         c.RealIP(),
		"user_agent": c.Request().UserAgent(),
		"path":       c.Path(),
	})
}
