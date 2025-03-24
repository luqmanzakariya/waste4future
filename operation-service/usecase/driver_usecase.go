package usecase

import (
	"context"
	"errors"
	"operation-service/model"
	"operation-service/repository"
	"time"

	"github.com/go-playground/validator/v10"
)

type IDriverUsecase interface {
	Create(ctx context.Context, payload model.PayloadCreateDriver) (model.ResponseDriver, error)
	FindAllActive(ctx context.Context) ([]model.Driver, error)
	FindAll(ctx context.Context) ([]model.Driver, error)
	FindByID(ctx context.Context, id string) (model.ResponseDriver, error)
	Update(ctx context.Context, id string, payload model.PayloadUpdateDriver) (model.ResponseDriver, error)
	Delete(ctx context.Context, id string) error
}

type driverUsecase struct {
	DriverRepo repository.IDriverRepository
	Validate   *validator.Validate
}

func NewDriverUsecase(driverRepo repository.IDriverRepository, validate *validator.Validate) IDriverUsecase {
	return &driverUsecase{
		DriverRepo: driverRepo,
		Validate:   validate,
	}
}

func (d *driverUsecase) Create(ctx context.Context, payload model.PayloadCreateDriver) (model.ResponseDriver, error) {
	err := d.Validate.Struct(payload)
	if err != nil {
		return model.ResponseDriver{}, err
	}

	driver := model.Driver{
		Name:         payload.Name,
		Phone:        payload.Phone,
		LicensePlate: payload.LicensePlate,
		Status:       model.DriverStatusActive,
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	}

	res, err := d.DriverRepo.Create(ctx, driver)
	if err != nil {
		return model.ResponseDriver{}, err
	}

	response := model.ResponseDriver{
		ID:           res.ID.Hex(),
		Name:         res.Name,
		Phone:        res.Phone,
		LicensePlate: res.LicensePlate,
		Status:       res.Status,
		CreatedAt:    res.CreatedAt,
		UpdatedAt:    res.UpdatedAt,
	}

	return response, nil
}

func (d *driverUsecase) FindAllActive(ctx context.Context) ([]model.Driver, error) {
	var drivers []model.Driver
	drivers, err := d.DriverRepo.ReadAllActive(ctx)
	if err != nil {
		return drivers, err
	}

	return drivers, nil
}

func (d *driverUsecase) FindAll(ctx context.Context) ([]model.Driver, error) {
	var drivers []model.Driver
	drivers, err := d.DriverRepo.ReadAll(ctx)
	if err != nil {
		return drivers, err
	}

	return drivers, nil
}

func (d *driverUsecase) FindByID(ctx context.Context, id string) (model.ResponseDriver, error) {
	var driver model.Driver
	driver, err := d.DriverRepo.ReadByID(ctx, id)
	if err != nil {
		return model.ResponseDriver{}, errors.New("driver not found")
	}

	response := model.ResponseDriver{
		ID:           driver.ID.Hex(),
		Name:         driver.Name,
		Phone:        driver.Phone,
		LicensePlate: driver.LicensePlate,
		Status:       driver.Status,
		CreatedAt:    driver.CreatedAt,
		UpdatedAt:    driver.UpdatedAt,
	}

	return response, nil
}

func (d *driverUsecase) Update(ctx context.Context, id string, payload model.PayloadUpdateDriver) (model.ResponseDriver, error) {
	err := d.Validate.Struct(payload)
	if err != nil {
		return model.ResponseDriver{}, err
	}

	driver, err := d.DriverRepo.ReadByID(ctx, id)
	if err != nil {
		return model.ResponseDriver{}, errors.New("driver not found")
	}

	payloadUpdated := model.Driver{
		ID:           driver.ID,
		Name:         payload.Name,
		Phone:        payload.Phone,
		LicensePlate: payload.LicensePlate,
		Status:       payload.Status,
		CreatedAt:    driver.CreatedAt,
		UpdatedAt:    time.Now(),
	}

	updatedOrder, err := d.DriverRepo.Update(ctx, id, payloadUpdated)
	if err != nil {
		return model.ResponseDriver{}, errors.New("error updating driver")
	}

	response := model.ResponseDriver{
		ID:           updatedOrder.ID.Hex(),
		Name:         updatedOrder.Name,
		Phone:        updatedOrder.Phone,
		LicensePlate: updatedOrder.LicensePlate,
		Status:       updatedOrder.Status,
		CreatedAt:    updatedOrder.CreatedAt,
		UpdatedAt:    updatedOrder.UpdatedAt,
	}

	return response, nil
}

func (d *driverUsecase) Delete(ctx context.Context, id string) error {
	return d.DriverRepo.Delete(ctx, id)
}
