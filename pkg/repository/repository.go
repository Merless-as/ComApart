package repository

import (
	ComApart "HelpWithComm"
	"github.com/jmoiron/sqlx"
)

type Count interface {
	SearchCounter(id int) (ComApart.SearchCounter, error)
	SearchPrice(id int) (ComApart.SearchPrice, error)
}

type Repository struct {
	Count
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		NewCountPostgres(db),
	}
}
