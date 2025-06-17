package repository

import (
	"github.com/TimmyTurner98/sharing/models"
	"github.com/jmoiron/sqlx"
)

type AuthPostgres struct {
	db *sqlx.DB
}

func NewAuthPostgres(db *sqlx.DB) *AuthPostgres {
	return &AuthPostgres{db: db}
}

func (r *AuthPostgres) CreateUser(user models.UserRegister) (int, error) {
	var id int
	err := r.db.QueryRow("INSERT INTO users (username, password, number) VALUES ($1, $2, $3) RETURNING id", user.Username, user.Password, user.Number).Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, nil
}
