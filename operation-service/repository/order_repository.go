package repository

import (
	"context"
	"errors"
	"operation-service/model"
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type IOrderRepository interface {
	Create(ctx context.Context, order model.Order, userId int64) (model.Order, error)
	ReadAll(ctx context.Context) ([]model.Order, error)
	ReadAllByUserID(ctx context.Context, userId int64) ([]model.Order, error)
	ReadByID(ctx context.Context, id string) (model.Order, error)
	Update(ctx context.Context, id string, address model.Order) (model.Order, error)
	Delete(ctx context.Context, id string) error
	SaveOrderDetail(ctx context.Context, orderId string, userId int64) error
	CheckoutOrder(ctx context.Context, userId int64) (model.Order, error)
	DeleteOrderDetailID(ctx context.Context, orderDetailId string, userId int64) error
	FindAllWithShippingStatusPickupWithDriver(ctx context.Context) ([]model.Order, error)
	FindAllWithShippingStatusPickupWithoutDriver(ctx context.Context) ([]model.Order, error)
	FindAllWithShippingStatusDelivery(ctx context.Context) ([]model.Order, error)
}

type orderRepository struct {
	OrderCollection *mongo.Collection
}

func NewOrderRepository(db *mongo.Database) IOrderRepository {
	return &orderRepository{
		OrderCollection: db.Collection("orders"),
	}
}

func (o *orderRepository) Create(ctx context.Context, order model.Order, userId int64) (model.Order, error) {
	res, err := o.OrderCollection.InsertOne(ctx, order)
	if err != nil {
		return model.Order{}, err
	}

	insertedID, ok := res.InsertedID.(bson.ObjectID)
	if !ok {
		return model.Order{}, errors.New("failed to get inserted ID")
	}

	order.ID = insertedID
	return order, nil
}

func (o *orderRepository) ReadAllByUserID(ctx context.Context, userId int64) ([]model.Order, error) {
	// Initialize as empty slice
	orders := make([]model.Order, 0)

	// Create filter for shipping_status = "pickup"
	filter := bson.M{
		"user_id": userId, // Match the user ID
	}

	// Find all matching orders
	cursor, err := o.OrderCollection.Find(ctx, filter)
	if err != nil {
		return orders, err
	}

	// Decode results into orders slice
	if err = cursor.All(ctx, &orders); err != nil {
		return orders, err
	}

	return orders, nil
}

func (o *orderRepository) ReadAll(ctx context.Context) ([]model.Order, error) {
	var orders []model.Order
	cursor, err := o.OrderCollection.Find(ctx, bson.D{})
	if err != nil {
		return orders, err
	}

	if err = cursor.All(ctx, &orders); err != nil {
		return orders, err
	}

	return orders, nil
}

func (o *orderRepository) ReadByID(ctx context.Context, id string) (model.Order, error) {
	// Convert the string ID to bson.ObjectID
	objectID, err := bson.ObjectIDFromHex(id)
	if err != nil {
		return model.Order{}, errors.New("invalid ID format")
	}

	// Define a filter to find the document by _id
	filter := bson.M{"_id": objectID}

	// Query the collection
	var order model.Order
	err = o.OrderCollection.FindOne(ctx, filter).Decode(&order)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return model.Order{}, errors.New("order not found")
		}
		return model.Order{}, err
	}

	return order, nil
}

func (o *orderRepository) Update(ctx context.Context, id string, order model.Order) (model.Order, error) {
	// Convert the string ID to bson.ObjectID
	objectID, err := bson.ObjectIDFromHex(id)
	if err != nil {
		return order, errors.New("invalid ID format")
	}

	// Update timestamp
	order.UpdatedAt = time.Now()

	// Define a filter to find the document by _id
	filter := bson.M{"_id": objectID}

	// Define an update document
	update := bson.M{"$set": order}

	// Perform the update operation
	res, err := o.OrderCollection.UpdateOne(ctx, filter, update)
	if err != nil {
		return order, err
	}

	// Check if any document was updated
	if res.MatchedCount == 0 {
		return order, errors.New("order not found")
	}

	return order, nil
}

func (o *orderRepository) Delete(ctx context.Context, id string) error {
	// Convert the string ID to bson.ObjectID
	objectID, err := bson.ObjectIDFromHex(id)
	if err != nil {
		return errors.New("invalid ID format")
	}

	// Define a filter to find the document by _id
	filter := bson.M{"_id": objectID}

	// Perform the delete operation
	res, err := o.OrderCollection.DeleteOne(ctx, filter)
	if err != nil {
		return err
	}

	// Check if any document was deleted
	if res.DeletedCount == 0 {
		return errors.New("order not found")
	}

	return nil
}

