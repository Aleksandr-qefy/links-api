package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) linksList(c *gin.Context) {
	id, ok := c.Get(userCtx)
	if !ok {
		newErrorResponse(c, http.StatusUnauthorized, "user is't authorized")
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

func (h *Handler) createLink(c *gin.Context) {

}

func (h *Handler) readLink(c *gin.Context) {

}

func (h *Handler) updateLink(c *gin.Context) {

}

func (h *Handler) deleteLink(c *gin.Context) {

}
