package service

import (
	"crypto/sha1"
	"errors"
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
	tokenTTL   = 30 * time.Minute
	signingKey = "bkuvvkjuv56df89h2r8h3290102-9012-0e_)()YT&78tg2de7gh12gihi21d"
)

type AuthService struct {
	repo repository.Authorization
}

func NewAuthService(repo repository.Authorization) *AuthService {
	return &AuthService{repo: repo}
}

func (s *AuthService) CreateUser(user model.UserAccount) (uuid.UUID, error) {
	return s.repo.CreateUser(repoModel.User{
		Name:         user.Name,
		PasswordHash: s.generatePasswordHash(user.Password),
	})
}

func (s *AuthService) GetUser(userAccount model.UserAccount) (model.User, error) {

	repoUser, err := s.repo.GetUser(repoModel.User{
		Name:         userAccount.Name,
		PasswordHash: s.generatePasswordHash(userAccount.Password),
	})
	if err != nil {
		return model.User{}, err
	}

	user := model.User{
		Id:           repoUser.Id,
		Name:         repoUser.Name,
		PasswordHash: repoUser.PasswordHash,
	}
	return user, nil
}

func (s *AuthService) GenerateToken(userAccount model.UserAccount) (uuid.UUID, string, error) {
	user, err := s.GetUser(userAccount)
	if err != nil {
		return "", "", err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &model.TokenClaims{
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(tokenTTL).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
		user.Id,
	})

	signedToken, err := token.SignedString([]byte(signingKey))

	return user.Id, signedToken, err
}

func (s *AuthService) ParseToken(accessToken string) (uuid.UUID, error) {
	token, err := jwt.ParseWithClaims(accessToken, &model.TokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, isOk := token.Method.(*jwt.SigningMethodHMAC); !isOk {
			return nil, errors.New("invalid signing method")
		}
		return []byte(signingKey), nil
	})
	if err != nil {
		return "", err
	}

	claims, ok := token.Claims.(*model.TokenClaims)
	if !ok {
		return "", errors.New("token claims are not  of type *model.TokenClaims")
	}

	return claims.UserId, nil
}

func (s *AuthService) generatePasswordHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))

	return fmt.Sprintf("%x", hash.Sum([]byte(os.Getenv("CRYPTO_SALT"))))
}

func (s *AuthService) DeleteAccount(userId uuid.UUID) error {
	return s.repo.DeleteAccount(userId)
}
