package handler

import (
	"fmt"
	"github.com/Aleksandr-qefy/links-api/internal/handler/model"
	serviceModel "github.com/Aleksandr-qefy/links-api/internal/service/model"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) signUp(c *gin.Context) {
	var inputUser model.UserSign
	if err := c.BindJSON(&inputUser); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.services.Authorization.CreateUser(serviceModel.User{
		Name:     inputUser.Name,
		Password: inputUser.Password,
	})
	if err != nil {
		newErrorResponse(
			c,
			http.StatusInternalServerError,
			fmt.Sprintf(
				"User with name '%s' exists already (%s)",
				inputUser.Name,
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
	var inputUser model.UserSign
	if err := c.BindJSON(&inputUser); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.services.Authorization.GenerateToken(
		inputUser.Name,
		inputUser.Password,
	)
	if err != nil {
		newErrorResponse(
			c,
			http.StatusInternalServerError,
			fmt.Sprintf("Incorrect name or password (%s)", err.Error()), // err.Error()
		)
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}
