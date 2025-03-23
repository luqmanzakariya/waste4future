package model

import (
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
)

// Driver Status Enum
type DriverStatus string

const DriverStatusActive DriverStatus = "active"
const DriverStatusInactive DriverStatus = "inactive"

type Driver struct {
	ID           bson.ObjectID `bson:"_id,omitempty" json:"id"`
	Name         string        `bson:"name" json:"name"`
	Phone        string        `bson:"phone" json:"phone"`
	LicensePlate string        `bson:"license_plate" json:"license_plate"`
	Status       DriverStatus  `bson:"status" json:"status"`
	CreatedAt    time.Time     `bson:"created_at" json:"created_at"`
	UpdatedAt    time.Time     `bson:"updated_at" json:"updated_at"`
}

// PayloadCreateDriver represents the structure of the create driver request data.
// @Description PayloadCreateDriver details.
type PayloadCreateDriver struct {
	Name         string `json:"name" example:"driver grab"`
	Phone        string `json:"phone" example:"081394859283"`
	LicensePlate string `json:"license_plate" example:"B 4134 LUZ"`
}
