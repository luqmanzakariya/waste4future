package grpcHandler

import (
	"context"
	"fmt"
	"user-service/middleware"
	userPB "user-service/pb/user"
	"user-service/usecase"
	"user-service/utils"
)

type server struct {
	userPB.UnimplementedUserServiceServer
	userUsecase usecase.IUserUsecase
}

func NewUserGrpcServer(userUsecase usecase.IUserUsecase) userPB.UserServiceServer {
	return &server{
		userUsecase: userUsecase,
	}
}

func (s *server) Validate(ctx context.Context, in *userPB.ValidateRequest) (*userPB.User, error) {
	// Retrieve userInfo from context
	userInfo, ok := ctx.Value(middleware.UserInfoKey).(*utils.UserInfo)
	if !ok || userInfo == nil {
		return nil, fmt.Errorf("unauthorized: user info not found")
	}

	// Map the user to the gRPC response
	return &userPB.User{
		ID:        int64(userInfo.ID),
		Name:      userInfo.Name,
		Email:     userInfo.Email,
		AddressID: userInfo.AddressID,
	}, nil
}
