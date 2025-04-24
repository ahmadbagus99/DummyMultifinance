package repositories

import "DummyMultifinance/domain/models"

type UserRepository interface {
	CreateUser(user *models.User) (*models.User, error)
	GetUserByUsername(username string) (*models.User, error)
}
