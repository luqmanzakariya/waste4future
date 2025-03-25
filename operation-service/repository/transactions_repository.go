package repository

import (
	"context"
	"errors"
	"operation-service/model"
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type ITransactionRepository interface {
	Create(ctx context.Context, transaction model.Transaction) (model.Transaction, error)
	ReadAll(ctx context.Context) ([]model.Transaction, error)
	ReadByID(ctx context.Context, id string) (model.Transaction, error)
	Update(ctx context.Context, id string, transaction model.Transaction) (model.Transaction, error)
	Delete(ctx context.Context, id string) error
}

type transactionRepository struct {
	TransactionCollection *mongo.Collection
}

func NewTransactionRepository(db *mongo.Database) ITransactionRepository {
	return &transactionRepository{
		TransactionCollection: db.Collection("transactions"),
	}
}

func (t *transactionRepository) Create(ctx context.Context, transaction model.Transaction) (model.Transaction, error) {
	res, err := t.TransactionCollection.InsertOne(ctx, transaction)
	if err != nil {
		return model.Transaction{}, err
	}

	insertedID, ok := res.InsertedID.(bson.ObjectID)
	if !ok {
		return model.Transaction{}, errors.New("failed to get inserted ID")
	}

	transaction.ID = insertedID
	return transaction, nil
}

func (t *transactionRepository) ReadAll(ctx context.Context) ([]model.Transaction, error) {
	var transactions []model.Transaction
	cursor, err := t.TransactionCollection.Find(ctx, bson.D{})
	if err != nil {
		return transactions, err
	}

	if err = cursor.All(ctx, &transactions); err != nil {
		return transactions, err
	}

	return transactions, nil
}

func (t *transactionRepository) ReadByID(ctx context.Context, id string) (model.Transaction, error) {
	objectID, err := bson.ObjectIDFromHex(id)
	if err != nil {
		return model.Transaction{}, errors.New("invalid ID format")
	}

	filter := bson.M{"_id": objectID}
	var transaction model.Transaction
	err = t.TransactionCollection.FindOne(ctx, filter).Decode(&transaction)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return model.Transaction{}, errors.New("transaction not found")
		}
		return model.Transaction{}, err
	}

	return transaction, nil
}

func (t *transactionRepository) Update(ctx context.Context, id string, transaction model.Transaction) (model.Transaction, error) {
	objectID, err := bson.ObjectIDFromHex(id)
	if err != nil {
		return transaction, errors.New("invalid ID format")
	}

	transaction.UpdatedAt = time.Now()
	filter := bson.M{"_id": objectID}
	update := bson.M{"$set": transaction}

	res, err := t.TransactionCollection.UpdateOne(ctx, filter, update)
	if err != nil {
		return transaction, err
	}

	if res.MatchedCount == 0 {
		return transaction, errors.New("transaction not found")
	}

	return transaction, nil
}

func (t *transactionRepository) Delete(ctx context.Context, id string) error {
	objectID, err := bson.ObjectIDFromHex(id)
	if err != nil {
		return errors.New("invalid ID format")
	}

	filter := bson.M{"_id": objectID}
	res, err := t.TransactionCollection.DeleteOne(ctx, filter)
	if err != nil {
		return err
	}

	if res.DeletedCount == 0 {
		return errors.New("transaction not found")
	}

	return nil
}
