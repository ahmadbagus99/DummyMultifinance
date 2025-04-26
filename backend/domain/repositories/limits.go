package repositories

import (
	"DummyMultifinance/domain/models"
	"context"
)

type LimitRepository interface {
	GetByID(ctx context.Context, id int) (*models.Limits, error)
	Insert(ctx context.Context, tx *models.Limits) (*models.Limits, error)
}
