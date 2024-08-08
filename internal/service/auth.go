package service

import (
	"crypto/sha1"
	"fmt"
	"github.com/Aleksandr-qefy/links-api/internal/repository"
	repoModel "github.com/Aleksandr-qefy/links-api/internal/repository/model"
	model "github.com/Aleksandr-qefy/links-api/internal/service/model"
	"github.com/Aleksandr-qefy/links-api/internal/uuid"
	"github.com/dgrijalva/jwt-go"
	"os"
	"time"
)

const (
	tokenTTL   = 15 * time.Minute
	signingKey = "bkuvvkjuv56df89h2r8h3290102-9012-0e_)()YT&78tg2de7gh12gihi21d"
)

type AuthService struct {
	repo repository.Authorization
}

func NewAuthService(repo repository.Authorization) *AuthService {
	return &AuthService{repo: repo}
}

func (s *AuthService) CreateUser(user model.User) (uuid.UUID, error) {
	return s.repo.CreateUser(repoModel.User{
		Name:         user.Name,
		PasswordHash: s.generatePasswordHash(user.Password),
	})
}

func (s *AuthService) GetUser(name string, password string) (model.User, error) {

	repoUser, err := s.repo.GetUser(repoModel.User{
		Name:         name,
		PasswordHash: s.generatePasswordHash(password),
	})
	if err != nil {
		return model.User{}, err
	}

	user := model.User{
		Id:       repoUser.Id,
		Name:     repoUser.Name,
		Password: repoUser.PasswordHash,
	}
	return user, nil
}

func (s *AuthService) GenerateToken(name string, password string) (string, error) {
	user, err := s.GetUser(name, password)
	if err != nil {
		return "", err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &model.TokenClaims{
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(tokenTTL).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
		user.Id,
	})

	return token.SignedString([]byte(signingKey))
}

func (s *AuthService) generatePasswordHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))

	return fmt.Sprintf("%x", hash.Sum([]byte(os.Getenv("CRYPTO_SALT"))))
}
