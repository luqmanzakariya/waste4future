package repository

import (
	"address-service/model"
	"context"
	"errors"
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type IAddressRepository interface {
	Create(ctx context.Context, address model.Address) (model.Address, error)
	ReadAll(ctx context.Context) ([]model.Address, error)
	ReadByID(ctx context.Context, id string) (model.Address, error)
	Update(ctx context.Context, id string, address model.Address) (model.Address, error)
	Delete(ctx context.Context, id string) error
	FindByName(ctx context.Context, name string) (model.Address, error)
}

type addressRepository struct {
	AddressCollection *mongo.Collection
}

func NewAddressRepository(db *mongo.Database) IAddressRepository {
	return &addressRepository{
		AddressCollection: db.Collection("address"),
	}
}

func (a *addressRepository) Create(ctx context.Context, address model.Address) (model.Address, error) {
	// Set timestamps
	address.CreatedAt = time.Now()
	address.UpdatedAt = time.Now()

	res, err := a.AddressCollection.InsertOne(ctx, address)
	if err != nil {
		return model.Address{}, err
	}

	insertedID, ok := res.InsertedID.(bson.ObjectID)
	if !ok {
		return model.Address{}, errors.New("failed to get inserted ID")
	}

	address.ID = insertedID
	return address, nil
}

func (a *addressRepository) ReadAll(ctx context.Context) ([]model.Address, error) {
	var addresses []model.Address
	cursor, err := a.AddressCollection.Find(ctx, bson.D{})
	if err != nil {
		return addresses, err
	}

	if err = cursor.All(ctx, &addresses); err != nil {
		return addresses, err
	}

	return addresses, nil
}

func (a *addressRepository) ReadByID(ctx context.Context, id string) (model.Address, error) {
	// Convert the string ID to bson.ObjectID
	objectID, err := bson.ObjectIDFromHex(id)
	if err != nil {
		return model.Address{}, errors.New("invalid ID format")
	}

	// Define a filter to find the document by _id
	filter := bson.M{"_id": objectID}

	// Query the collection
	var address model.Address
	err = a.AddressCollection.FindOne(ctx, filter).Decode(&address)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return model.Address{}, errors.New("address not found")
		}
		return model.Address{}, err
	}

	return address, nil
}

func (a *addressRepository) Update(ctx context.Context, id string, address model.Address) (model.Address, error) {
	// Convert the string ID to bson.ObjectID
	objectID, err := bson.ObjectIDFromHex(id)
	if err != nil {
		return address, errors.New("invalid ID format")
	}

	// Update timestamp
	address.UpdatedAt = time.Now()

	// Define a filter to find the document by _id
	filter := bson.M{"_id": objectID}

	// Define an update document
	update := bson.M{"$set": address}

	// Perform the update operation
	res, err := a.AddressCollection.UpdateOne(ctx, filter, update)
	if err != nil {
		return address, err
	}

	// Check if any document was updated
	if res.MatchedCount == 0 {
		return address, errors.New("address not found")
	}

	return address, nil
}

func (a *addressRepository) Delete(ctx context.Context, id string) error {
	// Convert the string ID to bson.ObjectID
	objectID, err := bson.ObjectIDFromHex(id)
	if err != nil {
		return errors.New("invalid ID format")
	}

	// Define a filter to find the document by _id
	filter := bson.M{"_id": objectID}

	// Perform the delete operation
	res, err := a.AddressCollection.DeleteOne(ctx, filter)
	if err != nil {
		return err
	}

	// Check if any document was deleted
	if res.DeletedCount == 0 {
		return errors.New("address not found")
	}

	return nil
}

func (a *addressRepository) FindByName(ctx context.Context, name string) (model.Address, error) {
	// Define a filter to find the document by name
	filter := bson.M{"name": name}

	// Query the collection
	var address model.Address
	err := a.AddressCollection.FindOne(ctx, filter).Decode(&address)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return model.Address{}, errors.New("address not found")
		}
		return model.Address{}, err
	}

	return address, nil
}
