package service

import "github.com/Aleksandr-qefy/links-api/internal/repository"

type Authorization interface {
}

type Link interface {
}

type Service struct {
	Authorization
	Link
}

func NewService(repos *repository.Repository) *Service {
	return &Service{}
}
