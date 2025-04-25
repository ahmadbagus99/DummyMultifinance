package usecases

import (
	"DummyMultifinance/domain/models"
	"DummyMultifinance/domain/repositories"
	"context"
)

type ConsumerUseCase interface {
	CreateConsumer(ctx context.Context, tx *models.Consumers) (*models.Consumers, error)
	GetConsumerById(ctx context.Context, id int) (*models.Consumers, error)
}

type consumerUsecase struct {
	repo repositories.ConsumerRepository
}

func NewConsumerUsecase(r repositories.ConsumerRepository) ConsumerUseCase {
	return &consumerUsecase{repo: r}
}

func (uc *consumerUsecase) GetConsumerById(ctx context.Context, id int) (*models.Consumers, error) {
	return uc.repo.GetByID(ctx, id)
}

func (uc *consumerUsecase) CreateConsumer(ctx context.Context, tx *models.Consumers) (*models.Consumers, error) {
	return uc.repo.Insert(ctx, tx)
}
