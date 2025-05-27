package integration

import (
	"DummyMultifinance/domain/models"
	"DummyMultifinance/domain/repositories"
	"DummyMultifinance/infrastructure/config"
	repositoriesMySQL "DummyMultifinance/infrastructure/repositories/users"
	usecases "DummyMultifinance/usecases/users"
	"context"
	"fmt"
	"log"
	"path/filepath"
	"testing"

	"database/sql"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
)

var db *sql.DB
var tx *sql.Tx
var userRepo repositories.UserRepository
var userUseCase usecases.UserUseCase

func setup() {
	filename := fmt.Sprintf(".env.%s", config.GetAppEnv())
	envPath := filepath.Join("..", "..", filename)

	err := godotenv.Load(envPath)
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	user := config.GetEnv("DB_USERNAME", "root")
	password := config.GetEnv("DB_PASSWORD", "")
	host := config.GetEnv("DB_HOST", "localhost")
	port := config.GetEnv("DB_PORT", "3306")
	dbname := config.GetEnv("DB_NAME", "dummy_multifinance_dev")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", user, password, host, port, dbname)

	db, err = sql.Open("mysql", dsn)
	if err != nil {
		log.Fatalf("Error opening DB: %v", err)
	}

	tx, err = db.Begin()
	if err != nil {
		log.Fatalf("Error starting transaction: %v", err)
	}

	userRepo = repositoriesMySQL.NewMysqlUserRepo(tx)
	userUseCase = usecases.NewUserUsecase(userRepo)

	if userRepo == nil || userUseCase == nil {
		log.Fatal("userRepo or userUseCase is nil!")
	}

	_, err = tx.Exec(`
		CREATE TABLE IF NOT EXISTS users (
			id INT AUTO_INCREMENT PRIMARY KEY,
			username VARCHAR(255) NOT NULL,
			password VARCHAR(255) NOT NULL,
			role_id INT NOT NULL,
			email VARCHAR(255),
			created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
			CONSTRAINT fk_role_id FOREIGN KEY (role_id) REFERENCES roles(id)
		)
	`)
	if err != nil {
		panic("failed to migrate database schema: " + err.Error())
	}
}

func teardown() {
	if tx != nil {
		tx.Rollback()
	}
}

func TestCreateUserIntegration(t *testing.T) {
	setup()
	defer teardown()

	user := &models.Users{
		Username: "testuser",
		Password: "password123",
		RoleID:   "1",
		Email:    "test@email.com",
	}

	createdUser, err := userUseCase.CreateUser(context.Background(), user)

	assert.NoError(t, err)
	assert.NotNil(t, createdUser)
	assert.Equal(t, "testuser", createdUser.Username)
	assert.Equal(t, "1", createdUser.RoleID)
	assert.Equal(t, "test@email.com", createdUser.Email)
}
