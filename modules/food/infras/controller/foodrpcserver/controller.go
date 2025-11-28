package foodrpcserver

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type FoodRPCServer struct {
	db *gorm.DB
}

func NewFoodRPCServer(db *gorm.DB) *FoodRPCServer {
	return &FoodRPCServer{db: db}
}

func (s *FoodRPCServer) SetupRouter(v1 *gin.RouterGroup) {
	foodRPCGroup := v1.Group("/rpc/foods")
	foodRPCGroup.POST("/get-food", s.GetFoodRPCAPI())
	foodRPCGroup.POST("/get-foods", s.GetFoodsRPCAPI())
}

