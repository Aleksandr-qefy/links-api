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
	allCategoriesTag  = "all_categories"
	createCategoryTag = "create_category"
	getCategoryTag    = "get_category"
	updateCategoryTag = "update_category"
	deleteCategoryTag = "delete_category"
)

// @Summary Categories List
// @Description Show categories list
// @Tags categories
// @Security ApiKeyAuth
// @Produce json
// @Success 200 {object} model.AllCategories
// @Failure 400 {object} Error
// @Router /api/categories/all [get]
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
			Id:   servCategory.Id,
			Name: servCategory.Name,
		}
	}

	h.services.Statistic.Create(servModel.Statistic{
		UserId:    userId,
		CreatedAt: time.Now(),
		Activity:  allCategoriesTag,
	})

	c.JSON(http.StatusCreated, model.AllCategories{
		Data: categories,
	})
}

// @Summary Create Category
// @Description Create new category
// @Tags categories
// @Security ApiKeyAuth
// @Accept json
// @Produce json
// @Param input body model.Category true "New category"
// @Success 200 {object} IDResponse
// @Failure 400 {object} Error
// @Router /api/categories [put]
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

	comment := string(id)
	h.services.Statistic.Create(servModel.Statistic{
		UserId:    userId,
		CreatedAt: time.Now(),
		Activity:  createCategoryTag,
		Comment:   &comment,
	})

	c.JSON(http.StatusCreated, IDResponse{
		ID: id,
	})
}

// @Summary Get Category By Id
// @Description Get category by id
// @Tags categories
// @Security ApiKeyAuth
// @Accept json
// @Produce json
// @Success 200 {object} model.Category
// @Failure 400 {object} Error
// @Param id path string true "Category ID"
// @Router /api/categories/{id} [get]
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

	category, err := h.services.Category.GetById(userId, categoryId)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "incorrect category id (or not accessible for this user)")
		return
	}

	comment := string(categoryId)
	h.services.Statistic.Create(servModel.Statistic{
		UserId:    userId,
		CreatedAt: time.Now(),
		Activity:  getCategoryTag,
		Comment:   &comment,
	})

	c.JSON(http.StatusCreated, model.Category{
		Id:   category.Id,
		Name: category.Name,
	})
}

// @Summary Update Category
// @Description Update category
// @Tags categories
// @Security ApiKeyAuth
// @Accept json
// @Produce json
// @Param input body model.Category true "Update category"
// @Success 200 {object} StatusResponse
// @Failure 400 {object} Error
// @Router /api/categories [post]
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

	comment := string(categoryUpdate.Id)
	h.services.Statistic.Create(servModel.Statistic{
		UserId:    userId,
		CreatedAt: time.Now(),
		Activity:  updateCategoryTag,
		Comment:   &comment,
	})

	c.JSON(http.StatusOK, StatusResponse{
		Status: "ok",
	})
}

// @Summary Delete Category By Id
// @Description Delete category by id
// @Tags categories
// @Security ApiKeyAuth
// @Produce json
// @Success 200 {object} StatusResponse
// @Failure 400 {object} Error
// @Param id path string true "Category ID"
// @Router /api/categories/{id} [delete]
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

	comment := string(categoryId)
	h.services.Statistic.Create(servModel.Statistic{
		UserId:    userId,
		CreatedAt: time.Now(),
		Activity:  deleteCategoryTag,
		Comment:   &comment,
	})

	c.JSON(http.StatusOK, StatusResponse{
		Status: "ok",
	})
}
