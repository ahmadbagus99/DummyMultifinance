package testing

import (
	"DummyMultifinance/domain/models"
	usecases "DummyMultifinance/usecases/users"
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"golang.org/x/crypto/bcrypt"
)

type MockUserRepository struct {
	mock.Mock
}

func (m *MockUserRepository) Insert(ctx context.Context, user *models.Users) (*models.Users, error) {
	args := m.Called(ctx, user)
	return args.Get(0).(*models.Users), args.Error(1)
}

func (m *MockUserRepository) GetUserByID(ctx context.Context, id int) (*models.Users, error) {
	args := m.Called(ctx, id)
	return args.Get(0).(*models.Users), args.Error(1)
}

func (m *MockUserRepository) GetUserByUsername(ctx context.Context, username string) (*models.Users, error) {
	args := m.Called(ctx, username)
	return args.Get(0).(*models.Users), args.Error(1)
}

func TestCreateUser(t *testing.T) {
	mockRepo := new(MockUserRepository)
	userUseCase := usecases.NewUserUsecase(mockRepo)

	user := &models.Users{
		Username: "testuser",
		Password: "password123",
		RoleID:   "admin",
	}

	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	user.Password = string(hashedPassword)

	mockRepo.On("Insert", mock.Anything, user).Return(user, nil)
	createdUser, err := userUseCase.CreateUser(context.Background(), user)

	assert.NoError(t, err)
	assert.NotNil(t, createdUser)
	assert.Equal(t, "testuser", createdUser.Username)
	mockRepo.AssertExpectations(t)
}

func TestGetUserById(t *testing.T) {
	mockRepo := new(MockUserRepository)
	userUseCase := usecases.NewUserUsecase(mockRepo)

	user := &models.Users{
		ID:       1,
		Username: "testuser",
	}

	mockRepo.On("GetUserByID", mock.Anything, 1).Return(user, nil)

	result, err := userUseCase.GetUserById(context.Background(), 1)

	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, "testuser", result.Username)
	mockRepo.AssertExpectations(t)
}

func TestLogin(t *testing.T) {
	mockRepo := new(MockUserRepository)
	userUseCase := usecases.NewUserUsecase(mockRepo)

	originalPassword := "password123"
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(originalPassword), bcrypt.DefaultCost)
	user := &models.Users{
		Username: "testuser",
		Password: string(hashedPassword),
	}

	mockRepo.On("GetUserByUsername", mock.Anything, "testuser").Return(user, nil)

	token, expiration, err := userUseCase.Login(context.Background(), "testuser", originalPassword)

	assert.NoError(t, err)
	assert.NotEmpty(t, token)
	assert.NotEmpty(t, expiration)
	mockRepo.AssertExpectations(t)
}
