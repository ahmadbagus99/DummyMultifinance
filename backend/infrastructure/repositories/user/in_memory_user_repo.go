package infrastructure

import (
	"DummyMultifinance/domain/models"
	"fmt"
)

type InMemoryUserRepo struct {
	users map[int]*models.User
}

func NewInMemoryUserRepo() *InMemoryUserRepo {
	return &InMemoryUserRepo{
		users: make(map[int]*models.User),
	}
}

func (repo *InMemoryUserRepo) CreateUser(user *models.User) (*models.User, error) {
	user.ID = len(repo.users) + 1
	repo.users[user.ID] = user
	return user, nil
}

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
