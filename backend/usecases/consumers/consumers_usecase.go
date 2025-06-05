package usecases

import (
	"DummyMultifinance/domain/models"
	"DummyMultifinance/domain/repositories"
	"context"
)

type ConsumerUseCase interface {
	CreateConsumer(ctx context.Context, tx *models.Consumers) (*models.Consumers, error)
	GetAllConsumer(ctx context.Context) (*models.Consumers, error)
	GetConsumerById(ctx context.Context, id int) (*models.Consumers, error)
	GetConsumerLimit(ctx context.Context, consumer_id int) ([]models.ConsumersLimit, error)
}

type consumerUsecase struct {
	repo repositories.ConsumerRepository
}

func NewConsumerUsecase(r repositories.ConsumerRepository) ConsumerUseCase {
	return &consumerUsecase{repo: r}
}

func (uc *consumerUsecase) CreateConsumer(ctx context.Context, tx *models.Consumers) (*models.Consumers, error) {
	return uc.repo.Insert(ctx, tx)
}

func (uc *consumerUsecase) GetAllConsumer(ctx context.Context) (*models.Consumers, error) {
	return uc.repo.GetAllData(ctx)
}

func (uc *consumerUsecase) GetConsumerById(ctx context.Context, id int) (*models.Consumers, error) {
	return uc.repo.GetByID(ctx, id)
}

func (uc *consumerUsecase) GetConsumerLimit(ctx context.Context, consumer_id int) ([]models.ConsumersLimit, error) {
	return uc.repo.GetConsumerLimit(ctx, consumer_id)
}
