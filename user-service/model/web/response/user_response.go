package response

// UserResponse represents the structure of user response data.
// @Description UserResponse details.
type UserResponse struct {
	ID        int    `validate:"required,min=1,max=200" json:"id"`
	Email     string `validate:"required,min=1,max=200" json:"email"`
	Name      string `validate:"required,min=1,max=200" json:"name"`
	AddressID string `validate:"required,min=1,max=200" json:"address_id"`
}
