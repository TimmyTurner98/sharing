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
	key := "code:" + number
	return r.rdb.Set(context.Background(), key, code, 2*time.Minute).Err()
}

func (r *AuthRedis) GetCode(user_id string) (string, error) {
	key := "code:" + user_id
	return r.rdb.Get(context.Background(), key).Result()
}

func (r *AuthRedis) DeleteCode(user_id string) error {
	key := "code:" + user_id
	return r.rdb.Del(context.Background(), key).Err()
}

func (r *AuthRedis) SaveRefreshToken(user_id int, token string) error {
	key := "refresh:" + user_id
	return r.rdb.Set(context.Background(), key, token, 30*24*time.Hour).Err()
}

func (r *AuthRedis) GetRefreshToken(user_id string) (string, error) {
	key := "refresh:" + user_id
	return r.rdb.Get(context.Background(), key).Result()
}

func (r *AuthRedis) DeleteRefreshToken(user_id string) error {
	key := "refresh:" + user_id
	return r.rdb.Del(context.Background(), key).Err()
}
