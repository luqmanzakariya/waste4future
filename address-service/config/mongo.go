package config

import (
	"context"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

func InitMongoDB(ctx context.Context) (*mongo.Client, *mongo.Database) {
	clientOpts := options.Client().ApplyURI(os.Getenv("MONGODB_URI")).
		SetMaxPoolSize(10).
		SetMinPoolSize(1).
		SetMaxConnIdleTime(60 * time.Second)
	client, err := mongo.Connect(clientOpts)
	if err != nil {
		panic(err)
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		panic(err)
	}

	db := client.Database(os.Getenv("MONGODB_DATABASE"))
	_, err = db.Collection("ping").Find(ctx, bson.D{})
	if err != nil {
		panic(err)
	}

	return client, db
}
