package handler

import (
	"github.com/Aleksandr-qefy/links-api/internal/service"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	auth := router.Group("/auth")
	{
		auth.POST("sign-up", h.signUp)
		auth.POST("sign-in", h.signIn)
	}

	api := router.Group("/api")
	{
		links := api.Group("/links")
		{
			links.GET("/all", h.linksList) // get all

			links.PUT("/", h.createLink)       // create link
			links.GET("/:id", h.readLink)      // read link
			links.POST("/:id", h.updateLink)   // update link
			links.DELETE("/:id", h.deleteLink) // delete link
		}

		statistics := api.Group("/statistics")
		{
			statistics.GET("/all/:userid", h.statisticList) // get all

			statistics.PUT("/", h.createStatistic)       // create
			statistics.GET("/:id", h.readStatistic)      // read
			statistics.POST("/:id", h.updateStatistic)   // update
			statistics.DELETE("/:id", h.deleteStatistic) // delete
		}
	}

	return router
}
