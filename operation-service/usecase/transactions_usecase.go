package usecase

import (
	"context"
	"time"

	"operation-service/model"
	"operation-service/repository"
	"operation-service/utils"

	"github.com/go-playground/validator/v10"
	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Transaction represents a payment transaction in the system
type Transaction struct {
	ID              primitive.ObjectID `bson:"_id,omitempty"`
	OrderID         primitive.ObjectID `bson:"order_id"`
	PaymentMethod   string             `bson:"payment_method"`
	GrandTotal      float64            `bson:"grand_total"`
	PaymentStatus   string             `bson:"payment_status"`
	TransactionDate time.Time          `bson:"transaction_date"`
}

// TransactionRepository defines the interface for transaction persistence operations
type TransactionRepository interface {
	Create(ctx context.Context, transaction *model.Transaction) error
	GetByID(ctx context.Context, id primitive.ObjectID) (*model.Transaction, error)
	GetByOrderID(ctx context.Context, orderID string) ([]*model.Transaction, error)
	UpdateStatus(ctx context.Context, id primitive.ObjectID, status string) error
	Delete(ctx context.Context, id primitive.ObjectID) error
	ListByDateRange(ctx context.Context, start, end time.Time) ([]*model.Transaction, error)
}

// TransactionUsecase implements the business logic for transaction operations
type TransactionUsecase struct {
	repo      TransactionRepository
	validator *validator.Validate
	orderRepo repository.IOrderRepository
}

// NewTransactionUsecase creates a new transaction usecase instance
func NewTransactionUsecase(repo TransactionRepository, validator *validator.Validate, orderRepo repository.IOrderRepository) *TransactionUsecase {
	return &TransactionUsecase{
		repo:      repo,
		validator: validator,
		orderRepo: orderRepo,
	}
}

// CreateTransaction handles the creation of a new transaction
func (uc *TransactionUsecase) CreateTransaction(ctx context.Context, req *model.CreateTransactionRequest, userID int64) (*model.Transaction, error) {
	// Validate request
	if err := uc.validator.Struct(req); err != nil {
		return nil, utils.ValidationError(err)
	}

	// Create new transaction model
	transaction := model.NewTransaction(
		req.OrderID,
		req.PaymentMethod,
		req.GrandTotal,
	)

	// Persist the transaction
	if err := uc.repo.Create(ctx, transaction); err != nil {
		return nil, errors.Wrap(err, "failed to create transaction")
	}

	err := uc.orderRepo.CheckoutOrder(ctx, userID)
	if err != nil {
		return nil, errors.Wrap(err, "failed to checkout order")
	}

	return transaction, nil
}

// GetTransaction retrieves a transaction by its ID
func (uc *TransactionUsecase) GetTransaction(ctx context.Context, id primitive.ObjectID) (*model.Transaction, error) {
	if id.IsZero() {
		return nil, errors.New("invalid transaction ID")
	}

	transaction, err := uc.repo.GetByID(ctx, id)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get transaction")
	}

	if transaction == nil {
		return nil, errors.New("transaction not found")
	}

	return transaction, nil
}

// GetTransactionsByOrder retrieves all transactions for a specific order
func (uc *TransactionUsecase) GetTransactionsByOrder(ctx context.Context, orderID string) ([]*model.Transaction, error) {
	if orderID == "" {
		return nil, errors.New("order ID is required")
	}

	transactions, err := uc.repo.GetByOrderID(ctx, orderID)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get transactions by order")
	}

	return transactions, nil
}

// UpdateTransactionStatus updates the payment status of a transaction
func (uc *TransactionUsecase) UpdateTransactionStatus(ctx context.Context, id primitive.ObjectID, status string) error {
	if id.IsZero() {
		return errors.New("invalid transaction ID")
	}

	if !model.TransactionStatus(status).IsValidStatus() {
		return errors.New("invalid payment status")
	}

	err := uc.repo.UpdateStatus(ctx, id, status)
	if err != nil {
		return errors.Wrap(err, "failed to update transaction status")
	}

	return nil
}

// DeleteTransaction removes a transaction from the system
func (uc *TransactionUsecase) DeleteTransaction(ctx context.Context, id primitive.ObjectID) error {
	if id.IsZero() {
		return errors.New("invalid transaction ID")
	}

	err := uc.repo.Delete(ctx, id)
	if err != nil {
		return errors.Wrap(err, "failed to delete transaction")
	}

	return nil
}

// GetTransactionsByDateRange retrieves transactions within a specific date range
func (uc *TransactionUsecase) GetTransactionsByDateRange(ctx context.Context, start, end time.Time) ([]*model.Transaction, error) {
	if start.IsZero() || end.IsZero() {
		return nil, errors.New("both start and end dates are required")
	}
	if end.Before(start) {
		return nil, errors.New("end date must be after start date")
	}

	transactions, err := uc.repo.ListByDateRange(ctx, start, end)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get transactions by date range")
	}

	return transactions, nil
}
