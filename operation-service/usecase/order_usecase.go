package usecase

import (
	"context"
	"errors"
	"fmt"
	"operation-service/model"
	"operation-service/repository"
	"time"

	"github.com/go-playground/validator/v10"
)

type IOrderUsecase interface {
	Create(ctx context.Context, payload model.PayloadCreateOrder, userId int64) (model.ResponseOrder, error)
	FindAll(ctx context.Context) ([]model.Order, error)
	FindAllByUserID(ctx context.Context, userId int64) ([]model.Order, error)
	FindByID(ctx context.Context, id string) (model.ResponseOrder, error)
	Update(ctx context.Context, id string, payload model.PayloadUpdateOrder) (model.ResponseOrder, error)
	Delete(ctx context.Context, id string) error
	SaveOrderDetail(ctx context.Context, orderId string, userId int64) error
	CheckoutOrder(ctx context.Context, userId int64) error
	SchedulerUpdateOrderAndShipping(ctx context.Context) error
}

type orderUsecase struct {
	OrderRepo       repository.IOrderRepository
	Validate        *validator.Validate
	OrderDetailRepo repository.IOrderDetailRepository
	TransactionRepo repository.ITransactionRepository
	DriverRepo      repository.IDriverRepository
}

func NewOrderUsecase(
	OrderRepo repository.IOrderRepository,
	validate *validator.Validate,
	transactionRepo repository.ITransactionRepository,
	orderDetailRepo repository.IOrderDetailRepository,
	driverRepo repository.IDriverRepository,
) IOrderUsecase {
	return &orderUsecase{
		OrderRepo:       OrderRepo,
		Validate:        validate,
		TransactionRepo: transactionRepo,
		OrderDetailRepo: orderDetailRepo,
		DriverRepo:      driverRepo,
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

func (o *orderUsecase) FindAllByUserID(ctx context.Context, userId int64) ([]model.Order, error) {
	var orders []model.Order
	orders, err := o.OrderRepo.ReadAllByUserID(ctx, userId)
	if err != nil {
		return orders, err
	}

	return orders, nil
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
	res, err := o.OrderRepo.CheckoutOrder(ctx, userId)
	if err != nil {
		return errors.New("error updating order")
	}

	var grandTotal float64 = 0
	// Loop through each order detail ID
	for _, orderDetailID := range res.OrderDetailIDs {
		// Call ReadByID for each order detail
		orderDetail, err := o.OrderDetailRepo.ReadByID(ctx, orderDetailID)
		if err != nil {
			// Handle error (you might want to continue or return)
			fmt.Printf("error reading order detail %s: %v\n", orderDetailID, err)
			continue
		}

		// Do something with the orderDetail
		grandTotal += orderDetail.SubTotal + orderDetail.DeliveryPrice
	}

	transaction := model.Transaction{
		OrderID:         res.ID.Hex(),
		PaymentMethod:   model.PaymentMethodBankTransfer,
		GrandTotal:      grandTotal,
		PaymentStatus:   model.PaymentStatusPending,
		TransactionDate: time.Now(),
		CreatedAt:       time.Now(),
		UpdatedAt:       time.Now(),
	}

	_, err = o.TransactionRepo.Create(ctx, transaction)
	if err != nil {
		return errors.New("error create transaction")
	}

	return nil
}

func (o *orderUsecase) SchedulerUpdateOrderAndShipping(ctx context.Context) error {
	// ##### 1 #####
	// # Find all delivery orders
	ordersDelivery, err := o.OrderRepo.FindAllWithShippingStatusDelivery(ctx)
	if err != nil {
		return fmt.Errorf("failed to find delivery orders: %w", err)
	}
	fmt.Println("Current delivery orders:", len(ordersDelivery))

	// # Update each pickup order to delivery status
	for _, order := range ordersDelivery {
		// Prepare the update payload
		updateData := model.Order{
			ID:              order.ID,
			UserID:          order.UserID,
			DriverID:        order.DriverID,
			OrderDetailIDs:  order.OrderDetailIDs,
			OrderDate:       order.OrderDate,
			OrderStatus:     order.OrderStatus,
			ShippingStatus:  model.ShippingStatusFinish,
			UpdatedShipping: time.Now(),
			Note:            order.Note,
			CreatedAt:       order.CreatedAt,
			UpdatedAt:       time.Now(),
		}

		// # Update the order
		_, err := o.OrderRepo.Update(ctx, order.ID.Hex(), updateData)
		if err != nil {
			// Log the error but continue with other orders
			fmt.Printf("failed to update order %s: %v\n", order.ID.Hex(), err)
			continue
		}

		// # Find Driver
		driver, err := o.DriverRepo.ReadByID(ctx, order.DriverID)
		if err != nil {
			return fmt.Errorf("failed find driver")
		}

		// # Update Driver Status
		updatedDriver := model.Driver{
			ID:           driver.ID,
			Name:         driver.Name,
			Phone:        driver.Phone,
			LicensePlate: driver.LicensePlate,
			Status:       model.DriverStatusActive,
			CreatedAt:    driver.CreatedAt,
			UpdatedAt:    time.Now(),
		}

		_, err = o.DriverRepo.Update(ctx, driver.ID.Hex(), updatedDriver)
		if err != nil {
			return fmt.Errorf("error update driver")
		}
	}

	// ##### 2 #####
	// # Find all pickup orders with drivers assigned
	ordersPickup, err := o.OrderRepo.FindAllWithShippingStatusPickupWithDriver(ctx)
	if err != nil {
		return fmt.Errorf("failed to find pickup orders: %w", err)
	}
	fmt.Println("Found pickup orders:", len(ordersPickup))

	// Update each pickup order to delivery status
	for _, order := range ordersPickup {
		// Prepare the update payload
		updateData := model.Order{
			ID:              order.ID,
			UserID:          order.UserID,
			DriverID:        order.DriverID,
			OrderDetailIDs:  order.OrderDetailIDs,
			OrderDate:       order.OrderDate,
			OrderStatus:     order.OrderStatus,
			ShippingStatus:  model.ShippingStatusDelivery,
			UpdatedShipping: time.Now(),
			Note:            order.Note,
			CreatedAt:       order.CreatedAt,
			UpdatedAt:       time.Now(),
		}

		// Update the order
		_, err := o.OrderRepo.Update(ctx, order.ID.Hex(), updateData)
		if err != nil {
			// Log the error but continue with other orders
			fmt.Printf("failed to update order %s: %v\n", order.ID.Hex(), err)
			continue
		}

		fmt.Printf("Updated order %s to delivery status\n", order.ID.Hex())
	}

	// ##### 3 #####
	// # Find all delivery orders
	ordersPickupWithoutDriver, err := o.OrderRepo.FindAllWithShippingStatusPickupWithoutDriver(ctx)
	if err != nil {
		return fmt.Errorf("failed to find delivery orders: %w", err)
	}
	fmt.Println("Current delivery orders:", len(ordersPickupWithoutDriver))

	// # Update each pickup order to delivery status
	for _, order := range ordersPickupWithoutDriver {
		drivers, err := o.DriverRepo.ReadAllActive(ctx)
		if err != nil {
			return errors.New("failed find drivers")
		}

		driverId := ""
		if len(drivers) > 0 {
			driverId = drivers[0].ID.Hex()

			foundDriver, err := o.DriverRepo.ReadByID(ctx, drivers[0].ID.Hex())
			if err != nil {
				return errors.New("failed found driver")
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

			_, err = o.DriverRepo.Update(ctx, drivers[0].ID.Hex(), assignedDriver)
			if err != nil {
				return errors.New("failed found driver")
			}
		}

		// Prepare the update payload
		updateData := model.Order{
			ID:              order.ID,
			UserID:          order.UserID,
			DriverID:        driverId,
			OrderDetailIDs:  order.OrderDetailIDs,
			OrderDate:       order.OrderDate,
			OrderStatus:     order.OrderStatus,
			ShippingStatus:  model.ShippingStatusPickup,
			UpdatedShipping: time.Now(),
			Note:            order.Note,
			CreatedAt:       order.CreatedAt,
			UpdatedAt:       time.Now(),
		}

		// # Update the order
		_, err = o.OrderRepo.Update(ctx, order.ID.Hex(), updateData)
		if err != nil {
			// Log the error but continue with other orders
			fmt.Printf("failed to update order %s: %v\n", order.ID.Hex(), err)
			continue
		}
	}

	return nil
}
