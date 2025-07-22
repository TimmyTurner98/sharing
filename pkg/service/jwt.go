package service

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var jwtSecret = []byte(os.Getenv("JWT_SECRET"))

// Генерация access-токена (на 15 минут)
func GenerateAccessToken(phone string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": phone,
		"exp": time.Now().Add(15 * time.Minute).Unix(),
	})
	return token.SignedString(jwtSecret)
}

// Генерация refresh-токена (на 30 дней)
func GenerateRefreshToken(phone string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": phone,
		"exp": time.Now().Add(30 * 24 * time.Hour).Unix(),
	})
	return token.SignedString(jwtSecret)
}
