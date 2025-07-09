package service

import (
	"errors"
	"regexp"

	"github.com/TimmyTurner98/sharing/models"
	"github.com/TimmyTurner98/sharing/pkg/repository"
	"github.com/redis/go-redis/v9"
)

type AuthService struct {
	repo  *repository.AuthPostgres
	redis *redis.Client
}

func NewAuthService(repo *repository.AuthPostgres, reporedis *repository.RedisConfig) *AuthService {
	return &AuthService{repo: repo, reporedis: reporedis}
}

var ErrInvalidNumber = errors.New("invalid phone number format")

func (s *AuthService) CreateUser(user models.UserRegister) (int, error) {
	if !isValidKZNumber(user.Number) {
		return 0, ErrInvalidNumber
	}
	return s.repo.CreateUser(user)
}

func isValidKZNumber(kzNumber string) bool {
	r := regexp.MustCompile(`^\+7(7\d{9})$`)
	return r.MatchString(kzNumber)
}
