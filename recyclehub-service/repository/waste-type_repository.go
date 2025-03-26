package repository

import (
	"context"
	"errors"
	"reyclehub-service/model"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type IWasteTypeRepository interface {
	ReadAll(ctx context.Context) ([]model.WasteType, error)
	ReadByID(ctx context.Context, id string) (model.WasteType, error)
}

type wasteTypeRepository struct {
	WasteTypeCollection *mongo.Collection
}

func NewWasteTypeRepository(db *mongo.Database) IWasteTypeRepository {
	return &wasteTypeRepository{
		WasteTypeCollection: db.Collection("wastes"),
	}
}

func (w *wasteTypeRepository) ReadAll(ctx context.Context) ([]model.WasteType, error) {
	var wasteTypes []model.WasteType
	cursor, err := w.WasteTypeCollection.Find(ctx, bson.D{})
	if err != nil {
		return wasteTypes, err
	}

	if err = cursor.All(ctx, &wasteTypes); err != nil {
		return wasteTypes, err
	}

	return wasteTypes, nil
}

func (w *wasteTypeRepository) ReadByID(ctx context.Context, id string) (model.WasteType, error) {
	objectID, err := bson.ObjectIDFromHex(id)
	if err != nil {
		return model.WasteType{}, errors.New("invalid ID format")
	}

	filter := bson.M{"_id": objectID}
	var wasteType model.WasteType
	err = w.WasteTypeCollection.FindOne(ctx, filter).Decode(&wasteType)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return model.WasteType{}, errors.New("waste type not found")
		}
		return model.WasteType{}, err
	}

	return wasteType, nil
}
