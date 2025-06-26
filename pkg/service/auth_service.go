package service

import (
	"github.com/TimmyTurner98/sharing/models"
	"github.com/TimmyTurner98/sharing/pkg/repository"
)

type AuthService struct {
	repo *repository.AuthPostgres
}

func NewAuthService(repo *repository.AuthPostgres) *AuthService {
	return &AuthService{repo: repo}
}

func (s *AuthService) CreateUser(user models.UserRegister) (int, error) {
	return s.repo.CreateUser(user)
}

