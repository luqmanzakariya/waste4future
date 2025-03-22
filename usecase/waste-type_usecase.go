package usecase

import (
	"context"
	"recyclehub-service/model/domain"
	"recyclehub-service/repository"
)

type IWasteTypeUsecase interface {
	FindAll(ctx context.Context) ([]domain.WasteType, error)
	FindById(ctx context.Context, id string) (domain.WasteType, error)
}

type wasteTypeUsecase struct {
	repo repository.WasteTypeRepository
}

func NewWasteTypeUsecase(repo repository.WasteTypeRepository) IWasteTypeUsecase {
	return &wasteTypeUsecase{repo: repo}
}

func (u *wasteTypeUsecase) FindAll(ctx context.Context) ([]domain.WasteType, error) {
	return u.repo.FindAll(ctx)
}

func (u *wasteTypeUsecase) FindById(ctx context.Context, id string) (domain.WasteType, error) {
	return u.repo.FindById(ctx, id)
}
