package restaurantrpcserver

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type RestaurantRPCServer struct {
	db *gorm.DB
}

func NewRestaurantRPCServer(db *gorm.DB) *RestaurantRPCServer {
	return &RestaurantRPCServer{db: db}
}

func (s *RestaurantRPCServer) SetupRouter(v1 *gin.RouterGroup) {
	rstRPCGroup := v1.Group("/rpc/restaurants")
	rstRPCGroup.POST("/get-restaurant", s.GetRestaurantRPCAPI())
	rstRPCGroup.POST("/get-restaurants", s.GetRestaurantsRPCAPI())
}

