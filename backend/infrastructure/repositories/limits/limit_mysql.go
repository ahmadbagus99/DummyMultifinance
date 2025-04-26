package infrastructure

import (
	"DummyMultifinance/domain/models"
	"DummyMultifinance/domain/repositories"
	"context"
	"database/sql"
	"fmt"
)

type mysqlLimitRepo struct {
	DB *sql.DB
}

func NewMysqlLimitRepo(db *sql.DB) repositories.LimitRepository {
	return &mysqlLimitRepo{DB: db}
}

func (m *mysqlLimitRepo) Insert(ctx context.Context, tx *models.Limits) (*models.Limits, error) {
	query := `INSERT INTO limits 
        (consumer_id, tenor_id, amount)
        VALUES (?, ?, ?)`

	result, err := m.DB.ExecContext(ctx, query,
		tx.ConsumerID,
		tx.TenorID,
		tx.Amount,
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
	query := `SELECT id, consumer_id, tenor_id, amount
              FROM limits WHERE id = ?`

	row := m.DB.QueryRowContext(ctx, query, id)

	var tx models.Limits
	err := row.Scan(
		&tx.ID,
		&tx.ConsumerID,
		&tx.TenorID,
		&tx.Amount,
	)
	if err != nil {
		return nil, err
	}

	return &tx, nil
}

func (m *mysqlLimitRepo) UpdateLimit(ctx context.Context, consumer_id int, tenor int, amount float64) error {
	var limit float64
	query := `SELECT amount FROM limits WHERE consumer_id = ? AND tenor_id = ?`

	row := m.DB.QueryRowContext(ctx, query, consumer_id, tenor)
	err := row.Scan(&limit)
	if err != nil {
		if err == sql.ErrNoRows {
			return fmt.Errorf("no limit found for consumer_id %d and tenor %d", consumer_id, tenor)
		}
		return fmt.Errorf("failed to get limit: %v", err)
	}

	if amount > limit {
		return fmt.Errorf("amount exceeds available limit")
	}

	newLimit := limit - amount
	updateQuery := `UPDATE limits SET amount = ? WHERE consumer_id = ? AND tenor_id = ?`

	_, err = m.DB.ExecContext(ctx, updateQuery, newLimit, consumer_id, tenor)
	if err != nil {
		return fmt.Errorf("failed to update limit: %v", err)
	}

	return nil
}
