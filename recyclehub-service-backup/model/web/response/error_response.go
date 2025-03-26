package response

// ErrorResponse represents the response payload for an error.
// @Description ErrorResponse represents an error message.
type ErrorResponse struct {
	Message string `json:"message"` // Error message
}
