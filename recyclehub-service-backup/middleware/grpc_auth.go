package middleware

import (
	"context"
	"log"

	"google.golang.org/grpc"
)

// GrpcAuth is an example gRPC interceptor for authentication
func GrpcAuth(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	// Add your authentication logic here
	log.Println("gRPC interceptor: checking authentication")

	// Example: Check for a token in the context
	// token := ExtractTokenFromContext(ctx)
	// if !IsValidToken(token) {
	//     return nil, status.Errorf(codes.Unauthenticated, "invalid token")
	// }

	// Call the next handler
	return handler(ctx, req)
}
