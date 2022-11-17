package repository

import (
	ComApart "HelpWithComm"
	"fmt"
	"github.com/jmoiron/sqlx"
)

type CountPostgres struct {
	db *sqlx.DB
}

func NewCountPostgres(db *sqlx.DB) *CountPostgres {
	return &CountPostgres{db: db}
}

func (r *CountPostgres) SearchCounter(id int) (ComApart.SearchCounter, error) {
	var answer ComApart.SearchCounter
	query := fmt.Sprintf("SELECT * FROM counter WHERE id=%d AND MAX(count);", id)
	if err := r.db.Get(&answer, query); err != nil {
		return ComApart.SearchCounter{}, err
	}
	return answer, nil
}

func (r *CountPostgres) SearchPrice(id int) (ComApart.SearchPrice, error) {
	var answer ComApart.SearchPrice
	query := fmt.Sprintf("SELECT * FROM price WHERE id=%d;", id)
	if err := r.db.Get(&answer, query); err != nil {
		return ComApart.SearchPrice{}, err
	}
	return answer, nil
}
