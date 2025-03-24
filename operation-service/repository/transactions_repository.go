package repository

import (
	"context"
	"errors"
	"time"

	"operation-service/model"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

type TransactionRepository struct {
	collection *mongo.Collection
}

func NewTransactionRepository(db *mongo.Database, collectionName string) *TransactionRepository {
	return &TransactionRepository{
		collection: db.Collection(collectionName),
	}
}

func (r *TransactionRepository) Create(ctx context.Context, transaction *model.Transaction) error {
	if transaction == nil {
		return errors.New("transaction is nil")
	}

	if transaction.OrderID == "" {
		return errors.New("order ID is required")
	}

	_, err := r.collection.InsertOne(ctx, transaction)
	return err
}

func (r *TransactionRepository) GetByID(ctx context.Context, id primitive.ObjectID) (*model.Transaction, error) {
	if id.IsZero() {
		return nil, errors.New("invalid transaction ID")
	}

	var transaction model.Transaction
	err := r.collection.FindOne(ctx, bson.M{"_id": id}).Decode(&transaction)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, nil
		}
		return nil, err
	}

	return &transaction, nil
}

func (r *TransactionRepository) GetByOrderID(ctx context.Context, orderID string) ([]*model.Transaction, error) {
	if orderID == "" {
		return nil, errors.New("order ID is required")
	}

	cursor, err := r.collection.Find(ctx, bson.M{"order_id": orderID})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var transactions []*model.Transaction
	if err := cursor.All(ctx, &transactions); err != nil {
		return nil, err
	}

	return transactions, nil
}

func (r *TransactionRepository) UpdateStatus(ctx context.Context, id primitive.ObjectID, status string) error {
	if id.IsZero() {
		return errors.New("invalid transaction ID")
	}

	if !model.TransactionStatus(status).IsValidStatus() {
		return errors.New("invalid payment status")
	}

	update := bson.M{
		"$set": bson.M{
			"payment_status": status,
			"updated_at":     time.Now(),
		},
	}

	_, err := r.collection.UpdateByID(ctx, id, update)
	return err
}

func (r *TransactionRepository) Delete(ctx context.Context, id primitive.ObjectID) error {
	if id.IsZero() {
		return errors.New("invalid transaction ID")
	}

	_, err := r.collection.DeleteOne(ctx, bson.M{"_id": id})
	return err
}

func (r *TransactionRepository) ListByDateRange(ctx context.Context, start, end time.Time) ([]*model.Transaction, error) {
	if start.IsZero() || end.IsZero() {
		return nil, errors.New("both start and end dates are required")
	}
	if end.Before(start) {
		return nil, errors.New("end date must be after start date")
	}

	filter := bson.M{
		"transaction_date": bson.M{
			"$gte": start,
			"$lte": end,
		},
	}

	opts := options.Find().SetSort(bson.D{{Key: "transaction_date", Value: 1}})

	cursor, err := r.collection.Find(ctx, filter, opts)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var transactions []*model.Transaction
	if err := cursor.All(ctx, &transactions); err != nil {
		return nil, err
	}

	return transactions, nil
}
