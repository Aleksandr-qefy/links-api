package service

import (
	"github.com/Aleksandr-qefy/links-api/internal/repository"
	model "github.com/Aleksandr-qefy/links-api/internal/service/model"
	"github.com/Aleksandr-qefy/links-api/internal/uuid"
)

type Authorization interface {
	CreateUser(model.UserAccount) (uuid.UUID, error)
	GenerateToken(model.UserAccount) (string, error)
	GetUser(model.UserAccount) (model.User, error)
	ParseToken(token string) (uuid.UUID, error)
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
