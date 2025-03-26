package usecase

import (
	"context"
	"errors"
	"reyclehub-service/model"
	"reyclehub-service/repository"
)

type IWasteTypeUsecase interface {
	FindAll(ctx context.Context) ([]model.WasteType, error)
	FindByID(ctx context.Context, id string) (model.WasteType, error)
}

type wasteTypeUsecase struct {
	WasteTypeRepo repository.IWasteTypeRepository
}

func NewWasteTypeUsecase(wasteTypeRepo repository.IWasteTypeRepository) IWasteTypeUsecase {
	return &wasteTypeUsecase{
		WasteTypeRepo: wasteTypeRepo,
	}
}

func (w *wasteTypeUsecase) FindAll(ctx context.Context) ([]model.WasteType, error) {
	return w.WasteTypeRepo.ReadAll(ctx)
}

func (w *wasteTypeUsecase) FindByID(ctx context.Context, id string) (model.WasteType, error) {
	wasteType, err := w.WasteTypeRepo.ReadByID(ctx, id)
	if err != nil {
		return model.WasteType{}, errors.New("waste type not found")
	}
	return wasteType, nil
}
