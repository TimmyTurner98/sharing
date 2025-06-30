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
	var email interface{}
	if user.Email == "" {
		email = nil // ← это важно
	} else {
		email = user.Email
	}

	var id int
	err := r.db.QueryRow(`INSERT INTO users (username, password, number, email) VALUES ($1, $2, $3, $4) RETURNING id`, user.Username, user.Password, user.Number, email).Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, nil
}
