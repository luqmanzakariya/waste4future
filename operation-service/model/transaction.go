package model

import (
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
)

// PaymentMethod Enum
type PaymentMethod string

const (
	PaymentMethodCOD          PaymentMethod = "cod"
	PaymentMethodBankTransfer PaymentMethod = "bank_transfer"
)

// PaymentStatus Enum
type PaymentStatus string

const (
	PaymentStatusPending   PaymentStatus = "pending"
	PaymentStatusCompleted PaymentStatus = "completed"
	PaymentStatusRejected  PaymentStatus = "rejected"
)

type Transaction struct {
	ID              bson.ObjectID `bson:"_id,omitempty" json:"id"`
	OrderID         string        `bson:"order_id" json:"order_id"`
	PaymentMethod   PaymentMethod `bson:"payment_method" json:"payment_method"`
	GrandTotal      float64       `bson:"grand_total" json:"grand_total"`
	PaymentStatus   PaymentStatus `bson:"payment_status" json:"payment_status"`
	TransactionDate time.Time     `bson:"transaction_date" json:"transaction_date"`
	CreatedAt       time.Time     `bson:"created_at" json:"created_at"`
	UpdatedAt       time.Time     `bson:"updated_at" json:"updated_at"`
}

// PayloadCreateTransaction represents the structure of the create transaction request data.
// @Description PayloadCreateTransaction details.
type PayloadCreateTransaction struct {
	OrderID       string        `json:"order_id" validate:"required"`
	PaymentMethod PaymentMethod `json:"payment_method" validate:"required,oneof=cod bank_transfer"`
	GrandTotal    float64       `json:"grand_total" validate:"required,gt=0"`
}

// PayloadUpdateTransaction represents the structure of the update transaction request data.
// @Description PayloadUpdateTransaction details.
type PayloadUpdateTransaction struct {
	PaymentMethod PaymentMethod `json:"payment_method" validate:"omitempty,oneof=cod bank_transfer"`
	GrandTotal    float64       `json:"grand_total" validate:"omitempty,gt=0"`
	PaymentStatus PaymentStatus `json:"payment_status" validate:"omitempty,oneof=pending completed rejected"`
}

// ResponseTransaction represents the structure transaction response data.
// @Description ResponseTransaction details.
type ResponseTransaction struct {
	ID              string        `json:"id"`
	OrderID         string        `json:"order_id"`
	PaymentMethod   PaymentMethod `json:"payment_method"`
	GrandTotal      float64       `json:"grand_total"`
	PaymentStatus   PaymentStatus `json:"payment_status"`
	TransactionDate time.Time     `json:"transaction_date"`
	CreatedAt       time.Time     `json:"created_at"`
	UpdatedAt       time.Time     `json:"updated_at"`
}