func (o *orderRepository) SaveOrderDetail(ctx context.Context, orderDetailId string, userId int64) error {
	// Convert the string orderDetailId to bson.ObjectID
	// objectID, err := bson.ObjectIDFromHex(orderDetailId)
	// if err != nil {
	// 	return errors.New("invalid order ID format")
	// }

	// Define a filter to find the order by _id and OrderStatus = "draft"
	filter := bson.M{
		"user_id":      userId,                 // Match the user ID
		"order_status": model.OrderStatusDraft, // Match the order status
	}

	// Query the collection to find the order
	foundStatusDraftOrder := true
	var foundOrder model.Order
	err := o.OrderCollection.FindOne(ctx, filter).Decode(&foundOrder)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			foundStatusDraftOrder = false
		} else {
			return err
		}
	}

	if foundStatusDraftOrder {
		// # Insert Order ID
		update := bson.M{
			"$push": bson.M{"order_detail_ids": orderDetailId}, // Append to the array
			"$set":  bson.M{"updated_at": time.Now()},          // Update the updated_at timestamp
		}

		// Perform the update operation
		_, err := o.OrderCollection.UpdateOne(ctx, filter, update)
		if err != nil {
			return err
		}

	} else {
		// # Create order and insert id
		newOrder := model.Order{
			UserID:          userId,
			DriverID:        "",
			OrderDetailIDs:  model.OrderDetailIDs{orderDetailId},
			OrderDate:       time.Now(),
			OrderStatus:     model.OrderStatusDraft,
			ShippingStatus:  model.ShippingStatusUnassigned,
			UpdatedShipping: time.Now(),
			Note:            "",
			CreatedAt:       time.Now(),
			UpdatedAt:       time.Now(),
		}

		res, err := o.OrderCollection.InsertOne(ctx, newOrder)
		if err != nil {
			return err
		}

		insertedID, ok := res.InsertedID.(bson.ObjectID)
		if !ok {
			return errors.New("failed to get inserted ID")
		}

		newOrder.ID = insertedID
		return nil
	}

	return nil
}

func (o *orderRepository) CheckoutOrder(ctx context.Context, userId int64) (model.Order, error) {
	// Define a filter to find the order by _id and OrderStatus = "draft"
	filter := bson.M{
		"user_id":      userId,                 // Match the user ID
		"order_status": model.OrderStatusDraft, // Match the order status
	}

	// Query the collection to find the order
	var order model.Order
	err := o.OrderCollection.FindOne(ctx, filter).Decode(&order)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return order, errors.New("order not found or not in draft status")
		}
		return order, err
	}

	// Define an update to set the order_status to "pending" and update the updated_at timestamp
	update := bson.M{
		"$set": bson.M{
			"order_status": model.OrderStatusPending, // Update order status to "pending"
			"updated_at":   time.Now(),               // Update the updated_at timestamp
		},
	}

	// Perform the update operation
	_, err = o.OrderCollection.UpdateOne(ctx, filter, update)
	if err != nil {
		return order, err
	}

	return order, nil
}

func (o *orderRepository) DeleteOrderDetailID(ctx context.Context, orderDetailId string, userId int64) error {
	// Define a filter to find the order by user_id and OrderStatus = "draft"
	filter := bson.M{
		"user_id":      userId,                 // Match the user ID
		"order_status": model.OrderStatusDraft, // Match the order status
	}

	// Define an update to remove the orderDetailId from the array and update the timestamp
	update := bson.M{
		"$pull": bson.M{"order_detail_ids": orderDetailId}, // Remove from the array
		"$set":  bson.M{"updated_at": time.Now()},          // Update the updated_at timestamp
	}

	// Perform the update operation
	result, err := o.OrderCollection.UpdateOne(ctx, filter, update)
	if err != nil {
		return err
	}

	// Check if any document was modified
	if result.MatchedCount == 0 {
		return errors.New("no draft order found for this user")
	}

	return nil
}

func (o *orderRepository) FindAllWithShippingStatusPickupWithDriver(ctx context.Context) ([]model.Order, error) {
	// Initialize as empty slice
	orders := make([]model.Order, 0)

	// Create filter for shipping_status = "pickup"
	filter := bson.M{
		"shipping_status": model.ShippingStatusPickup,
		"driver_id": bson.M{
			"$ne": "", // Not equal to empty string
		},
	}

	// Find all matching orders
	cursor, err := o.OrderCollection.Find(ctx, filter)
	if err != nil {
		return orders, err
	}

	// Decode results into orders slice
	if err = cursor.All(ctx, &orders); err != nil {
		return orders, err
	}

	return orders, nil
}

func (o *orderRepository) FindAllWithShippingStatusPickupWithoutDriver(ctx context.Context) ([]model.Order, error) {
	// Initialize as empty slice
	orders := make([]model.Order, 0)

	// Create filter for shipping_status = "pickup"
	filter := bson.M{
		"shipping_status": model.ShippingStatusPickup,
		"driver_id":       "",
	}

	// Find all matching orders
	cursor, err := o.OrderCollection.Find(ctx, filter)
	if err != nil {
		return orders, err
	}

	// Decode results into orders slice
	if err = cursor.All(ctx, &orders); err != nil {
		return orders, err
	}

	return orders, nil
}

func (o *orderRepository) FindAllWithShippingStatusDelivery(ctx context.Context) ([]model.Order, error) {
	// Initialize as empty slice
	orders := make([]model.Order, 0)

	// Create filter for shipping_status = "pickup"
	filter := bson.M{
		"shipping_status": model.ShippingStatusDelivery,
	}

	// Find all matching orders
	cursor, err := o.OrderCollection.Find(ctx, filter)
	if err != nil {
		return orders, err
	}

	// Decode results into orders slice
	if err = cursor.All(ctx, &orders); err != nil {
		return orders, err
	}

	return orders, nil
}
