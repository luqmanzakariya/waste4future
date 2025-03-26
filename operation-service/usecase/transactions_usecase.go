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
	ApprovePayment(ctx context.Context, id string) (model.ResponseTransaction, error)
	RejectPayment(ctx context.Context, id string) (model.ResponseTransaction, error)
}

type transactionUsecase struct {
	TransactionRepo repository.ITransactionRepository
	Validate        *validator.Validate
	OrderRepo       repository.IOrderRepository
	DriverRepo      repository.IDriverRepository
}

func NewTransactionUsecase(transactionRepo repository.ITransactionRepository, validate *validator.Validate, orderRepo repository.IOrderRepository, driverRepo repository.IDriverRepository) ITransactionUsecase {
	return &transactionUsecase{
		TransactionRepo: transactionRepo,
		Validate:        validate,
		OrderRepo:       orderRepo,
		DriverRepo:      driverRepo,
	}
}

func (t *transactionUsecase) ApprovePayment(ctx context.Context, id string) (model.ResponseTransaction, error) {
	// # Find Driver
	drivers, err := t.DriverRepo.ReadAllActive(ctx)
	if err != nil {
		return model.ResponseTransaction{}, errors.New("failed find drivers")
	}

	transaction, err := t.TransactionRepo.ReadByID(ctx, id)
	if err != nil {
		return model.ResponseTransaction{}, errors.New("transaction not found")
	}

	// # Update Transaction
	payloadUpdate := model.Transaction{
		ID:              transaction.ID,
		OrderID:         transaction.OrderID,
		PaymentMethod:   transaction.PaymentMethod,
		GrandTotal:      transaction.GrandTotal,
		PaymentStatus:   model.PaymentStatusCompleted,
		TransactionDate: transaction.TransactionDate,
		CreatedAt:       transaction.CreatedAt,
		UpdatedAt:       transaction.UpdatedAt,
	}

	// # Update Transaction
	_, err = t.TransactionRepo.Update(ctx, id, payloadUpdate)
	if err != nil {
		return model.ResponseTransaction{}, errors.New("failed update transaction")
	}

	driverId := ""
	if len(drivers) > 0 {
		driverId = drivers[0].ID.Hex()

		foundDriver, err := t.DriverRepo.ReadByID(ctx, drivers[0].ID.Hex())
		if err != nil {
			return model.ResponseTransaction{}, errors.New("failed found driver")
		}

		// # Find Assigned Driver Detail
		assignedDriver := model.Driver{
			ID:           foundDriver.ID,
			Name:         foundDriver.Name,
			Phone:        foundDriver.Phone,
			LicensePlate: foundDriver.LicensePlate,
			Status:       model.DriverStatusWorking,
			CreatedAt:    foundDriver.CreatedAt,
			UpdatedAt:    time.Now(),
		}

		_, err = t.DriverRepo.Update(ctx, drivers[0].ID.Hex(), assignedDriver)
		if err != nil {
			return model.ResponseTransaction{}, errors.New("failed found driver")
		}
	}

	// # Find Order
	order, err := t.OrderRepo.ReadByID(ctx, transaction.OrderID)
	if err != nil {
		return model.ResponseTransaction{}, errors.New("failed found order")
	}

	// # Payload Update Order
	payloadUpdateOrder := model.Order{
		ID:              order.ID,
		UserID:          order.UserID,
		DriverID:        driverId,
		OrderDetailIDs:  order.OrderDetailIDs,
		OrderDate:       order.OrderDate,
		OrderStatus:     model.OrderStatusPaid,
		ShippingStatus:  model.ShippingStatusPickup,
		UpdatedShipping: time.Now(),
		Note:            order.Note,
		CreatedAt:       order.CreatedAt,
		UpdatedAt:       time.Now(),
	}

	// # Update Order
	_, err = t.OrderRepo.Update(ctx, order.ID.Hex(), payloadUpdateOrder)
	if err != nil {
		return model.ResponseTransaction{}, errors.New("failed update order")
	}

	response := model.ResponseTransaction{
		ID:              payloadUpdate.ID.Hex(),
		OrderID:         payloadUpdate.OrderID,
		PaymentMethod:   payloadUpdate.PaymentMethod,
		GrandTotal:      payloadUpdate.GrandTotal,
		PaymentStatus:   payloadUpdate.PaymentStatus,
		TransactionDate: payloadUpdate.TransactionDate,
		CreatedAt:       payloadUpdate.CreatedAt,
		UpdatedAt:       payloadUpdate.UpdatedAt,
	}

	return response, nil
}

func (t *transactionUsecase) RejectPayment(ctx context.Context, id string) (model.ResponseTransaction, error) {
	transaction, err := t.TransactionRepo.ReadByID(ctx, id)
	if err != nil {
		return model.ResponseTransaction{}, errors.New("transaction not found")
	}

	payloadUpdate := model.Transaction{
		ID:              transaction.ID,
		OrderID:         transaction.OrderID,
		PaymentMethod:   transaction.PaymentMethod,
		GrandTotal:      transaction.GrandTotal,
		PaymentStatus:   model.PaymentStatusRejected,
		TransactionDate: transaction.TransactionDate,
		CreatedAt:       transaction.CreatedAt,
		UpdatedAt:       transaction.UpdatedAt,
	}

	_, err = t.TransactionRepo.Update(ctx, id, payloadUpdate)
	if err != nil {
		return model.ResponseTransaction{}, errors.New("failed update transaction")
	}

	order, err := t.OrderRepo.ReadByID(ctx, transaction.OrderID)
	if err != nil {
		return model.ResponseTransaction{}, errors.New("failed found order")
	}

	payloadUpdateOrder := model.Order{
		ID:              order.ID,
		UserID:          order.UserID,
		DriverID:        order.DriverID,
		OrderDetailIDs:  order.OrderDetailIDs,
		OrderDate:       order.OrderDate,
		OrderStatus:     model.OrderStatusRejected,
		ShippingStatus:  order.ShippingStatus,
		UpdatedShipping: time.Now(),
		Note:            order.Note,
		CreatedAt:       order.CreatedAt,
		UpdatedAt:       order.UpdatedAt,
	}

	_, err = t.OrderRepo.Update(ctx, order.ID.Hex(), payloadUpdateOrder)
	if err != nil {
		return model.ResponseTransaction{}, errors.New("failed update order")
	}

	response := model.ResponseTransaction{
		ID:              payloadUpdate.ID.Hex(),
		OrderID:         payloadUpdate.OrderID,
		PaymentMethod:   payloadUpdate.PaymentMethod,
		GrandTotal:      payloadUpdate.GrandTotal,
		PaymentStatus:   payloadUpdate.PaymentStatus,
		TransactionDate: payloadUpdate.TransactionDate,
		CreatedAt:       payloadUpdate.CreatedAt,
		UpdatedAt:       payloadUpdate.UpdatedAt,
	}

	return response, nil
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
