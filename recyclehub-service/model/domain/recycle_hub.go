package domain

import "go.mongodb.org/mongo-driver/bson/primitive"

// RecycleHub represents a recycle hub
type RecycleHub struct {
	ID          primitive.ObjectID `json:"id" bson:"_id"`
	Name        string             `json:"name" bson:"name"`
	Phone       string             `json:"phone" bson:"phone"`
	AddressID   string             `json:"address_id" bson:"address_id"`
	WasteTypeID string             `json:"waste_type_id" bson:"waste_type_id"`
}
