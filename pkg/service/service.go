package service

import "github.com/TimmyTurner98/sharing/pkg/repository"

type Service struct{}

func NewService(repos *repository.Repository) *Service {
	return &Service{}
}
