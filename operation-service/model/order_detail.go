package model

import (
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
)

type OrderDetail struct {
	ID                   bson.ObjectID `bson:"_id,omitempty" json:"id"`
	UserID               int           `bson:"user_id" json:"user_id"`
	RecycleHubID         string        `bson:"recycle_hub_id" json:"recycle_hub_id"`
	WasteWeight          float64       `bson:"waste_weight" json:"waste_weight"`
	SubTotal             float64       `bson:"sub_total" json:"sub_total"`
	OriginAddressID      string        `bson:"origin_address_id" json:"origin_address_id"`
	DestinationAddressID string        `bson:"destination_address_id" json:"destination_address_id"`
	DeliveryPrice        float64       `bson:"delivery_price" json:"delivery_price"`
	CreatedAt            time.Time     `bson:"created_at" json:"created_at"`
	UpdatedAt            time.Time     `bson:"updated_at" json:"updated_at"`
}

type PayloadCreateOrderDetail struct {
	RecycleHubID         string  `json:"recycle_hub_id" validate:"required"`
	WasteWeight          float64 `json:"waste_weight" validate:"required"`
	OriginAddressID      string  `json:"origin_address_id" validate:"required"`
	DestinationAddressID string  `json:"destination_address_id" validate:"required"`
}

type PayloadUpdateOrderDetail struct {
	WasteWeight          float64 `json:"waste_weight" validate:"required"`
	OriginAddressID      string  `json:"origin_address_id" validate:"required"`
	DestinationAddressID string  `json:"destination_address_id" validate:"required"`
}

type ResponseOrderDetail struct {
	ID                   string    `json:"id"`
	UserID               int       `json:"user_id"`
	RecycleHubID         string    `json:"recycle_hub_id"`
	WasteWeight          float64   `json:"waste_weight"`
	SubTotal             float64   `json:"sub_total"`
	OriginAddressID      string    `json:"origin_address_id"`
	DestinationAddressID string    `json:"destination_address_id"`
	DeliveryPrice        float64   `json:"delivery_price"`
	CreatedAt            time.Time `json:"created_at"`
	UpdatedAt            time.Time `json:"updated_at"`
}
