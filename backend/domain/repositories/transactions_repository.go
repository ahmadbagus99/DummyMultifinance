package repositories

import (
	"DummyMultifinance/domain/models"
	"context"
	"database/sql"
)

type TransactionRepository interface {
	GetByID(ctx context.Context, id int) (*models.Transactions, error)
	Insert(ctx context.Context, dbTx *sql.Tx, tx *models.Transactions) (*models.Transactions, error)
	GetDB() *sql.DB
}
