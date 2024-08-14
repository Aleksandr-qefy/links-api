package handler

import (
	"github.com/Aleksandr-qefy/links-api/internal/handler/model"
	servModel "github.com/Aleksandr-qefy/links-api/internal/service/model"
	"github.com/Aleksandr-qefy/links-api/internal/uuid"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) linksList(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}

	servLinks, err := h.services.Link.GetAll(userId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	links := make([]model.Link, len(servLinks))
	for i, servLink := range servLinks {
		links[i] = model.Link{
			Id:          servLink.Id,
			UserId:      servLink.UserId,
			Ref:         servLink.Ref,
			Description: servLink.Description,
			Categories:  servLink.Categories,
		}
	}

	c.JSON(http.StatusCreated, map[string]interface{}{
		"data": links,
	})
}

func (h *Handler) createLink(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}

	var inputLink model.Link
	if err := c.BindJSON(&inputLink); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	inputLink.UserId = userId

	id, err := h.services.Link.Create(servModel.Link{
		UserId:      inputLink.UserId,
		Ref:         inputLink.Ref,
		Description: inputLink.Description,
		Categories:  inputLink.Categories,
	})
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusCreated, map[string]interface{}{
		"id": id,
	})
}

func (h *Handler) getLinkById(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}

	linkIdStr := c.Param("id")
	if linkIdStr == "" {
		newErrorResponse(c, http.StatusBadRequest, "invalid link id param")
		return
	}

	linkId := uuid.UUID(linkIdStr)

	link, err := h.services.Link.GetById(userId, linkId)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "incorrect link id (or not accessible for this user)")
		return
	}

	c.JSON(http.StatusCreated, link)
}

func (h *Handler) updateLink(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}

	var linkUpdate model.LinkUpdate
	if err := c.BindJSON(&linkUpdate); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	var categories []uuid.UUID = nil
	if linkUpdate.Categories != nil {
		categories = *linkUpdate.Categories
	}

	err = h.services.Link.Update(servModel.Link{
		UserId:      userId,
		Id:          linkUpdate.Id,
		Ref:         linkUpdate.Ref,
		Description: linkUpdate.Description,
		Categories:  categories,
	})
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, StatusResponse{
		Status: "ok",
	})
}

func (h *Handler) deleteLinkById(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}

	linkIdStr := c.Param("id")
	if linkIdStr == "" {
		newErrorResponse(c, http.StatusBadRequest, "incorrect link id (or not accessible for this user)")
		return
	}

	linkId := uuid.UUID(linkIdStr)

	err = h.services.Link.DeleteById(userId, linkId)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, StatusResponse{
		Status: "ok",
	})
}
