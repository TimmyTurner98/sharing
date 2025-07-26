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

	_, err := s.repo.GetUserByNumber(user.Number)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			if err := s.repo.CreateUser(user.Number); err != nil {
				return err
			}
		} else {
			return err
		}
	}

	code := generateSMSCode()
	fmt.Println("Generated code:", code)

	return s.redis.SaveCode(user.Number, code)
}

func isValidKZNumber(kzNumber string) bool {
	r := regexp.MustCompile(`^\+7(7\d{9})$`)
	return r.MatchString(kzNumber)
}

func (s *AuthService) VerifyCode(number string, inputCode string) (string, string, error) {
	// 1. Получаем код из Redis
	storedCode, err := s.redis.GetCode(number)
	if err != nil {
		return "", "", err
	}

	// 2. Сравниваем
	if storedCode != inputCode {
		return "", "", errors.New("invalid code")
	}

	// 3. Удаляем одноразовый код
	_ = s.redis.DeleteCode(number)

	// 4. Генерируем access и refresh токены
	userId, err := s.repo.GetUserByNumber(number)
	if err != nil {
		return "", "", err
	}
	accessToken, err := GenerateAccessToken(userId)
	if err != nil {
		return "", "", err
	}

	refreshToken, err := GenerateRefreshToken(userId)
	if err != nil {
		return "", "", err
	}

	// 5. Сохраняем refresh токен в Redis
	err = s.redis.SaveRefreshToken(number, refreshToken)
	if err != nil {
		return "", "", err
	}

	// 6. Возвращаем оба токена
	return accessToken, refreshToken, nil
}
