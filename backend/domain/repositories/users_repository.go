package repositories

import "DummyMultifinance/domain/models"

type UserRepository interface {
	GetByID(id int) (*models.Users, error)
	CreateUser(user *models.Users) (*models.Users, error)
	GetUserByUsername(username string) (*models.Users, error)
}
