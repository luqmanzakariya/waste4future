package utils

import (
	// "operation-service/handler"
	// "operation-service/usecase"
	"strconv"

	"github.com/labstack/echo/v4"
)

type HTTPError struct {
	ErrorCode string      `json:"error_code"`
	Message   interface{} `json:"message"`
	Detail    interface{} `json:"detail,omitempty"`
}

var ErrorMap = map[string]string{
	// handler.ErrUserIsNotSuperadmin: "ERR_001",
	// handler.ErrRoleIsNotRegistered: "ERR_002",
	// handler.ErrImageIsTooLarge:     "ERR_001",

	// usecase.ErrSpecificInUsecase: "ERR_001",
	// ... akan terus bertambah sesuai error baru yang muncul
}

func HTTPErrorHandler(err error, c echo.Context) {
	httpError, ok := err.(*echo.HTTPError)
	if !ok {
		httpError = echo.ErrInternalServerError
	}

	errorCode := strconv.Itoa(httpError.Code)

	// customCode, ok := ErrorMap[httpError.Message.(string)]
	// if !ok {
	// 	httpError = echo.ErrInternalServerError
	// }

	errObj := HTTPError{
		ErrorCode: errorCode,
		Message:   httpError.Message,
	}

	//...
	// Bisa utilize map
	LogEntry(c).Error(httpError)
	// LogEntry(c).Warn(httpError)
	// LogEntry(c).Debug(httpError)

	c.JSON(httpError.Code, errObj)
}
