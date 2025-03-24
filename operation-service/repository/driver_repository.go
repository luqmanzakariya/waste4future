package repository

import (
	"context"
	"errors"
	"operation-service/model"
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type IDriverRepository interface {
	Create(ctx context.Context, driver model.Driver) (model.Driver, error)
	ReadAll(ctx context.Context) ([]model.Driver, error)
	ReadByID(ctx context.Context, id string) (model.Driver, error)
	Update(ctx context.Context, id string, driver model.Driver) (model.Driver, error)
	Delete(ctx context.Context, id string) error
}

type driverRepository struct {
	DriverCollection *mongo.Collection
}

func NewDriverRepository(db *mongo.Database) IDriverRepository {
	return &driverRepository{
		DriverCollection: db.Collection("drivers"),
	}
}

func (d *driverRepository) Create(ctx context.Context, driver model.Driver) (model.Driver, error) {
	res, err := d.DriverCollection.InsertOne(ctx, driver)
	if err != nil {
		return model.Driver{}, err
	}

	insertedID, ok := res.InsertedID.(bson.ObjectID)
	if !ok {
		return model.Driver{}, errors.New("failed to get inserted ID")
	}

	driver.ID = insertedID
	return driver, nil
}

func (d *driverRepository) ReadAll(ctx context.Context) ([]model.Driver, error) {
	var drivers []model.Driver
	cursor, err := d.DriverCollection.Find(ctx, bson.D{})
	if err != nil {
		return drivers, err
	}

	if err = cursor.All(ctx, &drivers); err != nil {
		return drivers, err
	}

	return drivers, nil
}

func (d *driverRepository) ReadByID(ctx context.Context, id string) (model.Driver, error) {
	// Convert the string ID to bson.ObjectID
	objectID, err := bson.ObjectIDFromHex(id)
	if err != nil {
		return model.Driver{}, errors.New("invalid ID format")
	}

	// Define a filter to find the document by _id
	filter := bson.M{"_id": objectID}

	// Query the collection
	var driver model.Driver
	err = d.DriverCollection.FindOne(ctx, filter).Decode(&driver)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return model.Driver{}, errors.New("driver not found")
		}
		return model.Driver{}, err
	}

	return driver, nil
}

func (d *driverRepository) Update(ctx context.Context, id string, driver model.Driver) (model.Driver, error) {
	// Convert the string ID to bson.ObjectID
	objectID, err := bson.ObjectIDFromHex(id)
	if err != nil {
		return driver, errors.New("invalid ID format")
	}

	// Update timestamp
	driver.UpdatedAt = time.Now()

	// Define a filter to find the document by _id
	filter := bson.M{"_id": objectID}

	// Define an update document
	update := bson.M{"$set": driver}

	// Perform the update operation
	res, err := d.DriverCollection.UpdateOne(ctx, filter, update)
	if err != nil {
		return driver, err
	}

	// Check if any document was updated
	if res.MatchedCount == 0 {
		return driver, errors.New("driver not found")
	}

	return driver, nil
}

func (d *driverRepository) Delete(ctx context.Context, id string) error {
	// Convert the string ID to bson.ObjectID
	objectID, err := bson.ObjectIDFromHex(id)
	if err != nil {
		return errors.New("invalid ID format")
	}

	// Define a filter to find the document by _id
	filter := bson.M{"_id": objectID}

	// Perform the delete operation
	res, err := d.DriverCollection.DeleteOne(ctx, filter)
	if err != nil {
		return err
	}

	// Check if any document was deleted
	if res.DeletedCount == 0 {
		return errors.New("driver not found")
	}

	return nil
}
