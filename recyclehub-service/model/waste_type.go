package model

import (
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
)

// WasteType represents the structure of a waste type in the database.
// @Description WasteType details.
type WasteType struct {
	ID        bson.ObjectID `bson:"_id,omitempty" json:"id"`
	Name      string        `bson:"name" json:"name"`
	Price     float64       `bson:"price" json:"price"`
	CreatedAt time.Time     `bson:"created_at" json:"created_at"`
	UpdatedAt time.Time     `bson:"updated_at" json:"updated_at"`
}
