package handler

import (
	"fmt"
	"github.com/Aleksandr-qefy/links-api/internal/handler/model"
	serviceModel "github.com/Aleksandr-qefy/links-api/internal/service/model"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) signUp(c *gin.Context) {
	var userAccount model.UserAccount
	if err := c.BindJSON(&userAccount); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.services.Authorization.CreateUser(serviceModel.UserAccount{
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

	id, err := h.services.Authorization.GenerateToken(
		serviceModel.UserAccount{
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

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}
