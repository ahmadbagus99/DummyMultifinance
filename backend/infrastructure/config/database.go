package config

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func NewDB() *sql.DB {
	LoadEnv()

	user := GetEnv("DB_USERNAME", "root")
	password := GetEnv("DB_PASSWORD", "")
	host := GetEnv("DB_HOST", "localhost")
	port := GetEnv("DB_PORT", "3306")
	dbname := GetEnv("DB_NAME", "3306")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true",
		user, password, host, port, dbname)

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatalf("Error opening DB: %v", err)
	}

	if err := db.Ping(); err != nil {
		log.Fatalf("Error connecting to DB: %v", err)
	}

	log.Println("Database connected successfully.")
	return db
}
