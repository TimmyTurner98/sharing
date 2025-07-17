package repository

import (
	"github.com/jmoiron/sqlx"
)

type AuthPostgres struct {
	db *sqlx.DB
}

func NewAuthPostgres(db *sqlx.DB) *AuthPostgres {
	return &AuthPostgres{db: db}
}

func (r *AuthPostgres) GetUserByNumber(number string) error {
	var exists bool
	err := r.db.QueryRow(`SELECT EXISTS (SELECT 1 FROM users WHERE number = $1)`, number).Scan(&exists)
	if err != nil {
		return err
	}
	if !exists {
		return r.CreateUser(number)
	}
	return nil
}

func (r *AuthPostgres) CreateUser(number string) error {
	var id int
	err := r.db.QueryRow(`INSERT INTO users (number) VALUES ($1) RETURNING id`, number).Scan(&id)
	if err != nil {
		return err
	}
	return nil
}

// var id int
// err := r.db.QueryRow(`INSERT INTO users (username, password, number, email) VALUES ($1, $2, $3, $4) RETURNING id`, user.Username, user.Password, user.Number, email).Scan(&id)
// if err != nil {
// 	return 0, err
// }
// return id, nil
