package main

import (
	"context"
	"fmt"
	"os"

	"crypto/tls"

	_ "github.com/joho/godotenv/autoload"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/metadata"

	pb "hacktiv/pb/user"
)

func main() {
	// Set up TLS credentials
	creds := credentials.NewTLS(&tls.Config{
		InsecureSkipVerify: false, // Set to true only for testing with self-signed certificates
	})

	opts := []grpc.DialOption{
		grpc.WithTransportCredentials(creds),
		grpc.WithUnaryInterceptor(WithAuth),
	}

	conn, err := grpc.NewClient(os.Getenv("USER_SERVICE_GRPC_URL"), opts...)
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	userServiceClient := pb.NewUserServiceClient(conn)
	users, err := userServiceClient.Validate(context.Background(), &pb.ValidateRequest{})
	fmt.Println(users, err)
}

func WithAuth(ctx context.Context, method string, req, reply any, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
	newCtx := metadata.AppendToOutgoingContext(ctx, "authorization", "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhZGRyZXNzX2lkIjoiMSIsImVtYWlsIjoibmV3dXNlckBtYWlsLmNvbSIsImV4cCI6MTc0MjYwNTU2MSwiaWQiOjQsIm5hbWUiOiJuZXcgdXNlciJ9.cZoCLZqCa7JCB22lb1Vn0kAkLx2wHiCQ4Q8Vt-yFPXU")
	return invoker(newCtx, method, req, reply, cc, opts...)
}
