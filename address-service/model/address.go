package model

import (
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
)

type Address struct {
	ID        bson.ObjectID `bson:"_id,omitempty" json:"id"`
	Name      string        `bson:"name" json:"name"`
	Latitude  string        `bson:"latitude" json:"latitude"`
	Longitude string        `bson:"longitude" json:"longitude"`
	CreatedAt time.Time     `bson:"created_at" json:"created_at"`
	UpdatedAt time.Time     `bson:"updated_at" json:"updated_at"`
}
