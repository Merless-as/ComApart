package repository

import "database/sql"

const Conn = "user=? password=? dbname=? sslmode=?"

func NewPostgresDB() (*sql.DB, error) {
	db, err := sql.Open("postgres", Conn)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}
