package handler

import (
	"github.com/Aleksandr-qefy/links-api/internal/uuid"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type Error struct {
	Message string `json:"message"`
}

type StatusResponse struct {
	Status string `json:"status"`
}

type IDResponse struct {
	ID uuid.UUID `json:"id"`
}

func newErrorResponse(c *gin.Context, statusCode int, message string) {
	logrus.Errorf(message)
	c.AbortWithStatusJSON(statusCode, Error{Message: message})
}
