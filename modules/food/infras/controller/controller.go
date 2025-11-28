package foodcontroller

import (
	fooddomain "vht-go/modules/food/domain"
	foodservice "vht-go/modules/food/service"
	"vht-go/shared"

	"github.com/gin-gonic/gin"
)

type HTTPFoodController struct {
	createHandler shared.ICommandResultHandler[*foodservice.CreateFoodResultCommand, *int]
	getHandler    shared.IQueryHandler[*foodservice.GetFoodQuery, *fooddomain.Food]
	listHandler   shared.IQueryHandler[*foodservice.ListFoodQuery, []fooddomain.Food]
	updateHandler shared.ICommandHandler[*foodservice.UpdateFoodCommand]
	deleteHandler shared.ICommandHandler[*foodservice.DeleteFoodCommand]
}

func NewHTTPFoodController(
	createHandler shared.ICommandResultHandler[*foodservice.CreateFoodResultCommand, *int],
	getHandler shared.IQueryHandler[*foodservice.GetFoodQuery, *fooddomain.Food],
	listHandler shared.IQueryHandler[*foodservice.ListFoodQuery, []fooddomain.Food],
	updateHandler shared.ICommandHandler[*foodservice.UpdateFoodCommand],
	deleteHandler shared.ICommandHandler[*foodservice.DeleteFoodCommand]) *HTTPFoodController {
	return &HTTPFoodController{
		createHandler: createHandler,
		getHandler:    getHandler,
		listHandler:   listHandler,
		updateHandler: updateHandler,
		deleteHandler: deleteHandler,
	}
}

func (ctrl *HTTPFoodController) SetupRoutes(v1 *gin.RouterGroup) {
	foodGroup := v1.Group("/foods")
	{
		foodGroup.POST("", ctrl.CreateFoodAPI())
		foodGroup.GET("/:id", ctrl.GetFoodByIdAPI())
		foodGroup.GET("", ctrl.ListFoodAPI())
		foodGroup.PATCH("/:id", ctrl.UpdateFoodAPI())
		foodGroup.DELETE("/:id", ctrl.DeleteFoodAPI())
	}
}

