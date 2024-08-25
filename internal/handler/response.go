package handler

import (
	"github.com/Aleksandr-qefy/links-api/internal/uuid"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type Error struct {
	Message string `json:"message" example:"some error message"`
}

type StatusResponse struct {
	Status string `json:"status" example:"ok"`
}

type IDResponse struct {
	ID uuid.UUID `json:"id" example:"ffffffff-ffff-ffff-ffff-ffffffffffff"`
}

type TokenResponse struct {
	Token string `json:"token" example:"<jwt token>"`
}

func newErrorResponse(c *gin.Context, statusCode int, message string) {
	logrus.Errorf(message)
	c.AbortWithStatusJSON(statusCode, Error{Message: message})
}
