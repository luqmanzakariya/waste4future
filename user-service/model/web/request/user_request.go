package request

// UserCreateRequest represents the structure of the create user request data.
// @Description UserCreateRequest details.
type UserCreateRequest struct {
	Name      string `validate:"required,min=1,max=200" json:"name" example:"new user"`
	Email     string `validate:"required,min=1,max=200" json:"email" example:"newuser@mail.com"`
	Password  string `validate:"required,min=1,max=200" json:"password" example:"password123"`
	AddressID string `validate:"required,min=1,max=200" json:"address_id" example:""`
}

// UserUpdateRequest represents the structure of the update user request data.
// @Description UserUpdateRequest details.
type UserUpdateRequest struct {
	Name      string `validate:"required,min=1,max=200" json:"name" example:"new user"`
	Email     string `validate:"required,min=1,max=200" json:"email" example:"newuser@mail.com"`
	Password  string `validate:"required,min=1,max=200" json:"password" example:"password123"`
	AddressID string `validate:"required,min=1,max=200" json:"address_id" example:""`
}

// UserLoginRequest represents the structure of the login user request data.
// @Description UserLoginRequest details.
type UserLoginRequest struct {
	Email    string `validate:"required,min=1,max=200" json:"email" example:"luqman@mail.com"`
	Password string `validate:"required,min=1,max=200" json:"password" example:"password123"`
}
