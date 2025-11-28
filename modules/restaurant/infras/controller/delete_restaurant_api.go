package restaurantcontroller

import (
	"net/http"
	"strconv"
	restaurantservice "vht-go/modules/restaurant/service"

	"github.com/gin-gonic/gin"
)

func (ctrl *HTTPRestaurantController) DeleteRestaurantAPI() gin.HandlerFunc {
	return func(c *gin.Context) {
		idParam := c.Param("id")
		id, err := strconv.Atoi(idParam)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": "invalid id format"})
			return
		}

		if err := ctrl.deleteHandler.Handle(c.Request.Context(), &restaurantservice.DeleteRestaurantCommand{Id: id}); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"data": true})
	}
}

