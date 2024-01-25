package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"os"
)

func NewPostgresDB() (*sqlx.DB, error) {

	dbinfo := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"))

	db, err := sqlx.Open("postgres", dbinfo)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, err
}
