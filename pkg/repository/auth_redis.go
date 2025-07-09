package repository

import "github.com/redis/go-redis/v9"

type AuthRedis struct {
	rdb *redis.Client
}

func NewAuthRedis(rdb *redis.Client) *AuthRedis {
	return &AuthRedis{rdb: rdb}
}
