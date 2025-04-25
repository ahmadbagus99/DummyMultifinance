package infrastructure

import (
	"DummyMultifinance/domain/models"
	"fmt"
)

// InMemoryUserRepo adalah implementasi repository untuk user yang disimpan dalam memori
type InMemoryUserRepo struct {
	users map[int]*models.User
}

// NewInMemoryUserRepo membuat instance baru dari InMemoryUserRepo
func NewInMemoryUserRepo() *InMemoryUserRepo {
	return &InMemoryUserRepo{
		users: make(map[int]*models.User),
	}
}

// CreateUser menyimpan user baru ke memori
func (repo *InMemoryUserRepo) CreateUser(user *models.User) (*models.User, error) {
	user.ID = len(repo.users) + 1
	repo.users[user.ID] = user
	return user, nil
}

// GetUserByUsername mencari user berdasarkan username
func (repo *InMemoryUserRepo) GetUserByUsername(username string) (*models.User, error) {
	for _, user := range repo.users {
		if user.Username == username {
			return user, nil
		}
	}
	return nil, fmt.Errorf("user not found")
}

func (repo *InMemoryUserRepo) GetByID(id int) (*models.User, error) {
	for _, user := range repo.users {
		if user.ID == id {
			return user, nil
		}
	}
	return nil, fmt.Errorf("user not found")
}
