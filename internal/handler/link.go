package handler

import (
	"github.com/Aleksandr-qefy/links-api/internal/handler/model"
	servModel "github.com/Aleksandr-qefy/links-api/internal/service/model"
	"github.com/Aleksandr-qefy/links-api/internal/uuid"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

const (
	allLinksTag   = "all_links"
	createLinkTag = "create_link"
	getLinkTag    = "get_link"
	updateLinkTag = "update_link"
	deleteLinkTag = "delete_link"
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
		categories := make([]model.Category, len(servLink.Categories))
		for j, servCategory := range servLink.Categories {
			categories[j] = model.Category{
				Id:   servCategory.Id,
				Name: servCategory.Name,
			}
		}

		links[i] = model.Link{
			Id:          servLink.Id,
			Ref:         servLink.Ref,
			Description: servLink.Description,
			Categories:  categories,
		}
	}

	h.services.Statistic.Create(servModel.Statistic{
		UserId:    userId,
		CreatedAt: time.Now(),
		Activity:  allLinksTag,
	})

	c.JSON(http.StatusCreated, map[string]interface{}{
		"data": links,
	})
}

func (h *Handler) createLink(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}

	var inputLink model.LinkCreate
	if err := c.BindJSON(&inputLink); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	inputLink.UserId = userId

	id, err := h.services.Link.Create(servModel.LinkUpdate{
		UserId:      inputLink.UserId,
		Ref:         inputLink.Ref,
		Description: inputLink.Description,
		Categories:  inputLink.Categories,
	})
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	comment := string(id)
	h.services.Statistic.Create(servModel.Statistic{
		UserId:    userId,
		CreatedAt: time.Now(),
		Activity:  createLinkTag,
		Comment:   &comment,
	})

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

	categories := make([]model.Category, len(link.Categories))
	for j, servCategory := range link.Categories {
		categories[j] = model.Category{
			Id:   servCategory.Id,
			Name: servCategory.Name,
		}
	}

	comment := string(linkId)
	h.services.Statistic.Create(servModel.Statistic{
		UserId:    userId,
		CreatedAt: time.Now(),
		Activity:  getLinkTag,
		Comment:   &comment,
	})

	c.JSON(http.StatusCreated, model.Link{
		Id:          link.Id,
		Ref:         link.Ref,
		Description: link.Description,
		Categories:  categories,
	})
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

	err = h.services.Link.Update(servModel.LinkUpdate{
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

	comment := string(linkUpdate.Id)
	h.services.Statistic.Create(servModel.Statistic{
		UserId:    userId,
		CreatedAt: time.Now(),
		Activity:  updateLinkTag,
		Comment:   &comment,
	})

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

	comment := string(linkId)
	h.services.Statistic.Create(servModel.Statistic{
		UserId:    userId,
		CreatedAt: time.Now(),
		Activity:  deleteLinkTag,
		Comment:   &comment,
	})

	c.JSON(http.StatusOK, StatusResponse{
		Status: "ok",
	})
}
