package infrastructure

import (
	common "DummyMultifinance/domain/common"
	"DummyMultifinance/domain/models"
	"DummyMultifinance/domain/repositories"
	"context"
)

type mysqlUserRepo struct {
	DB common.Execer
}

func NewMysqlUserRepo(db common.Execer) repositories.UserRepository {
	return &mysqlUserRepo{DB: db}
}

func (m *mysqlUserRepo) Insert(ctx context.Context, user *models.Users) (*models.Users, error) {
	query := `INSERT INTO users 
        (username, password, role_id, email, created_at)
        VALUES (?, ?, ?, ?, ?)`

	result, err := m.DB.ExecContext(ctx, query,
		user.Username,
		user.Password,
		user.RoleID,
		user.Email,
		user.CreatedAt,
	)
	if err != nil {
		return nil, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}

	user.ID = int(id)
	return user, nil
}

func (m *mysqlUserRepo) GetUserByID(ctx context.Context, id int) (*models.Users, error) {
	query := `SELECT id, username, password, role_id, email, created_at
              FROM users WHERE id = ?`

	row := m.DB.QueryRowContext(ctx, query, id)

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
