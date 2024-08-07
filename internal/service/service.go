package service

import (
	api "github.com/Aleksandr-qefy/links-api"
	"github.com/Aleksandr-qefy/links-api/internal/repository"
)

type Authorization interface {
	CreateUser(user api.User) (api.UUID, error)
}

type Link interface {
}

type Service struct {
	Authorization
	Link
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
	}
}
