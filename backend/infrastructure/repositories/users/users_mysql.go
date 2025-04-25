package infrastructure

import (
	"DummyMultifinance/domain/models"
	"DummyMultifinance/domain/repositories"
	"context"
	"database/sql"
)

type mysqlUserRepo struct {
	DB *sql.DB
}

func NewMysqlUserRepo(db *sql.DB) repositories.UserRepository {
	return &mysqlUserRepo{DB: db}
}

func (m *mysqlUserRepo) Insert(ctx context.Context, tx *models.Users) (*models.Users, error) {
	query := `INSERT INTO users 
        (username, password, role_id, email, created_at)
        VALUES (?, ?, ?, ?, ?)`

	result, err := m.DB.ExecContext(ctx, query,
		tx.Username,
		tx.Password,
		tx.RoleID,
		tx.Email,
		tx.CreatedAt,
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

func (m *mysqlUserRepo) GetUserByID(ctx context.Context, id int) (*models.Users, error) {
	query := `SELECT id, username, password, role_id, email, created_at
              FROM transactions WHERE id = ?`

	row := m.DB.QueryRowContext(ctx, query, id)

	var tx models.Users
	err := row.Scan(
		&tx.ID,
		&tx.Username,
		&tx.Password,
		&tx.RoleID,
		&tx.Email,
		&tx.CreatedAt,
	)
	if err != nil {
		return nil, err
	}

	return &tx, nil
}

func (m *mysqlUserRepo) GetUserByUsername(ctx context.Context, username string) (*models.Users, error) {
	query := `SELECT id, username, password, role_id, email, created_at
              FROM users WHERE username = ?`

	row := m.DB.QueryRowContext(ctx, query, username)

	var user models.Users
	err := row.Scan(
		&user.ID,
		&user.Username,
		&user.Password,
		&user.RoleID,
		&user.Email,
		&user.CreatedAt,
	)
	if err != nil {
		return nil, err
	}

	return &user, nil
}
