package handler

import (
	"github.com/Aleksandr-qefy/links-api/internal/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, map[string]interface{}{
			"ping": "pong",
		})
	})

	auth := router.Group("/auth")
	{
		auth.POST("sign-up", h.signUp)
		auth.POST("sign-in", h.signIn)
	}

	api := router.Group("/api", h.userIdentity)
	{
		statistics := api.Group("/statistics")
		{
			statistics.GET("/all/:userid", h.statisticList) // get all

			statistics.DELETE("/:id", h.deleteStatistic) // delete
		}

		categories := api.Group("/categories")
		{
			categories.GET("/all", h.categoriesList) // get all

			categories.PUT("/", h.createCategory)       // create category
			categories.GET("/:id", h.readCategory)      // read category
			categories.POST("/", h.updateCategory)      // update category
			categories.DELETE("/:id", h.deleteCategory) // delete category
		}

		links := api.Group("/links")
		{
			links.GET("/all", h.linksList) // get all

			links.PUT("/", h.createLink)           // create link
			links.GET("/:id", h.getLinkById)       // read link
			links.POST("/:id", h.updateLinkById)   // update link
			links.DELETE("/:id", h.deleteLinkById) // delete link

			links.POST("/add-to-category", h.addLinkToCategory)
			links.POST("/remove-from-category", h.removeLinkFromCategory)
		}
	}

	return router
}
