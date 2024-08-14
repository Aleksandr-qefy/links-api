package service

import (
	"github.com/Aleksandr-qefy/links-api/internal/repository"
	model "github.com/Aleksandr-qefy/links-api/internal/service/model"
	"github.com/Aleksandr-qefy/links-api/internal/uuid"
)

type Authorization interface {
	CreateUser(userAcc model.UserAccount) (uuid.UUID, error)
	GenerateToken(userAcc model.UserAccount) (string, error)
	GetUser(userAcc model.UserAccount) (model.User, error)
	ParseToken(bearerToken string) (uuid.UUID, error)
}

type Link interface {
	Create(link model.Link) (uuid.UUID, error)
	GetAll(userId uuid.UUID) ([]model.Link, error)
	GetById(userId, linkId uuid.UUID) (model.Link, error)
	DeleteById(userId, linkId uuid.UUID) error
	Update(model.Link) error
}

type Service struct {
	Authorization
	Link
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
		Link:          NewLinkService(repos.Link),
	}
}
