package repository

import (
	"context"
	"recyclehub-service/model/domain"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type WasteTypeRepository interface {
	FindAll(ctx context.Context) ([]domain.WasteType, error)
	FindById(ctx context.Context, id string) (domain.WasteType, error)
}

type wasteTypeRepository struct {
	collection *mongo.Collection
}

func NewWasteTypeRepository(collection *mongo.Collection) WasteTypeRepository {
	return &wasteTypeRepository{collection}
}

func (r *wasteTypeRepository) FindAll(ctx context.Context) ([]domain.WasteType, error) {
	var wasteTypes []domain.WasteType
	cursor, err := r.collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	if err = cursor.All(ctx, &wasteTypes); err != nil {
		return nil, err
	}
	return wasteTypes, nil
}

func (r *wasteTypeRepository) FindById(ctx context.Context, id string) (domain.WasteType, error) {
	var wasteType domain.WasteType
	objectID, err := primitive.ObjectIDFromHex(id) // Convert string ID to ObjectID
	if err != nil {
		return domain.WasteType{}, err
	}
	err = r.collection.FindOne(ctx, bson.M{"_id": objectID}).Decode(&wasteType)
	if err != nil {
		return domain.WasteType{}, err
	}
	return wasteType, nil
}
