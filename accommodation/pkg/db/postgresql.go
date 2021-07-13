package db

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
)

type Config struct {
	Host     string
	Port     string
	Username string
	Password string
	DBName   string
	SSLMode  string
}

func NewPostgresDB(config Config) (*sql.DB, error) {
	connectionStr := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
		config.Host, config.Port, config.Username, config.DBName, config.Password, config.SSLMode)

	db, err := sql.Open("postgres", connectionStr)
	if err != nil {
		return nil, err
	}

	log.Println(" Pgsql connected")
	return db, nil
}
