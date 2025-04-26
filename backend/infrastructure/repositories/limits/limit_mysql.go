package infrastructure

import (
	"DummyMultifinance/domain/models"
	"DummyMultifinance/domain/repositories"
	"context"
	"database/sql"
)

type mysqlLimitRepo struct {
	DB *sql.DB
}

func NewMysqlLimitRepo(db *sql.DB) repositories.LimitRepository {
	return &mysqlLimitRepo{DB: db}
}

func (m *mysqlLimitRepo) Insert(ctx context.Context, tx *models.Limits) (*models.Limits, error) {
	query := `INSERT INTO limits 
        (consumer_id, limit_1, limit_2, limit_3, limit_6)
        VALUES (?, ?, ?, ?, ?)`

	result, err := m.DB.ExecContext(ctx, query,
		tx.ConsumerID,
		tx.Limit1,
		tx.Limit2,
		tx.Limit3,
		tx.Limit6,
	)
	if err != nil {
		return nil, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}

	tx.ID = int(id)
	return tx, nil
}

func (m *mysqlLimitRepo) GetByID(ctx context.Context, id int) (*models.Limits, error) {
	query := `SELECT id, consumer_id, limit_1, limit_2, limit_3, limit_6
              FROM transactions WHERE id = ?`

	row := m.DB.QueryRowContext(ctx, query, id)

	var tx models.Limits
	err := row.Scan(
		&tx.ID,
		&tx.ConsumerID,
		&tx.Limit1,
		&tx.Limit2,
		&tx.Limit3,
		&tx.Limit6,
	)
	if err != nil {
		return nil, err
	}

	return &tx, nil
}
