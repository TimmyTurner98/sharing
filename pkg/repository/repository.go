package repository

import (
	"github.com/jmoiron/sqlx"
)

type Authorization interface {
}

type Repository struct {
	Auth Authorization
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Auth: NewAuthPostgres(db),
	}
}
