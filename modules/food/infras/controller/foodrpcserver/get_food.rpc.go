package foodrpcserver

import (
	fooddomain "vht-go/modules/food/domain"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type GetFoodRPCRequest struct {
	Id int `json:"id" binding:"required"`
}

func (s *FoodRPCServer) GetFoodRPCAPI() gin.HandlerFunc {
	return func(c *gin.Context) {
		var req GetFoodRPCRequest

		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}

		var food fooddomain.Food

		if err := s.db.First(&food, "id = ?", req.Id).Error; err != nil {
			if err == gorm.ErrRecordNotFound {
				c.JSON(404, gin.H{"error": fooddomain.ErrFoodNotFound})
				return
			}
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}

		c.JSON(200, gin.H{"data": &food})
	}
}

