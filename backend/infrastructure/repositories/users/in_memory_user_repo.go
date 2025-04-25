package infrastructure

import (
	"DummyMultifinance/domain/models"
	"fmt"
)

type InMemoryUserRepo struct {
	users map[int]*models.Users
}

func NewInMemoryUserRepo() *InMemoryUserRepo {
	return &InMemoryUserRepo{
		users: make(map[int]*models.Users),
	}
}

func (repo *InMemoryUserRepo) CreateUser(user *models.Users) (*models.Users, error) {
	user.ID = len(repo.users) + 1
	repo.users[user.ID] = user
	return user, nil
}

func (repo *InMemoryUserRepo) GetUserByUsername(username string) (*models.Users, error) {
	for _, user := range repo.users {
		if user.Username == username {
			return user, nil
		}
	}
	return nil, fmt.Errorf("user not found")
}

func (repo *InMemoryUserRepo) GetByID(id int) (*models.Users, error) {
	for _, user := range repo.users {
		if user.ID == id {
			return user, nil
		}
	}
	return nil, fmt.Errorf("user not found")
}
