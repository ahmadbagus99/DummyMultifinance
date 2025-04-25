package repositories

import (
	"DummyMultifinance/domain/models"
	"context"
)

type UserRepository interface {
	Insert(ctx context.Context, tx *models.Users) (*models.Users, error)
	GetUserByID(ctx context.Context, id int) (*models.Users, error)
	GetUserByUsername(ctx context.Context, username string) (*models.Users, error)
}
