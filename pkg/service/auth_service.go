package service

import (
	"database/sql"
	"errors"
	"fmt"
	"regexp"

	"github.com/TimmyTurner98/sharing/models"
	"github.com/TimmyTurner98/sharing/pkg/repository"
)

type AuthService struct {
	repo  *repository.AuthPostgres
	redis *repository.AuthRedis
}

func NewAuthService(repo *repository.AuthPostgres, redis *repository.AuthRedis) *AuthService {
	return &AuthService{repo: repo, redis: redis}
}

var ErrInvalidNumber = errors.New("invalid phone number format")

func (s *AuthService) SendCode(user models.UserSignUp) error {
	if !isValidKZNumber(user.Number) {
		return ErrInvalidNumber
	}
	err := s.repo.GetUserByNumber(user.Number)
	if errors.Is(err, sql.ErrNoRows) {
		if err := s.repo.CreateUser(user.Number); err != nil {
			return err
		}
	}

	code := generateSMSCode()
	fmt.Println("Generated code:", code)

	return nil

}

func isValidKZNumber(kzNumber string) bool {
	r := regexp.MustCompile(`^\+7(7\d{9})$`)
	return r.MatchString(kzNumber)
}
