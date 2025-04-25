package testing

import (
	"DummyMultifinance/domain/models"

	"github.com/stretchr/testify/mock"
)

type MockUserRepo struct {
	mock.Mock
}

func (m *MockUserRepo) CreateUser(user *models.Users) (*models.Users, error) {
	args := m.Called(user)
	return args.Get(0).(*models.Users), args.Error(1)
}

func (m *MockUserRepo) GetUserByUsername(username string) (*models.Users, error) {
	args := m.Called(username)
	return args.Get(0).(*models.Users), args.Error(1)
}

func (m *MockUserRepo) GetByID(id int) (*models.Users, error) {
	args := m.Called(id)
	return args.Get(0).(*models.Users), args.Error(1)
}

// func TestCreateUser_Success(t *testing.T) {
// 	mockRepo := new(MockUserRepo)
// 	userUseCase := &userUseCase.UserUseCase{UserRepo: mockRepo}
// 	mockRepo.On("CreateUser", mock.Anything).Return(&models.Users{Username: "testuser", Password: "hashedPassword"}, nil)

// 	user, err := userUseCase.CreateUser("testuser", "password")

// 	assert.NoError(t, err)
// 	assert.Equal(t, "testuser", user.Username)
// 	mockRepo.AssertExpectations(t)
// }

// func TestCreateUser_Failure(t *testing.T) {
// 	mockRepo := new(MockUserRepo)
// 	userUseCase := &userUseCase.UserUseCase{UserRepo: mockRepo}

// 	mockRepo.On("CreateUser", mock.Anything).Return(nil, errors.New("error creating user"))
// 	user, err := userUseCase.CreateUser("testuser", "password")

// 	assert.Error(t, err)
// 	assert.Nil(t, user)
// 	mockRepo.AssertExpectations(t)
// }

// func TestLogin_Success(t *testing.T) {
// 	validPassword := "password"
// 	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(validPassword), bcrypt.DefaultCost)
// 	if err != nil {
// 		t.Fatalf("Error generating bcrypt hash: %v", err)
// 	}

// 	mockRepo := new(MockUserRepo)
// 	userUseCase := &userUseCase.UserUseCase{UserRepo: mockRepo}

// 	mockRepo.On("GetUserByUsername", "testuser").Return(&models.Users{Username: "testuser", Password: string(hashedPassword)}, nil)
// 	token, expiration, err := userUseCase.Login("testuser", validPassword)

// 	assert.NoError(t, err)
// 	assert.NotEmpty(t, token)
// 	assert.NotEmpty(t, expiration)

// 	mockRepo.AssertExpectations(t)
// }

// func TestLogin_UserNotFound(t *testing.T) {
// 	mockRepo := new(MockUserRepo)
// 	userUseCase := &userUseCase.UserUseCase{UserRepo: mockRepo}

// 	mockRepo.On("GetUserByUsername", "testuser").Return(nil, errors.New("user not found"))
// 	token, expiration, err := userUseCase.Login("testuser", "password")

// 	assert.Error(t, err)
// 	assert.Empty(t, token)
// 	assert.Empty(t, expiration)
// 	mockRepo.AssertExpectations(t)
// }
