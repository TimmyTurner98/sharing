package service

import (
	"errors"

	"github.com/TimmyTurner98/sharing/models"
	"github.com/golang-jwt/jwt/v5"
)

func (s *AuthService) ParseToken(Token string) (int, error) {
	token, err := jwt.ParseWithClaims(Token, &tokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid signing method")
		}

		return []byte(jwtSecret), nil
	})
	if err != nil {
		return 0, err
	}

	claims, ok := token.Claims.(*tokenClaims)
	if !ok {
		return 0, errors.New("token claims are not of type *tokenClaims")
	}
	return claims.UserId, nil
}

func (s *AuthService) RefreshTokens(input models.RefreshInput) (string, string, error) {
	// 1. Парсим refresh токен
	userId, err := s.ParseToken(input.RefreshToken)
	if err != nil {
		return "", "", err
	}

	// 2. Генерируем новые токены
	newAccess, err := GenerateAccessToken(userId)
	if err != nil {
		return "", "", err
	}

	newRefresh, err := GenerateRefreshToken(userId)
	if err != nil {
		return "", "", err
	}

	err = s.redis.SaveRefreshToken(userId, newRefresh)
	if err != nil {
		return "", "", err
	}

	// 3. Возвращаем
	return newAccess, newRefresh, nil

}
