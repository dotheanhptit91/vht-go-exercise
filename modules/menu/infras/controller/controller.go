package menucontroller

import (
	menudomain "vht-go/modules/menu/domain"
	menuservice "vht-go/modules/menu/service"
	"vht-go/shared"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type HTTPMenuController struct {
	createHandler shared.ICommandResultHandler[*menuservice.CreateMenuResultCommand, *uuid.UUID]
	getHandler    shared.IQueryHandler[*menuservice.GetMenuQuery, *menudomain.Menu]
	listHandler   shared.IQueryHandler[*menuservice.ListMenuQuery, *menuservice.ListMenuResult]
	updateHandler shared.ICommandHandler[*menuservice.UpdateMenuCommand]
	deleteHandler shared.ICommandHandler[*menuservice.DeleteMenuCommand]
}

func NewHTTPMenuController(
	createHandler shared.ICommandResultHandler[*menuservice.CreateMenuResultCommand, *uuid.UUID],
	getHandler shared.IQueryHandler[*menuservice.GetMenuQuery, *menudomain.Menu],
	listHandler shared.IQueryHandler[*menuservice.ListMenuQuery, *menuservice.ListMenuResult],
	updateHandler shared.ICommandHandler[*menuservice.UpdateMenuCommand],
	deleteHandler shared.ICommandHandler[*menuservice.DeleteMenuCommand],
) *HTTPMenuController {
	return &HTTPMenuController{
		createHandler: createHandler,
		getHandler:    getHandler,
		listHandler:   listHandler,
		updateHandler: updateHandler,
		deleteHandler: deleteHandler,
	}
}

func (ctrl *HTTPMenuController) SetupRoutes(v1 *gin.RouterGroup) {
	menuGroup := v1.Group("/menus")
	{
		menuGroup.POST("", ctrl.CreateMenuAPI())
		menuGroup.GET("/:id", ctrl.GetMenuByIdAPI())
		menuGroup.GET("", ctrl.ListMenuAPI())
		menuGroup.PATCH("/:id", ctrl.UpdateMenuAPI())
		menuGroup.DELETE("/:id", ctrl.DeleteMenuAPI())
	}
}

