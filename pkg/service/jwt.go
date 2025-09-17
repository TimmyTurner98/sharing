package service

import (
	"os"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var jwtSecret = []byte(os.Getenv("JWT_SECRET"))

type tokenClaims struct {
	jwt.RegisteredClaims
	UserId int `json:"user_id"`
}

// Генерация access-токена (на 15 минут)
func GenerateAccessToken(UserId int) (string, error) {
	claims := tokenClaims{
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(15 * time.Minute)),
			Subject:   strconv.Itoa(UserId),
		},
		UserId: UserId,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtSecret)
}

// Генерация refresh-токена (на 30 дней)
func GenerateRefreshToken(UserId int) (string, error) {
	claims := tokenClaims{
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(30 * 24 * time.Hour)),
			Subject:   strconv.Itoa(UserId),
		},
		UserId: UserId,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtSecret)
}

//func GenerateRefreshToken(phone string) (string, error) {
//	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
//	"sub": phone,
//	"exp": time.Now().Add(30 * 24 * time.Hour).Unix(),
//})
//return token.SignedString(jwtSecret)
