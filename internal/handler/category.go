package handler

import (
	"github.com/Aleksandr-qefy/links-api/internal/handler/model"
	servModel "github.com/Aleksandr-qefy/links-api/internal/service/model"
	"github.com/Aleksandr-qefy/links-api/internal/uuid"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) categoriesList(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}

	servCategories, err := h.services.Category.GetAll(userId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	categories := make([]model.Category, len(servCategories))
	for i, servCategory := range servCategories {
		categories[i] = model.Category{
			Id:     servCategory.Id,
			UserId: servCategory.UserId,
			Name:   servCategory.Name,
		}
	}

	c.JSON(http.StatusCreated, map[string]interface{}{
		"data": categories,
	})
}

func (h *Handler) createCategory(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}

	var inputCategory model.Category
	if err := c.BindJSON(&inputCategory); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.services.Category.Create(servModel.Category{
		UserId: userId,
		Name:   inputCategory.Name,
	})
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusCreated, map[string]interface{}{
		"id": id,
	})
}

func (h *Handler) getCategoryById(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}

	categoryIdStr := c.Param("id")
	if categoryIdStr == "" {
		newErrorResponse(c, http.StatusBadRequest, "invalid category id param")
		return
	}

	categoryId := uuid.UUID(categoryIdStr)

	link, err := h.services.Category.GetById(userId, categoryId)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "incorrect category id (or not accessible for this user)")
		return
	}

	c.JSON(http.StatusCreated, link)
}

func (h *Handler) updateCategory(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}

	var categoryUpdate model.CategoryUpdate
	if err := c.BindJSON(&categoryUpdate); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	err = h.services.Category.Update(servModel.Category{
		UserId: userId,
		Id:     categoryUpdate.Id,
		Name:   categoryUpdate.Name,
	})
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, StatusResponse{
		Status: "ok",
	})
}

func (h *Handler) deleteCategoryById(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}

	categoryIdStr := c.Param("id")
	if categoryIdStr == "" {
		newErrorResponse(c, http.StatusBadRequest, "incorrect category id (or not accessible for this user)")
		return
	}

	categoryId := uuid.UUID(categoryIdStr)

	err = h.services.Category.DeleteById(userId, categoryId)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, StatusResponse{
		Status: "ok",
	})
}
