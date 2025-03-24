package model

import (
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
)

type RecycleHub struct {
	ID          bson.ObjectID `bson:"_id,omitempty" json:"id"`
	Name        string        `bson:"name" json:"name"`
	Phone       string        `bson:"phone" json:"phone"`
	AddressID   string        `bson:"address_id" json:"address_id"`
	WasteTypeID string        `bson:"waste_type_id" json:"waste_type_id"`
	CreatedAt   time.Time     `bson:"created_at" json:"created_at"`
	UpdatedAt   time.Time     `bson:"updated_at" json:"updated_at"`
}

type PayloadCreateRecycleHub struct {
	Name        string `json:"name" validate:"required" example:"Green Hub"`
	Phone       string `json:"phone" validate:"required" example:"123-456-7890"`
	AddressID   string `json:"address_id" validate:"required" example:"67cdcb62a50a990a870d928f"`
	WasteTypeID string `json:"waste_type_id" validate:"required" example:"67cdcb62a50a990a870d928f"`
}

type PayloadUpdateRecycleHub struct {
	Name        string `json:"name" validate:"required" example:"Green Hub Updated"`
	Phone       string `json:"phone" validate:"required" example:"123-456-7890"`
	AddressID   string `json:"address_id" validate:"required" example:"67cdcb62a50a990a870d928f"`
	WasteTypeID string `json:"waste_type_id" validate:"required" example:"67cdcb62a50a990a870d928f"`
}

type ResponseRecycleHub struct {
	ID          string    `json:"id"`
	Name        string    `json:"name"`
	Phone       string    `json:"phone"`
	AddressID   string    `json:"address_id"`
	WasteTypeID string    `json:"waste_type_id"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
