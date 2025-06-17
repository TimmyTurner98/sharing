package repository

import (
	"github.com/jmoiron/sqlx"
)

type Repository struct {
	Auth *AuthPostgres
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Auth: NewAuthPostgres(db),
	}
}
