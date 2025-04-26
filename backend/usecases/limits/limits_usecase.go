package usecases

import (
	"DummyMultifinance/domain/models"
	"DummyMultifinance/domain/repositories"
	"context"
)

type LimiUseCase interface {
	GetLimitById(ctx context.Context, id int) (*models.Limits, error)
	CreateLimit(ctx context.Context, tx *models.Limits) (*models.Limits, error)
}

type limitUsecase struct {
	repo repositories.LimitRepository
}

func NewTransactionUsecase(r repositories.LimitRepository) LimiUseCase {
	return &limitUsecase{repo: r}
}

func (uc *limitUsecase) GetLimitById(ctx context.Context, id int) (*models.Limits, error) {
	return uc.repo.GetByID(ctx, id)
}

func (uc *limitUsecase) CreateLimit(ctx context.Context, tx *models.Limits) (*models.Limits, error) {
	return uc.repo.Insert(ctx, tx)
}
