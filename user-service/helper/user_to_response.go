package helper

import (
	"user-service/model/domain"
	"user-service/model/web/response"
)

func UserToResponse(user domain.User) response.UserResponse {
	return response.UserResponse{
		ID:        user.ID,
		Email:     user.Email,
		Name:      user.Name,
		AddressID: user.AddressID,
	}
}

func UserToResponses(users []domain.User) []response.UserResponse {
	var responses []response.UserResponse
	for _, user := range users {
		responses = append(responses, UserToResponse(user))
	}

	return responses
}
