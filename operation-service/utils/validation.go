package utils

import (
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

// ValidationError wraps validator.ValidationErrors into a proper HTTP error
func ValidationError(err error) error {
	if ve, ok := err.(validator.ValidationErrors); ok {
		errors := make(map[string]string)
		for _, fe := range ve {
			errors[fe.Field()] = msgForTag(fe.Tag())
		}
		return echo.NewHTTPError(echo.ErrBadRequest.Code, map[string]interface{}{
			"message": "Validation failed",
			"errors":  errors,
		})
	}
	return echo.NewHTTPError(echo.ErrBadRequest.Code, err.Error())
}

// msgForTag returns human-readable messages for validation tags
func msgForTag(tag string) string {
	switch tag {
	case "required":
		return "This field is required"
	case "email":
		return "Invalid email format"
	case "gt":
		return "Value must be greater than specified"
	case "gte":
		return "Value must be greater than or equal to specified"
	case "lt":
		return "Value must be less than specified"
	case "lte":
		return "Value must be less than or equal to specified"
	case "oneof":
		return "Value must be one of the allowed options"
	default:
		return "Validation failed for this field"
	}
}
