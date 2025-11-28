package foodcontroller

import (
	"net/http"
	"strconv"

	fooddtos "vht-go/modules/food/dtos"
	foodservice "vht-go/modules/food/service"

	"github.com/gin-gonic/gin"
)

func (ctrl *HTTPFoodController) UpdateFoodAPI() gin.HandlerFunc {
	return func(c *gin.Context) {
		idStr := c.Param("id")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": "invalid food id"})
			return
		}

		var dto fooddtos.UpdateFoodDTO

		if err := c.ShouldBindJSON(&dto); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
			return
		}

		err = ctrl.updateHandler.Handle(c.Request.Context(), &foodservice.UpdateFoodCommand{
			Id:  id,
			DTO: &dto,
		})
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"data": true})
	}
}

