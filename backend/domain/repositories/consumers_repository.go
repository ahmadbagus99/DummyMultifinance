package repositories

import (
	"DummyMultifinance/domain/models"
	"context"
)

type ConsumerRepository interface {
	Insert(ctx context.Context, tx *models.Consumers) (*models.Consumers, error)
	GetAllData(ctx context.Context) (*models.Consumers, error)
	GetByID(ctx context.Context, id int) (*models.Consumers, error)
	GetConsumerLimit(ctx context.Context, consumer_id int) ([]models.ConsumersLimit, error)
}
