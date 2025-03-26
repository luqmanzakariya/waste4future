package repository

import (
	"context"
	"errors"
	"operation-service/model"
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type IOrderDetailRepository interface {
	Create(ctx context.Context, orderDetail model.OrderDetail) (model.OrderDetail, error)
	ReadAll(ctx context.Context) ([]model.OrderDetail, error)
	ReadByID(ctx context.Context, id string) (model.OrderDetail, error)
	Update(ctx context.Context, id string, orderDetail model.OrderDetail) (model.OrderDetail, error)
	Delete(ctx context.Context, id string) error
}

type orderDetailRepository struct {
	OrderDetailCollection *mongo.Collection
}

func NewOrderDetailRepository(db *mongo.Database) IOrderDetailRepository {
	return &orderDetailRepository{
		OrderDetailCollection: db.Collection("order_details"),
	}
}

func (r *orderDetailRepository) Create(ctx context.Context, orderDetail model.OrderDetail) (model.OrderDetail, error) {
	orderDetail.CreatedAt = time.Now()
	orderDetail.UpdatedAt = time.Now()

	res, err := r.OrderDetailCollection.InsertOne(ctx, orderDetail)
	if err != nil {
		return model.OrderDetail{}, err
	}

	insertedID, ok := res.InsertedID.(bson.ObjectID)
	if !ok {
		return model.OrderDetail{}, errors.New("failed to get inserted ID")
	}

	orderDetail.ID = insertedID
	return orderDetail, nil
}

func (r *orderDetailRepository) ReadAll(ctx context.Context) ([]model.OrderDetail, error) {
	var orderDetails []model.OrderDetail
	cursor, err := r.OrderDetailCollection.Find(ctx, bson.D{})
	if err != nil {
		return orderDetails, err
	}

	if err = cursor.All(ctx, &orderDetails); err != nil {
		return orderDetails, err
	}
	return orderDetails, nil
}

func (r *orderDetailRepository) ReadByID(ctx context.Context, id string) (model.OrderDetail, error) {
	objectID, err := bson.ObjectIDFromHex(id)
	if err != nil {
		return model.OrderDetail{}, errors.New("invalid ID format")
	}

	filter := bson.M{"_id": objectID}
	var orderDetail model.OrderDetail
	err = r.OrderDetailCollection.FindOne(ctx, filter).Decode(&orderDetail)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return model.OrderDetail{}, errors.New("order detail not found")
		}
		return model.OrderDetail{}, err
	}
	return orderDetail, nil
}

func (r *orderDetailRepository) Update(ctx context.Context, id string, orderDetail model.OrderDetail) (model.OrderDetail, error) {
	objectID, err := bson.ObjectIDFromHex(id)
	if err != nil {
		return orderDetail, errors.New("invalid ID format")
	}

	orderDetail.UpdatedAt = time.Now()
	filter := bson.M{"_id": objectID}
	update := bson.M{"$set": orderDetail}

	res, err := r.OrderDetailCollection.UpdateOne(ctx, filter, update)
	if err != nil {
		return orderDetail, err
	}

	if res.MatchedCount == 0 {
		return orderDetail, errors.New("order detail not found")
	}
	return orderDetail, nil
}

func (r *orderDetailRepository) Delete(ctx context.Context, id string) error {
	objectID, err := bson.ObjectIDFromHex(id)
	if err != nil {
		return errors.New("invalid ID format")
	}

	filter := bson.M{"_id": objectID}
	res, err := r.OrderDetailCollection.DeleteOne(ctx, filter)
	if err != nil {
		return err
	}

	if res.DeletedCount == 0 {
		return errors.New("order detail not found")
	}
	return nil
}
