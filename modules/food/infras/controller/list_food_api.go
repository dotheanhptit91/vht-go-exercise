package foodcontroller

import (
	"net/http"
	"strconv"

	fooddtos "vht-go/modules/food/dtos"
	foodservice "vht-go/modules/food/service"
	"vht-go/shared"

	"github.com/gin-gonic/gin"
)

func (ctrl *HTTPFoodController) ListFoodAPI() gin.HandlerFunc {
	return func(c *gin.Context) {
		var dto fooddtos.ListFoodDTO

		if err := c.ShouldBindQuery(&dto); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
			return
		}

		// Parse pagination parameters
		page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
		limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))

		paging := shared.Paging{
			Page:  page,
			Limit: limit,
		}

		query := &foodservice.ListFoodQuery{
			RestaurantId: dto.RestaurantId,
			CategoryId:   dto.CategoryId,
			Status:       dto.Status,
			Paging:       &paging,
		}

		foods, err := ctrl.listHandler.Handle(c.Request.Context(), query)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"data":   foods,
			"paging": paging,
		})
	}
}

