package repository

import (
	"github.com/jmoiron/sqlx"
)

const Conn = "user=? password=? dbname=? sslmode=?"

func NewPostgresDB() (*sqlx.DB, error) {
	db, err := sqlx.Open("postgres", Conn)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}
