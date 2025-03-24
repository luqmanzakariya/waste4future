package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"

	"address-service/config"
	// _ "address-service/docs"
	grpcHandler "address-service/handler/grpc"
	httpHandler "address-service/handler/http"
	"address-service/middleware"
	"address-service/repository"
	"address-service/usecase"
	"address-service/utils"

	pb "address-service/pb/address"

	"github.com/go-playground/validator/v10"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	echoMiddleware "github.com/labstack/echo/v4/middleware"
	"github.com/soheilhy/cmux"
	echoSwagger "github.com/swaggo/echo-swagger"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"google.golang.org/grpc"
)

// @title Address Service API
// @version 1.0
// @description This is the documentation of Address Service API
// @host https://address-service-84457363535.asia-southeast2.run.app/
// @schemes https http
// @BasePath /
// @SecurityDefinitions.apikey BearerAuth
// @In header
// @Name Authorization
func main() {
	// # Load ENV
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found, relying on system environment variables")
	}

	// # Initialize MongoDB
	ctx := context.Background()
	mongoClient, db := config.InitMongoDB(ctx)
	defer func() {
		if err := mongoClient.Disconnect(ctx); err != nil {
			log.Printf("Failed to disconnect MongoDB client: %v", err)
		}
	}()

	sigCh := make(chan os.Signal, 1)
	errCh := make(chan error, 1)
	quitCh := make(chan bool, 1)
	signal.Notify(sigCh, os.Interrupt)

	go func() {
		for {
			select {
			case <-sigCh:
				quitCh <- true
			case err := <-errCh:
				log.Fatal(err)
				quitCh <- true
			}
		}
	}()

	// Start the combined HTTP/gRPC server
	go func() {
		StartCombinedServer(db, errCh)
	}()

	<-quitCh
	fmt.Println("exiting program...")
}

func StartCombinedServer(db *mongo.Database, errCh chan error) {
	// # Get PORT from ENV
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	// Create a TCP listener
	lis, err := net.Listen("tcp", ":"+port)
	if err != nil {
		errCh <- fmt.Errorf("failed to listen: %v", err)
		return
	}

	// Create a cmux multiplexer
	m := cmux.New(lis)

	// Match HTTP requests
	httpL := m.Match(cmux.HTTP1Fast())
	// Match gRPC requests
	grpcL := m.Match(cmux.Any())

	// # Initialize Echo
	e := echo.New()

	e.Use(echoMiddleware.Recover())
	e.Use(echoMiddleware.CORSWithConfig(echoMiddleware.CORSConfig{
		AllowOrigins:     []string{"https://address-service.example.com"},
		AllowMethods:     []string{echo.GET, echo.POST, echo.PUT, echo.DELETE, echo.OPTIONS},
		AllowHeaders:     []string{"Content-Type", "Authorization", "X-Requested-With"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))

	// Ensure correct CORS response for preflight requests
	e.OPTIONS("/*", func(c echo.Context) error {
		return c.NoContent(http.StatusNoContent)
	})

	// # Initialize validator for payload struct validation
	validator := validator.New()

	// # Built-in middleware
	e.Pre(middleware.TrailingSlashMiddleware)
	e.Use(middleware.WithLogger)

	// # Error handler
	e.HTTPErrorHandler = utils.HTTPErrorHandler

	// # Base Routes
	baseRoutes := e.Group("/api")

	// # Dependency Injection Address
	addressRoutes := baseRoutes.Group("/addresses")
	addressRepo := repository.NewAddressRepository(db)
	addressUsecase := usecase.NewAddressUsecase(addressRepo, validator)
	addressHTTPHandler := httpHandler.NewAddressHandler(addressUsecase)
	addressHTTPHandler.InitRoutes(addressRoutes)

	// # Swagger documentation route
	e.File("/swagger/doc.json", "docs/swagger.json")
	e.File("/swagger/doc.yaml", "docs/swagger.yaml")
	e.GET("/swagger/*", echoSwagger.WrapHandler)

	// Initialize gRPC server
	grpcServer := grpc.NewServer(grpc.UnaryInterceptor(middleware.AuthGRPCInterceptor))
	pb.RegisterAddressServiceServer(grpcServer, grpcHandler.NewAddressGRPCHandler(addressUsecase))

	// Start the HTTP server
	go func() {
		if err := e.Server.Serve(httpL); err != nil {
			errCh <- fmt.Errorf("HTTP server error: %v", err)
		}
	}()

	// Start the gRPC server
	go func() {
		if err := grpcServer.Serve(grpcL); err != nil {
			errCh <- fmt.Errorf("gRPC server error: %v", err)
		}
	}()

	// Start the cmux listener
	fmt.Printf("Combined HTTP/gRPC server started on port %s\n", port)
	if err := m.Serve(); err != nil {
		errCh <- fmt.Errorf("cmux serve error: %v", err)
	}
}
