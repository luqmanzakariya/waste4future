package domain

import "go.mongodb.org/mongo-driver/bson/primitive"

type WasteType struct {
	ID    primitive.ObjectID `json:"id" bson:"_id"`
	Name  string             `json:"name" bson:"name"`
	Price float64            `json:"price" bson:"price"`
}
