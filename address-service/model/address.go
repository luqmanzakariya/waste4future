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

// PayloadCreateAddress represents the structure of the create address request data.
// @Description PayloadCreateAddress details.
type PayloadCreateAddress struct {
	Name      string `json:"name" validate:"required" example:"Home"`
	Latitude  string `json:"latitude" validate:"required" example:"40.7128"`
	Longitude string `json:"longitude" validate:"required" example:"-74.0060"`
}

// PayloadUpdateAddress represents the structure of the update address request data.
// @Description PayloadUpdateAddress details.
type PayloadUpdateAddress struct {
	Name      string `json:"name" validate:"required" example:"Office"`
	Latitude  string `json:"latitude" validate:"required" example:"40.7128"`
	Longitude string `json:"longitude" validate:"required" example:"-74.0060"`
}

// ResponseAddress represents the structure address response data.
// @Description ResponseAddress details.
type ResponseAddress struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	Latitude  string    `json:"latitude"`
	Longitude string    `json:"longitude"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
