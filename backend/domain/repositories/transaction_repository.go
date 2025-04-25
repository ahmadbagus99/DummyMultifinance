package repositories

import (
	"DummyMultifinance/domain/models"
	"context"
)

type TransactionRepository interface {
	GetByID(ctx context.Context, id int) (*models.Transaction, error)
	Insert(ctx context.Context, tx *models.Transaction) (*models.Transaction, error)
}
