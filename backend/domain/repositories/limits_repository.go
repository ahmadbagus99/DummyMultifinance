package repositories

import (
	"DummyMultifinance/domain/models"
	"context"
)

type LimitRepository interface {
	Insert(ctx context.Context, tx *models.Limits) (*models.Limits, error)
	GetByID(ctx context.Context, id int) (*models.Limits, error)
	UpdateLimit(ctx context.Context, consumer_id int, tenor int, amount float64) error
}
