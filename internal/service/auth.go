package service

import (
	"crypto/sha1"
	"fmt"
	api "github.com/Aleksandr-qefy/links-api"
	"github.com/Aleksandr-qefy/links-api/internal/repository"
	"os"
)

type AuthService struct {
	repo repository.Authorization
}

func NewAuthService(repo repository.Authorization) *AuthService {
	return &AuthService{repo: repo}
}

func (s *AuthService) CreateUser(user api.User) (api.UUID, error) {
	user.Password = s.generatePasswordHash(user.Password)
	return s.repo.CreateUser(user)
}

func (s *AuthService) generatePasswordHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))

	return fmt.Sprintf("%x", hash.Sum([]byte(os.Getenv("CRYPTO_SALT"))))
}
