package usecase

import (
	"context"
	"errors"
	"recyclehub-service/helper"
	"recyclehub-service/model/web/response"
	"recyclehub-service/repository"

	"github.com/go-playground/validator/v10"
)

type IWasteTypeUsecase interface {
	FindAll(ctx context.Context) ([]response.WasteTypeResponse, error)
	FindById(ctx context.Context, id string) (response.WasteTypeResponse, error)
}

type wasteTypeUsecase struct {
	WasteTypeRepo repository.IWasteTypeRepository
	Validate      *validator.Validate
}

func NewWasteTypeUsecase(wasteTypeRepo repository.IWasteTypeRepository, validate *validator.Validate) IWasteTypeUsecase {
	return &wasteTypeUsecase{
		WasteTypeRepo: wasteTypeRepo,
		Validate:      validate,
	}
}

func (u *wasteTypeUsecase) FindAll(ctx context.Context) ([]response.WasteTypeResponse, error) {
	wasteTypes, err := u.WasteTypeRepo.FindAll(ctx)
	if err != nil {
		return []response.WasteTypeResponse{}, err
	}
	return helper.WasteTypeToResponses(wasteTypes), nil
}

func (u *wasteTypeUsecase) FindById(ctx context.Context, id string) (response.WasteTypeResponse, error) {
	wasteType, err := u.WasteTypeRepo.FindById(ctx, id)
	if err != nil {
		return response.WasteTypeResponse{}, errors.New("waste type not found")
	}
	return helper.WasteTypeToResponse(wasteType), nil
}
