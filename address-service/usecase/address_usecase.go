package usecase

import (
	"address-service/model"
	"address-service/repository"
	"context"
	"errors"
	"time"

	"github.com/go-playground/validator/v10"
)

type IAddressUsecase interface {
	Create(ctx context.Context, payload model.PayloadCreateAddress) (model.ResponseAddress, error)
	FindAll(ctx context.Context) ([]model.Address, error)
	FindByID(ctx context.Context, id string) (model.ResponseAddress, error)
	Update(ctx context.Context, id string, payload model.PayloadUpdateAddress) (model.ResponseAddress, error)
	Delete(ctx context.Context, id string) error
}

type addressUsecase struct {
	AddressRepo repository.IAddressRepository
	Validate    *validator.Validate
}

func NewAddressUsecase(addressRepo repository.IAddressRepository, validate *validator.Validate) IAddressUsecase {
	return &addressUsecase{
		AddressRepo: addressRepo,
		Validate:    validate,
	}
}

func (a *addressUsecase) Create(ctx context.Context, payload model.PayloadCreateAddress) (model.ResponseAddress, error) {
	err := a.Validate.Struct(payload)
	if err != nil {
		return model.ResponseAddress{}, err
	}

	address := model.Address{
		Name:      payload.Name,
		Latitude:  payload.Latitude,
		Longitude: payload.Longitude,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	res, err := a.AddressRepo.Create(ctx, address)
	if err != nil {
		return model.ResponseAddress{}, err
	}

	response := model.ResponseAddress{
		ID:        res.ID.Hex(),
		Name:      res.Name,
		Latitude:  res.Latitude,
		Longitude: res.Longitude,
		CreatedAt: res.CreatedAt,
		UpdatedAt: res.UpdatedAt,
	}

	return response, nil
}

func (a *addressUsecase) FindAll(ctx context.Context) ([]model.Address, error) {
	var addresses []model.Address
	addresses, err := a.AddressRepo.ReadAll(ctx)
	if err != nil {
		return addresses, err
	}

	return addresses, nil
}

func (a *addressUsecase) FindByID(ctx context.Context, id string) (model.ResponseAddress, error) {
	var address model.Address
	address, err := a.AddressRepo.ReadByID(ctx, id)
	if err != nil {
		return model.ResponseAddress{}, errors.New("address not found")
	}

	response := model.ResponseAddress{
		ID:        address.ID.Hex(),
		Name:      address.Name,
		Latitude:  address.Latitude,
		Longitude: address.Longitude,
		CreatedAt: address.CreatedAt,
		UpdatedAt: address.UpdatedAt,
	}

	return response, nil
}

func (a *addressUsecase) Update(ctx context.Context, id string, payload model.PayloadUpdateAddress) (model.ResponseAddress, error) {
	err := a.Validate.Struct(payload)
	if err != nil {
		return model.ResponseAddress{}, err
	}

	address, err := a.AddressRepo.ReadByID(ctx, id)
	if err != nil {
		return model.ResponseAddress{}, errors.New("address not found")
	}

	payloadUpdated := model.Address{
		ID:        address.ID,
		Name:      payload.Name,
		Latitude:  payload.Latitude,
		Longitude: payload.Longitude,
		CreatedAt: address.CreatedAt,
		UpdatedAt: time.Now(),
	}

	updatedAddress, err := a.AddressRepo.Update(ctx, id, payloadUpdated)
	if err != nil {
		return model.ResponseAddress{}, errors.New("error updating address")
	}

	response := model.ResponseAddress{
		ID:        updatedAddress.ID.Hex(),
		Name:      updatedAddress.Name,
		Latitude:  updatedAddress.Latitude,
		Longitude: updatedAddress.Longitude,
		CreatedAt: updatedAddress.CreatedAt,
		UpdatedAt: updatedAddress.UpdatedAt,
	}

	return response, nil
}

func (a *addressUsecase) Delete(ctx context.Context, id string) error {
	return a.AddressRepo.Delete(ctx, id)
}
