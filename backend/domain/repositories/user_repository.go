package repositories

import "DummyMultifinance/domain/models"

type UserRepository interface {
	GetByID(id int) (*models.User, error)
	CreateUser(user *models.User) (*models.User, error)
	GetUserByUsername(username string) (*models.User, error)
}
