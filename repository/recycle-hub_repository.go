package repository

import (
	"context"
	"recyclehub-service/model/domain"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type RecycleHubRepository interface {
	Create(ctx context.Context, hub *domain.RecycleHub) error
	FindAll(ctx context.Context) ([]domain.RecycleHub, error)
	FindByID(ctx context.Context, id string) (*domain.RecycleHub, error)
	Update(ctx context.Context, id string, hub *domain.RecycleHub) error
	Delete(ctx context.Context, id string) error
}

type recycleHubRepository struct {
	collection *mongo.Collection
}

func NewRecycleHubRepository(collection *mongo.Collection) RecycleHubRepository {
	return &recycleHubRepository{collection}
}

func (r *recycleHubRepository) Create(ctx context.Context, hub *domain.RecycleHub) error {
	_, err := r.collection.InsertOne(ctx, hub)
	return err
}

func (r *recycleHubRepository) FindAll(ctx context.Context) ([]domain.RecycleHub, error) {
	var hubs []domain.RecycleHub
	cursor, err := r.collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	if err = cursor.All(ctx, &hubs); err != nil {
		return nil, err
	}
	return hubs, nil
}

func (r *recycleHubRepository) FindByID(ctx context.Context, id string) (*domain.RecycleHub, error) {
	var hub domain.RecycleHub
	err := r.collection.FindOne(ctx, bson.M{"_id": id}).Decode(&hub)
	if err != nil {
		return nil, err
	}
	return &hub, nil
}

func (r *recycleHubRepository) Update(ctx context.Context, id string, hub *domain.RecycleHub) error {
	_, err := r.collection.UpdateOne(ctx, bson.M{"_id": id}, bson.M{"$set": hub})
	return err
}

func (r *recycleHubRepository) Delete(ctx context.Context, id string) error {
	_, err := r.collection.DeleteOne(ctx, bson.M{"_id": id})
	return err
}
