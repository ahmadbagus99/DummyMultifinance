package repositories

import (
	"DummyMultifinance/domain/models"
	"context"
)

type TransactionRepository interface {
	GetByID(ctx context.Context, id int) (*models.Transactions, error)
	Insert(ctx context.Context, tx *models.Transactions) (*models.Transactions, error)
}
