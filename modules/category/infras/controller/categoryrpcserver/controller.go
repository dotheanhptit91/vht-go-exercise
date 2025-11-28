package categoryrpcserver

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type CategoryRPCServer struct {
	db *gorm.DB
}

func NewCategoryRPCServer(db *gorm.DB) *CategoryRPCServer {
	return &CategoryRPCServer{db: db}
}

func (s *CategoryRPCServer) SetupRouter(v1 *gin.RouterGroup) {
	catRPCGroup := v1.Group("/rpc/categories")
	catRPCGroup.POST("/get-category", s.GetCategoryRPCAPI())
	catRPCGroup.POST("/get-categories", s.GetCategoriesRPCAPI())
}