package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"user-service/config"
	_ "user-service/docs" // docs is generated by Swag CLI, you have to import it.
	grpcHandler "user-service/handler/grpc"
	httpHandler "user-service/handler/http"
	"user-service/middleware"
	"user-service/repository"
	"user-service/usecase"
	"user-service/utils"

	userPB "user-service/pb/user"

	"github.com/go-playground/validator/v10"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	echoMiddleware "github.com/labstack/echo/v4/middleware"
	"github.com/soheilhy/cmux" // Add cmux
	echoSwagger "github.com/swaggo/echo-swagger"
	"google.golang.org/grpc"
	"gorm.io/gorm"
)

// @title Waste4Future User Service API
// @version 1.0
// @description This is the documentation of Waste4Future User Service API
// @host localhost:8080
// @schemes http https
// @BasePath /
// @SecurityDefinitions.apikey BearerAuth
// @In header
// @Name Authorization
func main() {
	// # Load ENV
	godotenv.Load()

	// # Initialize Database
	db := config.InitDB()

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

func StartCombinedServer(db *gorm.DB, errCh chan error) {
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
		AllowOrigins:     []string{"https://user-service-84457363535.asia-southeast2.run.app"},
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

	// # built-in middleware
	e.Pre(middleware.TrailingSlashMiddleware) // Custom middleware to remove trailing slash
	e.Use(middleware.WithLogger)              // Logger middleware

	// # Error handler
	e.HTTPErrorHandler = utils.HTTPErrorHandler

	// # Base Routes
	baseRoutes := e.Group("/api")

	// # Dependency Injection User
	userRoutes := baseRoutes.Group("/users")
	userRepo := repository.NewUserRepository(db)
	userUsecase := usecase.NewUserUsecase(userRepo, validator)
	userHandler := httpHandler.NewUserUsecase(userUsecase)
	userHandler.InitRoutes(userRoutes)

	// # Swagger documentation route
	e.GET("/swagger/*", echoSwagger.WrapHandler)

	// Initialize gRPC server
	grpcServer := grpc.NewServer(grpc.UnaryInterceptor(middleware.GrpcAuth))
	userPB.RegisterUserServiceServer(grpcServer, grpcHandler.NewUserGrpcServer(userUsecase))

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

	// # Log error
	// e.Logger.Fatal(e.Start(":" + port))
}
