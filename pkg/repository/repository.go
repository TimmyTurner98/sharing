package repository

import (
	"github.com/jmoiron/sqlx"
	"github.com/redis/go-redis/v9"
)

type Repository struct {
	Auth  *AuthPostgres
	Redis *AuthRedis
}

func NewRepository(db *sqlx.DB, rdb *redis.Client) *Repository {
	return &Repository{
		Auth:  NewAuthPostgres(db),
		Redis: NewAuthRedis(rdb),
	}
}
