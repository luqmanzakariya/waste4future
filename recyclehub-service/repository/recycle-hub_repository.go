package repository

import (
	"context"
	"errors"
	"reyclehub-service/model"
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type IRecycleHubRepository interface {
	Create(ctx context.Context, recycleHub model.RecycleHub) (model.RecycleHub, error)
	ReadAll(ctx context.Context) ([]model.RecycleHub, error)
	ReadByID(ctx context.Context, id string) (model.RecycleHub, error)
	Update(ctx context.Context, id string, recycleHub model.RecycleHub) (model.RecycleHub, error)
	Delete(ctx context.Context, id string) error
}

type recycleHubRepository struct {
	RecycleHubCollection *mongo.Collection
}

func NewRecycleHubRepository(db *mongo.Database) IRecycleHubRepository {
	return &recycleHubRepository{
		RecycleHubCollection: db.Collection("recycles"),
	}
}

func (r *recycleHubRepository) Create(ctx context.Context, recycleHub model.RecycleHub) (model.RecycleHub, error) {
	recycleHub.CreatedAt = time.Now()
	recycleHub.UpdatedAt = time.Now()

	res, err := r.RecycleHubCollection.InsertOne(ctx, recycleHub)
	if err != nil {
		return model.RecycleHub{}, err
	}

	insertedID, ok := res.InsertedID.(bson.ObjectID)
	if !ok {
		return model.RecycleHub{}, errors.New("failed to get inserted ID")
	}

	recycleHub.ID = insertedID
	return recycleHub, nil
}

func (r *recycleHubRepository) ReadAll(ctx context.Context) ([]model.RecycleHub, error) {
	var recycleHubs []model.RecycleHub
	cursor, err := r.RecycleHubCollection.Find(ctx, bson.D{})
	if err != nil {
		return recycleHubs, err
	}

	if err = cursor.All(ctx, &recycleHubs); err != nil {
		return recycleHubs, err
	}

	return recycleHubs, nil
}

func (r *recycleHubRepository) ReadByID(ctx context.Context, id string) (model.RecycleHub, error) {
	objectID, err := bson.ObjectIDFromHex(id)
	if err != nil {
		return model.RecycleHub{}, errors.New("invalid ID format")
	}

	filter := bson.M{"_id": objectID}
	var recycleHub model.RecycleHub
	err = r.RecycleHubCollection.FindOne(ctx, filter).Decode(&recycleHub)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return model.RecycleHub{}, errors.New("recycle hub not found")
		}
		return model.RecycleHub{}, err
	}

	return recycleHub, nil
}

func (r *recycleHubRepository) Update(ctx context.Context, id string, recycleHub model.RecycleHub) (model.RecycleHub, error) {
	objectID, err := bson.ObjectIDFromHex(id)
	if err != nil {
		return recycleHub, errors.New("invalid ID format")
	}

	recycleHub.UpdatedAt = time.Now()
	filter := bson.M{"_id": objectID}
	update := bson.M{"$set": recycleHub}

	res, err := r.RecycleHubCollection.UpdateOne(ctx, filter, update)
	if err != nil {
		return recycleHub, err
	}

	if res.MatchedCount == 0 {
		return recycleHub, errors.New("recycle hub not found")
	}

	return recycleHub, nil
}

func (r *recycleHubRepository) Delete(ctx context.Context, id string) error {
	objectID, err := bson.ObjectIDFromHex(id)
	if err != nil {
		return errors.New("invalid ID format")
	}

	filter := bson.M{"_id": objectID}
	res, err := r.RecycleHubCollection.DeleteOne(ctx, filter)
	if err != nil {
		return err
	}

	if res.DeletedCount == 0 {
		return errors.New("recycle hub not found")
	}

	return nil
}
