package psql

import (
	"fmt"

	"database/sql"

	_ "github.com/lib/pq"
)

type Config struct {
	Username string
	Password string
	Host     string
	Port     string
	DBName   string
}

func NewPostgresDB(cfg Config) (*sql.DB, error) {
	psqlInfo := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		cfg.Host, cfg.Port, cfg.Username, cfg.Password, cfg.DBName)

	db, err := sql.Open("postgres", psqlInfo)

	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}
