package service

import "github.com/TimmyTurner98/sharing/pkg/repository"

type Service struct {
	Auth *AuthService
	
}

func NewService(repos *repository.Repository, repos *repos) *Service {
	return &Service{
		Auth: NewAuthService(repos.Auth, repos.Redis),
	}
}
