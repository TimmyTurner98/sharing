package repository

import (
	"context"
	"strconv"
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

func (r *AuthRedis) GetCode(number string) (string, error) {
	key := "code:" + number
	return r.rdb.Get(context.Background(), key).Result()
}

func (r *AuthRedis) DeleteCode(number string) error {
	key := "code:" + number
	return r.rdb.Del(context.Background(), key).Err()
}

func (r *AuthRedis) SaveRefreshToken(user_id int, token string) error {
	key := "refresh:" + strconv.Itoa(user_id)
	return r.rdb.Set(context.Background(), key, token, 30*24*time.Hour).Err()
}

func (r *AuthRedis) GetRefreshToken(user_id int) (string, error) {
	key := "refresh:" + strconv.Itoa(user_id)
	return r.rdb.Get(context.Background(), key).Result()
}

func (r *AuthRedis) DeleteRefreshToken(user_id int) error {
	key := "refresh:" + strconv.Itoa(user_id)
	return r.rdb.Del(context.Background(), key).Err()
}
