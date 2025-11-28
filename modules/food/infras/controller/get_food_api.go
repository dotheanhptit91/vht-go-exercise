package foodcontroller

import (
	"net/http"
	"strconv"

	foodservice "vht-go/modules/food/service"

	"github.com/gin-gonic/gin"
)

func (ctrl *HTTPFoodController) GetFoodByIdAPI() gin.HandlerFunc {
	return func(c *gin.Context) {
		idStr := c.Param("id")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": "invalid food id"})
			return
		}

		food, err := ctrl.getHandler.Handle(c.Request.Context(), &foodservice.GetFoodQuery{Id: id})
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"data": food})
	}
}

