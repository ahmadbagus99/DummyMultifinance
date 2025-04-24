package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func LoadEnv() {
	env := os.Getenv("APP_ENV")
	if env == "" {
		env = "development"
	}

	filename := fmt.Sprintf(".env.%s", env)
	err := godotenv.Load(filename)
	if err != nil {
		log.Fatalf("Gagal load %s: %v", filename, err)
	}
}

func GetEnv(key, fallback string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return fallback
}
