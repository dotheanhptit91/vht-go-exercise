package restaurantcontroller

import (
	"net/http"
	restaurantdtos "vht-go/modules/restaurant/dtos"
	restaurantservice "vht-go/modules/restaurant/service"

	"github.com/gin-gonic/gin"
)

func (ctrl *HTTPRestaurantController) ListRestaurantsAPI() gin.HandlerFunc {
	return func(c *gin.Context) {
		var dto restaurantdtos.ListRestaurantDTO

		if err := c.ShouldBindQuery(&dto); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
			return
		}

		result, err := ctrl.listHandler.Handle(c.Request.Context(), &restaurantservice.ListRestaurantQuery{DTO: dto})
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"data": result})
	}
}

