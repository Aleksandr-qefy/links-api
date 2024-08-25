package handler

import (
	"github.com/Aleksandr-qefy/links-api/internal/handler/model"
	"github.com/gin-gonic/gin"
	"net/http"
)

// @Summary Statistic List
// @Description Show statistic list
// @Tags statistics
// @Security ApiKeyAuth
// @Produce json
// @Success 200 {object} model.AllStatistics
// @Failure 400 {object} Error
// @Failure 401 {object} Error
// @Router /api/statistics/all [get]
func (h *Handler) statisticList(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}

	servStatistics, err := h.services.Statistic.GetAll(userId)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	statistics := make([]model.Statistic, len(servStatistics))
	for i, servStatistic := range servStatistics {
		statistics[i] = model.Statistic{
			Id:        servStatistic.Id,
			CreatedAt: servStatistic.CreatedAt,
			Activity:  servStatistic.Activity,
			Comment:   servStatistic.Comment,
		}
	}

	c.JSON(http.StatusOK, model.AllStatistics{
		Data: statistics,
	})
}
