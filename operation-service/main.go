package main

import (
	"context"
	"crypto/tls"
	"log"
	"net/http"
	"os"

	"operation-service/config"
	_ "operation-service/docs"
	handler "operation-service/handler/http"
	httpHandler "operation-service/handler/http"
	"operation-service/middleware"
	"operation-service/repository"
	"operation-service/usecase"
	"operation-service/utils"

	userPB "operation-service/pb/user"

	"github.com/go-playground/validator/v10"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	echoMiddleware "github.com/labstack/echo/v4/middleware"
	echoSwagger "github.com/swaggo/echo-swagger"

	// "go.mongodb.org/mongo-driver/v2/mongo"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

// @title Operation Service API
// @version 1.0
// @description This is the documentation of Operation Service API
// @host localhost:8080
// @schemes http https
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

	// ===== GRPC Client Section =====
	userGrpcUrl := os.Getenv("USER_GRPC_SERVICE_URL")
	addressGrpcUrl := os.Getenv("ADDRESS_GRPC_SERVICE_URL")
	recyclehubGrpcUrl := os.Getenv("RECYCLEHUB_GRPC_SERVICE_URL")

	// Set up TLS credentials
	creds := credentials.NewTLS(&tls.Config{
		InsecureSkipVerify: false, // Set to true only for testing with self-signed certificates
	})

	// Dial User gRPC server with TLS
	userConn, err := grpc.NewClient(userGrpcUrl, grpc.WithTransportCredentials(creds))
	if err != nil {
		log.Fatalf("Failed to connect to gRPC server: %v", err)
	}
	defer userConn.Close()

	// Dial Address gRPC server with TLS
	addressConn, err := grpc.NewClient(addressGrpcUrl, grpc.WithTransportCredentials(creds))
	if err != nil {
		log.Fatalf("Failed to connect to gRPC server: %v", err)
	}
	defer addressConn.Close()

	// Dial Recyclehub gRPC server with TLS
	recyclehubConn, err := grpc.NewClient(recyclehubGrpcUrl, grpc.WithTransportCredentials(creds))
	if err != nil {
		log.Fatalf("Failed to connect to gRPC server: %v", err)
	}
	defer recyclehubConn.Close()

	// Create a gRPC client
	userGrpcClient := userPB.NewUserServiceClient(userConn)
	// ===============================

	// # Initialize Echo
	e := echo.New()

	e.Use(echoMiddleware.Recover())
	e.Use(echoMiddleware.CORSWithConfig(echoMiddleware.CORSConfig{
		AllowOrigins:     []string{"https://operation-service.example.com"},
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

	// # Dependency Driver
	driverRoutes := baseRoutes.Group("/drivers")
	driverRepo := repository.NewDriverRepository(db)
	driverUsecase := usecase.NewDriverUsecase(driverRepo, validator)
	driverHttpHandler := httpHandler.NewDriverUsecase(driverUsecase)
	driverHttpHandler.InitRoutes(driverRoutes)

	// # Dependency Injection Order
	orderRoutes := baseRoutes.Group("/orders")
	orderRepo := repository.NewOrderRepository(db)
	orderUsecase := usecase.NewOrderUsecase(orderRepo, validator)
	orderHttpHandler := httpHandler.NewOrderHandler(orderUsecase, userGrpcClient)
	orderHttpHandler.InitRoutes(orderRoutes)

	// # Dependency Injection Order Detail
	orderDetailRoutes := baseRoutes.Group("/order-details")
	orderDetailRepo := repository.NewOrderDetailRepository(db)
	orderDetailUsecase := usecase.NewOrderDetailUsecase(
		orderDetailRepo,
		orderRepo,
		addressConn,
		recyclehubConn,
		recyclehubConn,
		userConn,
		validator,
	)
	orderDetailHttpHandler := handler.NewOrderDetailHandler(orderDetailUsecase, userGrpcClient)
	orderDetailHttpHandler.InitRoutes(orderDetailRoutes)

	// # Swagger documentation route
	e.File("/swagger/doc.json", "docs/swagger.json")
	e.File("/swagger/doc.yaml", "docs/swagger.yaml")
	e.GET("/swagger/*", echoSwagger.WrapHandler)

	// # Swagger documentation route
	e.GET("/swagger/*", echoSwagger.WrapHandler)

	// # Get PORT from ENV
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	// # Log error
	e.Logger.Fatal(e.Start(":" + port))
}
