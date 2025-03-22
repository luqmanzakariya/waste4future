package middleware

import (
	"address-service/utils"
	"context"
	"fmt"

	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

type contextKey string

const (
	UserInfoKey contextKey = "userInfo" // Key to store userInfo in context
)

func AuthGRPCInterceptor(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (any, error) {
	// Extract metadata from the incoming context
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, fmt.Errorf("missing metadata")
	}

	// Check if the authorization header is present
	if len(md["authorization"]) == 0 {
		return nil, fmt.Errorf("missing token")
	}

	// Extract the token from the authorization header
	token := md["authorization"][0]

	// Validate the JWT token
	userInfo, err := utils.ValidateJWT(token)
	if err != nil {
		return nil, fmt.Errorf("invalid token: %v", err)
	}

	// Store the userInfo in the context
	ctx = context.WithValue(ctx, UserInfoKey, userInfo)

	// Call the next handler
	return handler(ctx, req)
}
