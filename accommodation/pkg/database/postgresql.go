package database

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
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
	connectionStr := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=%s",
		config.Username,
		config.Password,
		config.Host,
		config.Port,
		config.DBName,
		config.SSLMode)

	db, err := sql.Open("postgres", connectionStr)
	if err != nil {
		return nil, err
	}

	return db, nil
}
