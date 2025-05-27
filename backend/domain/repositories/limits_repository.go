package repositories

import (
	common "DummyMultifinance/domain/common"
	"DummyMultifinance/domain/models"
	"context"
	"database/sql"
)

type LimitRepository interface {
	Insert(ctx context.Context, execer common.Execer, limit *models.Limits) (*models.Limits, error)
	GetByID(ctx context.Context, id int) (*models.Limits, error)
	UpdateLimit(ctx context.Context, dbTx *sql.Tx, consumer_id int, tenor int, amount float64) error
	GetDB() *sql.DB
}
