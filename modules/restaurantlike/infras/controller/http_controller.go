package rstlikecontroller

import (
	"vht-go/middleware"
	rstlikeservice "vht-go/modules/restaurantlike/service"
	"vht-go/shared"

	"github.com/gin-gonic/gin"
)

type HTTPRestaurantLikeController struct {
	likeRestaurantCommandHandler   shared.ICommandHandler[*rstlikeservice.LikeRestaurantCommand]
	unlikeRestaurantCommandHandler shared.ICommandHandler[*rstlikeservice.UnlikeRestaurantCommand]
}

func NewHTTPRestaurantLikeController(
	likeRestaurantCommandHandler shared.ICommandHandler[*rstlikeservice.LikeRestaurantCommand],
	unlikeRestaurantCommandHandler shared.ICommandHandler[*rstlikeservice.UnlikeRestaurantCommand],
) *HTTPRestaurantLikeController {
	return &HTTPRestaurantLikeController{
		likeRestaurantCommandHandler:   likeRestaurantCommandHandler,
		unlikeRestaurantCommandHandler: unlikeRestaurantCommandHandler,
	}
}

func (ctrl *HTTPRestaurantLikeController) SetupRouter(v1 *gin.RouterGroup, middlewareProvider middleware.IMiddlewareProvider) {
	restaurants := v1.Group("/restaurants", middlewareProvider.Auth())
	{
		restaurants.POST("/:id/like", ctrl.UserLikeRestaurantAPI())
		restaurants.POST("/:id/unlike", ctrl.UserUnlikeRestaurantAPI())
	}
}