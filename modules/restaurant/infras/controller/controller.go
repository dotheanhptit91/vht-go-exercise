package restaurantcontroller

import (
	restaurantdomain "vht-go/modules/restaurant/domain"
	restaurantservice "vht-go/modules/restaurant/service"
	"vht-go/shared"

	"github.com/gin-gonic/gin"
)

type HTTPRestaurantController struct {
	createHandler shared.ICommandResultHandler[*restaurantservice.CreateRestaurantResultCommand, int]
	getHandler    shared.IQueryHandler[*restaurantservice.GetRestaurantQuery, *restaurantdomain.Restaurant]
	listHandler   shared.IQueryHandler[*restaurantservice.ListRestaurantQuery, *restaurantservice.ListRestaurantResult]
	updateHandler shared.ICommandHandler[*restaurantservice.UpdateRestaurantCommand]
	deleteHandler shared.ICommandHandler[*restaurantservice.DeleteRestaurantCommand]
}

func NewHTTPRestaurantController(
	createHandler shared.ICommandResultHandler[*restaurantservice.CreateRestaurantResultCommand, int],
	getHandler shared.IQueryHandler[*restaurantservice.GetRestaurantQuery, *restaurantdomain.Restaurant],
	listHandler shared.IQueryHandler[*restaurantservice.ListRestaurantQuery, *restaurantservice.ListRestaurantResult],
	updateHandler shared.ICommandHandler[*restaurantservice.UpdateRestaurantCommand],
	deleteHandler shared.ICommandHandler[*restaurantservice.DeleteRestaurantCommand]) *HTTPRestaurantController {
	return &HTTPRestaurantController{
		createHandler: createHandler,
		getHandler:    getHandler,
		listHandler:   listHandler,
		updateHandler: updateHandler,
		deleteHandler: deleteHandler,
	}
}

func (ctrl *HTTPRestaurantController) SetupRoutes(v1 *gin.RouterGroup) {
	restaurantGroup := v1.Group("/restaurants")
	restaurantGroup.POST("", ctrl.CreateRestaurantAPI())
	restaurantGroup.GET("/:id", ctrl.GetRestaurantByIdAPI())
	restaurantGroup.PATCH("/:id", ctrl.UpdateRestaurantAPI())
	restaurantGroup.GET("", ctrl.ListRestaurantsAPI())
	restaurantGroup.DELETE("/:id", ctrl.DeleteRestaurantAPI())
}

