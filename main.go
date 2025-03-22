package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"

	"recyclehub-service/handler/grpc"
	"recyclehub-service/handler/http"
	"recyclehub-service/repository"
	"recyclehub-service/usecase"

	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"google.golang.org/grpc"

	_ "github.com/joho/godotenv/autoload"
)

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

	// Initialize use cases
	wasteTypeUseCase := usecase.NewWasteTypeUseCase(wasteTypeRepo)
	recycleHubUseCase := usecase.NewRecycleHubUseCase(recycleHubRepo)

	// Initialize HTTP server
	e := echo.New()
	http.NewWasteTypeHandler(wasteTypeUseCase).RegisterRoutes(e)
	http.NewRecycleHubHandler(recycleHubUseCase).RegisterRoutes(e)

	// Initialize gRPC server
	grpcServer := grpc.NewServer()
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
