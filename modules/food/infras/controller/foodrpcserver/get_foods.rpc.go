package foodrpcserver

import (
	fooddomain "vht-go/modules/food/domain"

	"github.com/gin-gonic/gin"
)

type GetFoodsRPCRequest struct {
	Ids []int `json:"ids"`
}

func (s *FoodRPCServer) GetFoodsRPCAPI() gin.HandlerFunc {
	return func(c *gin.Context) {
		var req GetFoodsRPCRequest

		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}

		var foods []fooddomain.Food

		if err := s.db.Where("id IN (?)", req.Ids).Find(&foods).Error; err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}

		c.JSON(200, gin.H{"data": foods})
	}

}

