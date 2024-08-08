package links_api

import (
	"github.com/Aleksandr-qefy/links-api/internal/uuid"
	"github.com/dgrijalva/jwt-go"
)

type TokenClaims struct {
	jwt.StandardClaims
	UserId uuid.UUID `json:"userIp"`
}
