package testing

import (
	"testing"

	userUseCase "DummyMultifinance/usecases/users"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type UserUseCaseIntegrationTestSuite struct {
	suite.Suite
	userUseCase *userUseCase.UserUseCase
}

// func (suite *UserUseCaseIntegrationTestSuite) SetupSuite() {
// 	// Set up database or mock database connection here
// 	os.Setenv("SECRET_KEY", "supersecretkey")
// 	db := database.SetupDB()
// 	repo := &database.UserRepo{DB: db}
// 	suite.userUseCase = &usecases.UserUseCase{UserRepo: repo}
// }

func (suite *UserUseCaseIntegrationTestSuite) TestCreateUser() {
	Users, err := suite.userUseCase.CreateUser("testuser", "password")
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), "testuser", Users.Username)
}

func (suite *UserUseCaseIntegrationTestSuite) TestLogin() {
	// Creating user first
	suite.userUseCase.CreateUser("testuser", "password")

	token, expiration, err := suite.userUseCase.Login("testuser", "password")

	assert.NoError(suite.T(), err)
	assert.NotEmpty(suite.T(), token)
	assert.NotEmpty(suite.T(), expiration)
}

func TestUserUseCaseIntegrationTestSuite(t *testing.T) {
	suite.Run(t, new(UserUseCaseIntegrationTestSuite))
}
