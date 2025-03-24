package usecase

import (
	"context"
	"errors"
	"operation-service/model"
	"operation-service/repository"
	"time"

	"github.com/go-playground/validator/v10"
)

type IOrderUsecase interface {
	Create(ctx context.Context, payload model.PayloadCreateOrder, userId int64) (model.ResponseOrder, error)
	FindAll(ctx context.Context) ([]model.Order, error)
	FindByID(ctx context.Context, id string) (model.ResponseOrder, error)
	Update(ctx context.Context, id string, payload model.PayloadUpdateOrder) (model.ResponseOrder, error)
	Delete(ctx context.Context, id string) error
	SaveOrderDetail(ctx context.Context, orderId string, userId int64) error
	CheckoutOrder(ctx context.Context, userId int64) error
}

type orderUsecase struct {
	OrderRepo repository.IOrderRepository
	Validate  *validator.Validate
}

func NewOrderUsecase(OrderRepo repository.IOrderRepository, validate *validator.Validate) IOrderUsecase {
	return &orderUsecase{
		OrderRepo: OrderRepo,
		Validate:  validate,
	}
}

func (o *orderUsecase) Create(ctx context.Context, payload model.PayloadCreateOrder, userId int64) (model.ResponseOrder, error) {
	err := o.Validate.Struct(payload)
	if err != nil {
		return model.ResponseOrder{}, err
	}

	order := model.Order{
		UserID:          userId,
		DriverID:        "",
		OrderDetailIDs:  model.OrderDetailIDs{},
		OrderDate:       time.Now(),
		OrderStatus:     model.OrderStatusDraft,
		ShippingStatus:  model.ShippingStatusUnassigned,
		UpdatedShipping: time.Now(),
		Note:            payload.Note,
		CreatedAt:       time.Now(),
		UpdatedAt:       time.Now(),
	}

	res, err := o.OrderRepo.Create(ctx, order, userId)
	if err != nil {
		return model.ResponseOrder{}, err
	}

	response := model.ResponseOrder{
		ID:              res.ID.Hex(),
		DriverID:        res.DriverID,
		OrderDetailIDs:  res.OrderDetailIDs,
		OrderDate:       res.OrderDate,
		OrderStatus:     res.OrderStatus,
		ShippingStatus:  res.ShippingStatus,
		UpdatedShipping: res.UpdatedShipping,
		Note:            res.Note,
		CreatedAt:       res.CreatedAt,
		UpdatedAt:       res.UpdatedAt,
	}

	return response, nil
}

func (o *orderUsecase) FindAll(ctx context.Context) ([]model.Order, error) {
	var orders []model.Order
	orders, err := o.OrderRepo.ReadAll(ctx)
	if err != nil {
		return orders, err
	}

	return orders, nil
}

func (o *orderUsecase) FindByID(ctx context.Context, id string) (model.ResponseOrder, error) {
	var order model.Order
	order, err := o.OrderRepo.ReadByID(ctx, id)
	if err != nil {
		return model.ResponseOrder{}, errors.New("order not found")
	}

	response := model.ResponseOrder{
		ID:              order.ID.Hex(),
		DriverID:        order.DriverID,
		OrderDetailIDs:  order.OrderDetailIDs,
		OrderDate:       order.OrderDate,
		OrderStatus:     order.OrderStatus,
		ShippingStatus:  order.ShippingStatus,
		UpdatedShipping: order.UpdatedShipping,
		Note:            order.Note,
		CreatedAt:       order.CreatedAt,
		UpdatedAt:       order.UpdatedAt,
	}

	return response, nil
}

func (o *orderUsecase) Update(ctx context.Context, id string, payload model.PayloadUpdateOrder) (model.ResponseOrder, error) {
	err := o.Validate.Struct(payload)
	if err != nil {
		return model.ResponseOrder{}, err
	}

	order, err := o.OrderRepo.ReadByID(ctx, id)
	if err != nil {
		return model.ResponseOrder{}, errors.New("order not found")
	}

	payloadUpdated := model.Order{
		ID:             order.ID,
		OrderDetailIDs: payload.OrderDetailIDs,
		Note:           payload.Note,
		UpdatedAt:      time.Now(),
	}

	updatedOrder, err := o.OrderRepo.Update(ctx, id, payloadUpdated)
	if err != nil {
		return model.ResponseOrder{}, errors.New("error updating order")
	}

	response := model.ResponseOrder{
		ID:              updatedOrder.ID.Hex(),
		DriverID:        updatedOrder.DriverID,
		OrderDetailIDs:  updatedOrder.OrderDetailIDs,
		OrderDate:       updatedOrder.OrderDate,
		OrderStatus:     updatedOrder.OrderStatus,
		ShippingStatus:  updatedOrder.ShippingStatus,
		UpdatedShipping: updatedOrder.UpdatedShipping,
		Note:            updatedOrder.Note,
		CreatedAt:       updatedOrder.CreatedAt,
		UpdatedAt:       updatedOrder.UpdatedAt,
	}

	return response, nil
}

func (o *orderUsecase) Delete(ctx context.Context, id string) error {
	return o.OrderRepo.Delete(ctx, id)
}

func (o *orderUsecase) SaveOrderDetail(ctx context.Context, orderId string, userId int64) error {
	return o.OrderRepo.SaveOrderDetail(ctx, orderId, userId)
}

func (o *orderUsecase) CheckoutOrder(ctx context.Context, userId int64) error {
	return o.OrderRepo.CheckoutOrder(ctx, userId)
}
