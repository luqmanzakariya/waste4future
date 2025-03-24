package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Transaction represents a payment transaction in the system
type Transaction struct {
	ID              primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	OrderID         string             `bson:"order_id" json:"order_id"`
	PaymentMethod   string             `bson:"payment_method" json:"payment_method"`
	GrandTotal      float64            `bson:"grand_total" json:"grand_total"`
	PaymentStatus   string             `bson:"payment_status" json:"payment_status"`
	TransactionDate time.Time          `bson:"transaction_date" json:"transaction_date"`
	CreatedAt       time.Time          `bson:"created_at" json:"created_at"`
	UpdatedAt       time.Time          `bson:"updated_at" json:"updated_at"`
}

// TransactionStatus defines possible payment status values
type TransactionStatus string

const (
	TransactionStatusPending   TransactionStatus = "pending"
	TransactionStatusCompleted TransactionStatus = "completed"
	TransactionStatusFailed    TransactionStatus = "failed"
	TransactionStatusRefunded  TransactionStatus = "refunded"
	TransactionStatusCancelled TransactionStatus = "cancelled"
)

// IsValidStatus checks if a payment status is valid
func (ts TransactionStatus) IsValidStatus() bool {
	switch ts {
	case TransactionStatusPending, TransactionStatusCompleted,
		TransactionStatusFailed, TransactionStatusRefunded,
		TransactionStatusCancelled:
		return true
	}
	return false
}

// NewTransaction creates a new Transaction with default values
func NewTransaction(orderID string, paymentMethod string, grandTotal float64) *Transaction {
	now := time.Now()
	return &Transaction{
		OrderID:         orderID,
		PaymentMethod:   paymentMethod,
		GrandTotal:      grandTotal,
		PaymentStatus:   "pending",
		TransactionDate: now,
		CreatedAt:       now,
		UpdatedAt:       now,
	}
}

// CreateTransactionRequest represents the payload for creating a transaction
type CreateTransactionRequest struct {
	OrderID         string    `json:"order_id" validate:"required"`
	PaymentMethod   string    `json:"payment_method" validate:"required"`
	GrandTotal      float64   `json:"grand_total" validate:"required,gt=0"`
	PaymentStatus   string    `json:"payment_status,omitempty" validate:"omitempty,oneof=pending completed failed refunded cancelled"`
	TransactionDate time.Time `json:"transaction_date,omitempty"`
}

// UpdateTransactionStatusRequest represents the payload for updating transaction status
type UpdateTransactionStatusRequest struct {
	Status string `json:"status" validate:"required,oneof=pending completed failed refunded cancelled"`
}
