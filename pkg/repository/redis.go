package repository

import (
	"os"
	"strconv"

	"github.com/redis/go-redis/v9"
)

type RedisConfig struct {
	Addr     string
	Password string
	DB       int
}

func LoadRedisConfig() RedisConfig {
	dbStr := os.Getenv("REDIS_DB")
	db, err := strconv.Atoi(dbStr)
	if err != nil {
		db = 0
	}
	return RedisConfig{
		Addr:     os.Getenv("REDIS_ADDR"),
		Password: os.Getenv("REDIS_PASSWORD"),
		DB:       db,
	}
}

func NewRedisClient(cfg RedisConfig) *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:     cfg.Addr,
		Password: cfg.Password,
		DB:       cfg.DB,
	})
	return rdb
}
