package categoryrpcserver

import (
	categorydomain "vht-go/modules/category/domain"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type GetCategoryRPCRequest struct {
	Id uuid.UUID `json:"id" binding:"required"`
}

func (s *CategoryRPCServer) GetCategoryRPCAPI() gin.HandlerFunc {
	return func(c *gin.Context) {
		var req GetCategoryRPCRequest

		if err := c.ShouldBindJSON(&req);  err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}

		var category categorydomain.Category

		if err := s.db.First(&category, "id = ?", req.Id).Error; err != nil {
			if err == gorm.ErrRecordNotFound {
				c.JSON(404, gin.H{"error": categorydomain.ErrCategoryNotFound})
				return
			}
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}

		c.JSON(200, gin.H{"data": &category})
	}
}