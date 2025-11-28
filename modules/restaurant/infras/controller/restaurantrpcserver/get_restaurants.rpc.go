package restaurantrpcserver

import (
	restaurantdomain "vht-go/modules/restaurant/domain"

	"github.com/gin-gonic/gin"
)

type GetRestaurantsRPCRequest struct {
	Ids []int `json:"ids"`
}

func (s *RestaurantRPCServer) GetRestaurantsRPCAPI() gin.HandlerFunc {
	return func(c *gin.Context) {
		var req GetRestaurantsRPCRequest

		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}

		var restaurants []restaurantdomain.Restaurant

		if err := s.db.Where("id IN (?)", req.Ids).Find(&restaurants).Error; err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}

		c.JSON(200, gin.H{"data": restaurants})
	}

}

