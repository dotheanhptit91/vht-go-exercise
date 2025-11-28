package restaurantrpcserver

import (
	restaurantdomain "vht-go/modules/restaurant/domain"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type GetRestaurantRPCRequest struct {
	Id int `json:"id" binding:"required"`
}

func (s *RestaurantRPCServer) GetRestaurantRPCAPI() gin.HandlerFunc {
	return func(c *gin.Context) {
		var req GetRestaurantRPCRequest

		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}

		var restaurant restaurantdomain.Restaurant

		if err := s.db.First(&restaurant, "id = ?", req.Id).Error; err != nil {
			if err == gorm.ErrRecordNotFound {
				c.JSON(404, gin.H{"error": restaurantdomain.ErrRestaurantNotFound})
				return
			}
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}

		c.JSON(200, gin.H{"data": &restaurant})
	}
}

