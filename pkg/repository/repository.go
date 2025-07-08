package repository

import (
	"github.com/jmoiron/sqlx"
	"github.com/redis/go-redis/v9"
)

type Repository struct {
	Auth  *AuthPostgres
	Redis *redis.Client
}

func NewRepository(db *sqlx.DB, redisClient *redis.Client) *Repository {
	return &Repository{
		Auth:  NewAuthPostgres(db),
		Redis: redisClient,
	}
}
