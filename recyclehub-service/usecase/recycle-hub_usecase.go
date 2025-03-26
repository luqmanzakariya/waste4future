package usecase

import (
	"context"
	"errors"
	"reyclehub-service/model"
	"reyclehub-service/repository"
	"time"

	"github.com/go-playground/validator/v10"
)

type IRecycleHubUsecase interface {
	Create(ctx context.Context, payload model.PayloadCreateRecycleHub) (model.ResponseRecycleHub, error)
	FindAll(ctx context.Context) ([]model.RecycleHub, error)
	FindByID(ctx context.Context, id string) (model.ResponseRecycleHub, error)
	Update(ctx context.Context, id string, payload model.PayloadUpdateRecycleHub) (model.ResponseRecycleHub, error)
	Delete(ctx context.Context, id string) error
}

type recycleHubUsecase struct {
	RecycleHubRepo repository.IRecycleHubRepository
	Validate       *validator.Validate
}

func NewRecycleHubUsecase(recycleHubRepo repository.IRecycleHubRepository, validate *validator.Validate) IRecycleHubUsecase {
	return &recycleHubUsecase{
		RecycleHubRepo: recycleHubRepo,
		Validate:       validate,
	}
}

func (r *recycleHubUsecase) Create(ctx context.Context, payload model.PayloadCreateRecycleHub) (model.ResponseRecycleHub, error) {
	err := r.Validate.Struct(payload)
	if err != nil {
		return model.ResponseRecycleHub{}, err
	}

	recycleHub := model.RecycleHub{
		Name:        payload.Name,
		Phone:       payload.Phone,
		AddressID:   payload.AddressID,
		WasteTypeID: payload.WasteTypeID,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	res, err := r.RecycleHubRepo.Create(ctx, recycleHub)
	if err != nil {
		return model.ResponseRecycleHub{}, err
	}

	response := model.ResponseRecycleHub{
		ID:          res.ID.Hex(),
		Name:        res.Name,
		Phone:       res.Phone,
		AddressID:   res.AddressID,
		WasteTypeID: res.WasteTypeID,
		CreatedAt:   res.CreatedAt,
		UpdatedAt:   res.UpdatedAt,
	}

	return response, nil
}

func (r *recycleHubUsecase) FindAll(ctx context.Context) ([]model.RecycleHub, error) {
	return r.RecycleHubRepo.ReadAll(ctx)
}

func (r *recycleHubUsecase) FindByID(ctx context.Context, id string) (model.ResponseRecycleHub, error) {
	recycleHub, err := r.RecycleHubRepo.ReadByID(ctx, id)
	if err != nil {
		return model.ResponseRecycleHub{}, errors.New("recycle hub not found")
	}

	response := model.ResponseRecycleHub{
		ID:          recycleHub.ID.Hex(),
		Name:        recycleHub.Name,
		Phone:       recycleHub.Phone,
		AddressID:   recycleHub.AddressID,
		WasteTypeID: recycleHub.WasteTypeID,
		CreatedAt:   recycleHub.CreatedAt,
		UpdatedAt:   recycleHub.UpdatedAt,
	}

	return response, nil
}

func (r *recycleHubUsecase) Update(ctx context.Context, id string, payload model.PayloadUpdateRecycleHub) (model.ResponseRecycleHub, error) {
	err := r.Validate.Struct(payload)
	if err != nil {
		return model.ResponseRecycleHub{}, err
	}

	recycleHub, err := r.RecycleHubRepo.ReadByID(ctx, id)
	if err != nil {
		return model.ResponseRecycleHub{}, errors.New("recycle hub not found")
	}

	payloadUpdated := model.RecycleHub{
		ID:          recycleHub.ID,
		Name:        payload.Name,
		Phone:       payload.Phone,
		AddressID:   payload.AddressID,
		WasteTypeID: payload.WasteTypeID,
		CreatedAt:   recycleHub.CreatedAt,
		UpdatedAt:   time.Now(),
	}

	updatedRecycleHub, err := r.RecycleHubRepo.Update(ctx, id, payloadUpdated)
	if err != nil {
		return model.ResponseRecycleHub{}, errors.New("error updating recycle hub")
	}

	response := model.ResponseRecycleHub{
		ID:          updatedRecycleHub.ID.Hex(),
		Name:        updatedRecycleHub.Name,
		Phone:       updatedRecycleHub.Phone,
		AddressID:   updatedRecycleHub.AddressID,
		WasteTypeID: updatedRecycleHub.WasteTypeID,
		CreatedAt:   updatedRecycleHub.CreatedAt,
		UpdatedAt:   updatedRecycleHub.UpdatedAt,
	}

	return response, nil
}

func (r *recycleHubUsecase) Delete(ctx context.Context, id string) error {
	return r.RecycleHubRepo.Delete(ctx, id)
}
