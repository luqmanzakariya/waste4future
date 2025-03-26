package usecase

import (
	"context"
	"errors"
	"recyclehub-service/helper"
	"recyclehub-service/model/domain"
	web "recyclehub-service/model/web/request"
	"recyclehub-service/model/web/response"
	"recyclehub-service/repository"

	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type IRecycleHubUsecase interface {
	FindAll(ctx context.Context) ([]response.RecycleHubResponse, error)
	FindById(ctx context.Context, id string) (response.RecycleHubResponse, error)
	Create(ctx context.Context, payload web.RecycleHubCreateRequest) (response.RecycleHubResponse, error)
	Update(ctx context.Context, payload web.RecycleHubUpdateRequest, id string) (response.RecycleHubResponse, error)
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

func (u *recycleHubUsecase) FindAll(ctx context.Context) ([]response.RecycleHubResponse, error) {
	recycleHubs, err := u.RecycleHubRepo.FindAll(ctx)
	if err != nil {
		return []response.RecycleHubResponse{}, err
	}
	return helper.RecycleHubToResponses(recycleHubs), nil
}

func (u *recycleHubUsecase) FindById(ctx context.Context, id string) (response.RecycleHubResponse, error) {
	recycleHub, err := u.RecycleHubRepo.FindByID(ctx, id)
	if err != nil {
		return response.RecycleHubResponse{}, errors.New("recycle hub not found")
	}
	return helper.RecycleHubToResponse(*recycleHub), nil
}

func (u *recycleHubUsecase) Create(ctx context.Context, payload web.RecycleHubCreateRequest) (response.RecycleHubResponse, error) {
	err := u.Validate.Struct(payload)
	if err != nil {
		return response.RecycleHubResponse{}, err
	}

	newRecycleHub := &domain.RecycleHub{
		ID:          primitive.NewObjectID(), // Explicitly generate a new ObjectId
		Name:        payload.Name,
		Phone:       payload.Phone,
		AddressID:   payload.AddressID,
		WasteTypeID: payload.WasteTypeID,
	}

	recycleHubCreated, err := u.RecycleHubRepo.Create(ctx, newRecycleHub)
	if err != nil {
		return response.RecycleHubResponse{}, err
	}

	return helper.RecycleHubToResponse(*recycleHubCreated), nil
}

func (u *recycleHubUsecase) Update(ctx context.Context, payload web.RecycleHubUpdateRequest, id string) (response.RecycleHubResponse, error) {
	err := u.Validate.Struct(payload)
	if err != nil {
		return response.RecycleHubResponse{}, err
	}

	_, err = u.RecycleHubRepo.FindByID(ctx, id)
	if err != nil {
		return response.RecycleHubResponse{}, errors.New("recycle hub not found")
	}

	updatedRecycleHub := &domain.RecycleHub{
		Name:        payload.Name,
		Phone:       payload.Phone,
		AddressID:   payload.AddressID,
		WasteTypeID: payload.WasteTypeID,
	}

	updatedRecycleHub, err = u.RecycleHubRepo.Update(ctx, id, updatedRecycleHub)
	if err != nil {
		return response.RecycleHubResponse{}, err
	}

	return helper.RecycleHubToResponse(*updatedRecycleHub), nil
}

func (u *recycleHubUsecase) Delete(ctx context.Context, id string) error {
	recycleHub, err := u.RecycleHubRepo.FindByID(ctx, id)
	if err != nil {
		return errors.New("recycle hub not found")
	}

	return u.RecycleHubRepo.Delete(ctx, recycleHub.ID.Hex())
}
