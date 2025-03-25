package usecase

import (
	"context"
	"errors"
	"operation-service/model"
	"operation-service/repository"
	"time"

	"github.com/go-playground/validator/v10"
)

type ITransactionUsecase interface {
	Create(ctx context.Context, payload model.PayloadCreateTransaction) (model.ResponseTransaction, error)
	FindAll(ctx context.Context) ([]model.Transaction, error)
	FindByID(ctx context.Context, id string) (model.ResponseTransaction, error)
	Update(ctx context.Context, id string, payload model.PayloadUpdateTransaction) (model.ResponseTransaction, error)
	Delete(ctx context.Context, id string) error
}

type transactionUsecase struct {
	TransactionRepo repository.ITransactionRepository
	Validate        *validator.Validate
}

func NewTransactionUsecase(transactionRepo repository.ITransactionRepository, validate *validator.Validate) ITransactionUsecase {
	return &transactionUsecase{
		TransactionRepo: transactionRepo,
		Validate:        validate,
	}
}

func (t *transactionUsecase) Create(ctx context.Context, payload model.PayloadCreateTransaction) (model.ResponseTransaction, error) {
	err := t.Validate.Struct(payload)
	if err != nil {
		return model.ResponseTransaction{}, err
	}

	transaction := model.Transaction{
		OrderID:         payload.OrderID,
		PaymentMethod:   payload.PaymentMethod,
		GrandTotal:      payload.GrandTotal,
		PaymentStatus:   model.PaymentStatusPending,
		TransactionDate: time.Now(),
		CreatedAt:       time.Now(),
		UpdatedAt:       time.Now(),
	}

	res, err := t.TransactionRepo.Create(ctx, transaction)
	if err != nil {
		return model.ResponseTransaction{}, err
	}

	response := model.ResponseTransaction{
		ID:              res.ID.Hex(),
		OrderID:         res.OrderID,
		PaymentMethod:   res.PaymentMethod,
		GrandTotal:      res.GrandTotal,
		PaymentStatus:   res.PaymentStatus,
		TransactionDate: res.TransactionDate,
		CreatedAt:       res.CreatedAt,
		UpdatedAt:       res.UpdatedAt,
	}

	return response, nil
}

func (t *transactionUsecase) FindAll(ctx context.Context) ([]model.Transaction, error) {
	transactions, err := t.TransactionRepo.ReadAll(ctx)
	if err != nil {
		return nil, err
	}

	return transactions, nil
}

func (t *transactionUsecase) FindByID(ctx context.Context, id string) (model.ResponseTransaction, error) {
	transaction, err := t.TransactionRepo.ReadByID(ctx, id)
	if err != nil {
		return model.ResponseTransaction{}, errors.New("transaction not found")
	}

	response := model.ResponseTransaction{
		ID:              transaction.ID.Hex(),
		OrderID:         transaction.OrderID,
		PaymentMethod:   transaction.PaymentMethod,
		GrandTotal:      transaction.GrandTotal,
		PaymentStatus:   transaction.PaymentStatus,
		TransactionDate: transaction.TransactionDate,
		CreatedAt:       transaction.CreatedAt,
		UpdatedAt:       transaction.UpdatedAt,
	}

	return response, nil
}

func (t *transactionUsecase) Update(ctx context.Context, id string, payload model.PayloadUpdateTransaction) (model.ResponseTransaction, error) {
	err := t.Validate.Struct(payload)
	if err != nil {
		return model.ResponseTransaction{}, err
	}

	transaction, err := t.TransactionRepo.ReadByID(ctx, id)
	if err != nil {
		return model.ResponseTransaction{}, errors.New("transaction not found")
	}

	if payload.PaymentMethod != "" {
		transaction.PaymentMethod = payload.PaymentMethod
	}
	if payload.GrandTotal != 0 {
		transaction.GrandTotal = payload.GrandTotal
	}
	if payload.PaymentStatus != "" {
		transaction.PaymentStatus = payload.PaymentStatus
	}
	transaction.UpdatedAt = time.Now()

	updatedTransaction, err := t.TransactionRepo.Update(ctx, id, transaction)
	if err != nil {
		return model.ResponseTransaction{}, errors.New("error updating transaction")
	}

	response := model.ResponseTransaction{
		ID:              updatedTransaction.ID.Hex(),
		OrderID:         updatedTransaction.OrderID,
		PaymentMethod:   updatedTransaction.PaymentMethod,
		GrandTotal:      updatedTransaction.GrandTotal,
		PaymentStatus:   updatedTransaction.PaymentStatus,
		TransactionDate: updatedTransaction.TransactionDate,
		CreatedAt:       updatedTransaction.CreatedAt,
		UpdatedAt:       updatedTransaction.UpdatedAt,
	}

	return response, nil
}

func (t *transactionUsecase) Delete(ctx context.Context, id string) error {
	return t.TransactionRepo.Delete(ctx, id)
}
