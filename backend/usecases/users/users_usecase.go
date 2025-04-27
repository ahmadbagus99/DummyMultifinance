package usecases

import (
	"DummyMultifinance/domain/models"
	"DummyMultifinance/domain/repositories"
	"context"
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

type UserUseCase interface {
	CreateUser(ctx context.Context, tx *models.Users) (*models.Users, error)
	GetUserById(ctx context.Context, id int) (*models.Users, error)
	Login(ctx context.Context, username, password string) (string, string, error)
}

type userUseCase struct {
	repo repositories.UserRepository
}

func NewUserUsecase(r repositories.UserRepository) UserUseCase {
	return &userUseCase{repo: r}
}

func (uc *userUseCase) CreateUser(ctx context.Context, tx *models.Users) (*models.Users, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(tx.Password), bcrypt.DefaultCost) // Perbaikan ada di sini
	if err != nil {
		return nil, fmt.Errorf("failed to hash password: %v", err)
	}

	tx.CreatedAt = time.Now()
	tx.Password = string(hashedPassword)
	return uc.repo.Insert(ctx, tx)
}

func (uc *userUseCase) GetUserById(ctx context.Context, id int) (*models.Users, error) {
	return uc.repo.GetUserByID(ctx, id)
}

func (uc *userUseCase) Login(ctx context.Context, username, password string) (string, string, error) {
	user, err := uc.repo.GetUserByUsername(ctx, username)
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

	tokenString, err := token.SignedString([]byte(os.Getenv("JWT_SECRET_KEY")))
	if err != nil {
		return "", "", fmt.Errorf("failed to sign token: %v", err)
	}

	return tokenString, expirationString, nil
}
