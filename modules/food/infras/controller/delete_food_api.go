package foodcontroller

import (
	"net/http"
	"strconv"

	foodservice "vht-go/modules/food/service"

	"github.com/gin-gonic/gin"
)

func (ctrl *HTTPFoodController) DeleteFoodAPI() gin.HandlerFunc {
	return func(c *gin.Context) {
		idStr := c.Param("id")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": "invalid food id"})
			return
		}

		err = ctrl.deleteHandler.Handle(c.Request.Context(), &foodservice.DeleteFoodCommand{Id: id})
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"data": true})
	}
}

