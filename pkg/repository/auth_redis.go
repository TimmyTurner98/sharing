package repository

import (
	"context"
	"time"

	"github.com/redis/go-redis/v9"
)

type AuthRedis struct {
	rdb *redis.Client
}

func NewAuthRedis(rdb *redis.Client) *AuthRedis {
	return &AuthRedis{rdb: rdb}
}

func (r *AuthRedis) SaveCode(number string, code string) error {
	return r.rdb.Set(context.Background(), number, code, 2*time.Minute).Err()
}

func (r *AuthRedis) GetCode(number string) (string, error) {
	return r.rdb.Get(context.Background(), number).Result()
}

func (r *AuthRedis) DeleteCode(number string) error {
	return r.rdb.Del(context.Background(), number).Err()
}

func (r *AuthRedis) SaveRefreshToken(phone string, token string) error {
	key := "refresh:" + phone
	return r.rdb.Set(context.Background(), key, token, 30*24*time.Hour).Err()
}

func (r *AuthRedis) GetRefreshToken(phone string) (string, error) {
	key := "refresh:" + phone
	return r.rdb.Get(context.Background(), key).Result()
}

func (r *AuthRedis) DeleteRefreshToken(phone string) error {
	key := "refresh:" + phone
	return r.rdb.Del(context.Background(), key).Err()
}
