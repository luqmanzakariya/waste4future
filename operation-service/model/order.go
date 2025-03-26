package model

import (
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
)

// Order Status Enum
type OrderStatus string

const OrderStatusDraft OrderStatus = "draft"
const OrderStatusPending OrderStatus = "pending"
const OrderStatusPaid OrderStatus = "paid"
const OrderStatusRejected OrderStatus = "rejected"

// Shipping Status Enum
type ShippingStatus string

const ShippingStatusUnassigned ShippingStatus = "unassigned"
const ShippingStatusPickup ShippingStatus = "pickup"
const ShippingStatusDelivery ShippingStatus = "delivery"
const ShippingStatusFinish ShippingStatus = "finish"

type OrderDetailIDs []string

type Order struct {
	ID              bson.ObjectID  `bson:"_id,omitempty" json:"id"`
	UserID          int64          `bson:"user_id" json:"user_id"`
	DriverID        string         `bson:"driver_id" json:"driver_id"`
	OrderDetailIDs  []string       `bson:"order_detail_ids" json:"order_detail_ids"`
	OrderDate       time.Time      `bson:"order_date" json:"order_date"`
	OrderStatus     OrderStatus    `bson:"order_status" json:"order_status"`
	ShippingStatus  ShippingStatus `bson:"shipping_status" json:"shipping_status"`
	UpdatedShipping time.Time      `bson:"updated_shipping" json:"updated_shipping"`
	Note            string         `bson:"note" json:"note"`
	CreatedAt       time.Time      `bson:"created_at" json:"created_at"`
	UpdatedAt       time.Time      `bson:"updated_at" json:"updated_at"`
}

// PayloadCreateOrder represents the structure of the create order request data.
// @Description PayloadCreateOrder details.
type PayloadCreateOrder struct {
	Note string `json:"note" example:"mohon hati hati karena terdapat barang pecah belah"`
}

// PayloadUpdateOrder represents the structure of the update order request data.
// @Description PayloadUpdateOrder details.
type PayloadUpdateOrder struct {
	OrderDetailIDs []string `bson:"order_detail_ids"`
	Note           string   `json:"note" example:"40.7128"`
}

// ResponseOrder represents the structure order response data.
// @Description ResponseOrder details.
type ResponseOrder struct {
	ID              string         `json:"id"`
	DriverID        string         `json:"driver_id"`
	OrderDetailIDs  []string       `json:"order_detail_ids"`
	OrderDate       time.Time      `json:"order_date"`
	OrderStatus     OrderStatus    `json:"order_status"`
	ShippingStatus  ShippingStatus `json:"shipping_status"`
	UpdatedShipping time.Time      `json:"updated_shipping"`
	Note            string         `json:"note"`
	CreatedAt       time.Time      `json:"created_at"`
	UpdatedAt       time.Time      `json:"updated_at"`
}
