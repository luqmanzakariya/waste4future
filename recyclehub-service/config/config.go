package config

import (
	"os"

	_ "github.com/joho/godotenv/autoload"
	"go.mongodb.org/mongo-driver/mongo"
)

var Client *mongo.Client

// GetCollection returns a MongoDB collection by name
func GetCollection(collectionName string) *mongo.Collection {
	return Client.Database(os.Getenv("DB_NAME")).Collection(collectionName)
}
