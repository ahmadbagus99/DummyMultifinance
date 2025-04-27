package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func GetAppEnv() string {
	env := os.Getenv("APP_ENV")
	if env == "" {
		env = "development"
	}

	return env
}

func LoadEnv() {
	env := GetAppEnv()

	filename := fmt.Sprintf(".env.%s", env)
	err := godotenv.Load(filename)
	if err != nil {
		log.Fatalf("Gagal load %s: %v", filename, err)
	}

	log.Println("APP_ENV:", os.Getenv("APP_ENV"))
}

func GetEnv(key, fallback string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return fallback
}
