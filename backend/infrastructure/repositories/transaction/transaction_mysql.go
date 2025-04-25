package mysql

import (
	"DummyMultifinance/domain/models"
	"DummyMultifinance/domain/repositories"
	"context"
	"database/sql"
)

type mysqlTransactionRepo struct {
	DB *sql.DB
}

func NewMysqlTransactionRepo(db *sql.DB) repositories.TransactionRepository {
	return &mysqlTransactionRepo{DB: db}
}

func (m *mysqlTransactionRepo) Insert(ctx context.Context, tx *models.Transaction) (*models.Transaction, error) {
	query := `INSERT INTO transactions 
        (contract_number, customer_id, otr, admin_fee, installment, interest, asset_name, transaction_date)
        VALUES (?, ?, ?, ?, ?, ?, ?, ?)`

	result, err := m.DB.ExecContext(ctx, query,
		tx.ContractNumber,
		tx.CustomerID,
		tx.OTR,
		tx.AdminFee,
		tx.Installment,
		tx.Interest,
		tx.AssetName,
		tx.TransactionDate,
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

func (m *mysqlTransactionRepo) GetByID(ctx context.Context, id int) (*models.Transaction, error) {
	query := `SELECT id, contract_number, customer_id, otr, admin_fee, installment, interest, asset_name, transaction_date
              FROM transactions WHERE id = ?`

	row := m.DB.QueryRowContext(ctx, query, id)

	var tx models.Transaction
	err := row.Scan(
		&tx.ID,
		&tx.ContractNumber,
		&tx.CustomerID,
		&tx.OTR,
		&tx.AdminFee,
		&tx.Installment,
		&tx.Interest,
		&tx.AssetName,
		&tx.TransactionDate,
	)
	if err != nil {
		return nil, err
	}

	return &tx, nil
}
