package repositories

import (
	"DummyMultifinance/domain/models"
	"context"
)

type ConsumerRepository interface {
	GetByID(ctx context.Context, id int) (*models.Consumers, error)
	Insert(ctx context.Context, tx *models.Consumers) (*models.Consumers, error)
}
