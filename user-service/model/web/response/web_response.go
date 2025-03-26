package response

// WebResponse represents the standard API response format.
// @Description represents the standard API response format.
type WebResponse struct {
	Code   int         `json:"code"`
	Status string      `json:"status"`
	Data   interface{} `json:"data"`
}

// TokenResponse represents the standard token response format.
// @Description represents the standard token response format.
type TokenResponse struct {
	Token string `json:"token"`
}
