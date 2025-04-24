package usecases

import (
	"DummyMultifinance/domain/models"
	"DummyMultifinance/domain/repositories"
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
)

var jwtKey = []byte(os.Getenv("SECRET_KEY"))

type Claims struct {
	Username string `json:"username"`
	jwt.RegisteredClaims
}

type UserUseCase struct {
	UserRepo repositories.UserRepository
}

func (uc *UserUseCase) CreateUser(username, password string) (*models.User, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, fmt.Errorf("failed to hash password: %v", err)
	}

	user := &models.User{
		Username: username,
		Password: string(hashedPassword),
	}

	return uc.UserRepo.CreateUser(user)
}

func (uc *UserUseCase) Login(username, password string) (string, string, error) {
	user, err := uc.UserRepo.GetUserByUsername(username)
	if err != nil {
		return "", "", fmt.Errorf("user not found: %v", err)
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return "", "", fmt.Errorf("invalid password: %v", err)
	}

	expirationTime := time.Now().Add(24 * time.Hour)
	expirationString := expirationTime.Format("2006-01-02 15:04:05")
	claims := &Claims{
		Username: username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)

	if err != nil {
		return "", "", fmt.Errorf("failed to sign token: %v", err)
	}

	return tokenString, expirationString, nil
}
