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

// @Summary Sign Up
// @Description Create account
// @Tags auth
// @Accept json
// @Produce json
// @Param input body model.UserAccount true "Create account"
// @Success 200 {string} XXXXXXXX-XXXX-XXXX-XXXX-XXXXXXXXXXXX
// @Failure 400 {object} Error
// @Router /auth/sign-up [post]
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
			),
		)
		return
	}

	h.services.Statistic.Create(servModel.Statistic{
		UserId:    id,
		CreatedAt: time.Now(),
		Activity:  signUpTag,
	})

	c.JSON(http.StatusOK, IDResponse{
		ID: id,
	})
}

// @Summary Sign In
// @Description Log in
// @Tags auth
// @Accept json
// @Produce json
// @Param input body model.UserAccount true "Log in"
// @Success 200 {string} string "<jwt token>"
// @Failure 400 {object} Error
// @Router /auth/sign-in [post]
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
			fmt.Sprintf("Incorrect name or password (%s)", err.Error()),
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

// @Summary Sign In
// @Description Delete account
// @Tags auth
// @Security ApiKeyAuth
// @Produce json
// @Success 200 {string} ok
// @Failure 400 {object} Error
// @Router /auth/delete [get]
func (h *Handler) deleteAccount(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}

	if err := h.services.DeleteAccount(userId); err != nil {
		newErrorResponse(
			c,
			http.StatusBadRequest,
			err.Error(),
		)
	}

	c.JSON(http.StatusOK, StatusResponse{
		Status: "ok",
	})
}
