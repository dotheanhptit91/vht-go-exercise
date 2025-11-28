package categoryrpcserver

import (
	categorydomain "vht-go/modules/category/domain"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type GetCategoriesRPCRequest struct {
	Ids []uuid.UUID `json:"ids"`
}

func (s *CategoryRPCServer) GetCategoriesRPCAPI() gin.HandlerFunc {
	return func(c *gin.Context) {
		var req GetCategoriesRPCRequest

		if err := c.ShouldBindJSON(&req);  err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}

		var categories []categorydomain.Category

		if err := s.db.Where("id IN (?)", req.Ids).Find(&categories).Error; err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}

		c.JSON(200, gin.H{"data": categories})
	}

}