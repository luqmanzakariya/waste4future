package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"

	"recyclehub-service/handler/grpc"
	"recyclehub-service/handler/http"
	recyclehubMiddleware "recyclehub-service/middleware"
	"recyclehub-service/repository"
	"recyclehub-service/usecase"

	pb "recyclehub-service/pb/recyclehub"

	_ "recyclehub-service/docs"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	echoSwagger "github.com/swaggo/echo-swagger"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	grpcLib "google.golang.org/grpc"

	_ "github.com/joho/godotenv/autoload"
)

// @title RecycleHub Service API
// @version 1.0
// @description This is the API documentation for the RecycleHub Service.
// @host https://recyclehub-service-84457363535.asia-southeast2.run.app/
// @schemes https http
// @BasePath /
// @SecurityDefinitions.apikey BearerAuth
// @In header
// @Name Authorization
func main() {
	// MongoDB connection setup
	mongoURI := os.Getenv("MONGO_URI")
	if mongoURI == "" {
		log.Fatal("MONGO_URI environment variable is not set")
	}

	// Create a MongoDB client
	clientOptions := options.Client().ApplyURI(mongoURI)
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatalf("Failed to connect to MongoDB: %v", err)
	}

	// Ping the MongoDB server to verify the connection
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatalf("Failed to ping MongoDB: %v", err)
	}
	fmt.Println("Connected to MongoDB!")

	// Get the MongoDB collection for waste types
	wasteTypeCollection := client.Database(os.Getenv("DB_NAME")).Collection("wastes")
	recycleHubCollection := client.Database(os.Getenv("DB_NAME")).Collection("recycles")

	// Initialize repositories
	wasteTypeRepo := repository.NewWasteTypeRepository(wasteTypeCollection)
	recycleHubRepo := repository.NewRecycleHubRepository(recycleHubCollection)

	// Initialize validator
	validate := validator.New()

	// Initialize use cases
	wasteTypeUseCase := usecase.NewWasteTypeUsecase(wasteTypeRepo)
	recycleHubUseCase := usecase.NewRecycleHubUsecase(recycleHubRepo, validate)

	// Initialize HTTP server
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// swagger documentation route
	e.GET("/swagger/*", echoSwagger.WrapHandler)
	log.Println("Swagger route registered at /swagger/*")

	http.NewWasteTypeHandler(wasteTypeUseCase).RegisterRoutes(e)
	http.NewRecycleHubHandler(recycleHubUseCase).RegisterRoutes(e)

	// Initialize gRPC server with an interceptor (e.g., for authentication)
	grpcServer := grpcLib.NewServer(
		grpcLib.UnaryInterceptor(recyclehubMiddleware.GrpcAuth), // Add your gRPC interceptor here
	)
	pb.RegisterWasteTypeServiceServer(grpcServer, grpc.NewWasteTypeServer(wasteTypeUseCase))
	pb.RegisterRecycleHubServiceServer(grpcServer, grpc.NewRecycleHubServer(recycleHubUseCase))

	// Start gRPC server
	go func() {
		lis, err := net.Listen("tcp", ":50051")
		if err != nil {
			log.Fatalf("Failed to listen for gRPC: %v", err)
		}
		log.Println("gRPC server is running on port 50051")
		if err := grpcServer.Serve(lis); err != nil {
			log.Fatalf("Failed to serve gRPC: %v", err)
		}
	}()

	// Start HTTP server
	log.Println("HTTP server is running on port 8080")
	if err := e.Start(":8080"); err != nil {
		log.Fatalf("Failed to start HTTP server: %v", err)
	}
}
