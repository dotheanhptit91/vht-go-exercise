package restaurantcontroller

import (
	"net/http"
	restaurantdtos "vht-go/modules/restaurant/dtos"
	restaurantservice "vht-go/modules/restaurant/service"

	"github.com/gin-gonic/gin"
)

func (ctrl *HTTPRestaurantController) CreateRestaurantAPI() gin.HandlerFunc {
	return func(c *gin.Context) {
		var dto restaurantdtos.CreateRestaurantDTO

		if err := c.ShouldBindJSON(&dto); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
			return
		}

		newId, err := ctrl.createHandler.Handle(c.Request.Context(), &restaurantservice.CreateRestaurantResultCommand{DTO: &dto})
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
			return
		}

		c.JSON(http.StatusCreated, gin.H{"data": newId})
	}
}

