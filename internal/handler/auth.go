package handler

import (
	"fmt"
	"github.com/Aleksandr-qefy/links-api/internal/handler/model"
	servModel "github.com/Aleksandr-qefy/links-api/internal/service/model"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

const (
	signUpTag = "sign_up"
	signInTag = "sign_in"
)

func (h *Handler) signUp(c *gin.Context) {
	var userAccount model.UserAccount
	if err := c.BindJSON(&userAccount); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.services.Authorization.CreateUser(servModel.UserAccount{
		Name:     userAccount.Name,
		Password: userAccount.Password,
	})
	if err != nil {
		newErrorResponse(
			c,
			http.StatusBadRequest,
			fmt.Sprintf(
				"User with name '%s' exists already (%s)",
				userAccount.Name,
				err.Error(),
			), // err.Error()
		)
		return
	}

	h.services.Statistic.Create(servModel.Statistic{
		UserId:    id,
		CreatedAt: time.Now(),
		Activity:  signUpTag,
	})

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

func (h *Handler) signIn(c *gin.Context) {
	var userAccount model.UserAccount
	if err := c.BindJSON(&userAccount); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	userId, token, err := h.services.Authorization.GenerateToken(
		servModel.UserAccount{
			Name:     userAccount.Name,
			Password: userAccount.Password,
		},
	)
	if err != nil {
		newErrorResponse(
			c,
			http.StatusBadRequest,
			fmt.Sprintf("Incorrect name or password (%s)", err.Error()), // err.Error()
		)
		return
	}

	h.services.Statistic.Create(servModel.Statistic{
		UserId:    userId,
		CreatedAt: time.Now(),
		Activity:  signInTag,
	})

	c.JSON(http.StatusOK, map[string]interface{}{
		"token": token,
	})
}

func (h *Handler) deleteAccount(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}

	if err := h.services.DeleteAccount(userId); err != nil {
		newErrorResponse(
			c,
			http.StatusBadRequest,
			fmt.Sprintf(err.Error()), // err.Error()
		)
	}

	c.JSON(http.StatusOK, StatusResponse{
		Status: "ok",
	})
}
