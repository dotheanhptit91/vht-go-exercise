package foodcontroller

import (
	"net/http"

	fooddtos "vht-go/modules/food/dtos"
	foodservice "vht-go/modules/food/service"

	"github.com/gin-gonic/gin"
)

func (ctrl *HTTPFoodController) CreateFoodAPI() gin.HandlerFunc {
	return func(c *gin.Context) {
		var dto fooddtos.CreateFoodDTO

		if err := c.ShouldBindJSON(&dto); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
			return
		}

		newId, err := ctrl.createHandler.Handle(c.Request.Context(), &foodservice.CreateFoodResultCommand{DTO: &dto})
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
			return
		}

		c.JSON(http.StatusCreated, gin.H{"data": newId})
	}
}

