package categorycontroller

import (
	categorydomain "vht-go/modules/category/domain"
	categoryservice "vht-go/modules/category/service"
	"vht-go/shared"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type HTTPCategoryController struct {
	// svc categoryservice.ICategoryService/
	createHandler shared.ICommandResultHandler[*categoryservice.CreateCategoryResultCommand, *uuid.UUID]
	getHandler shared.IQueryHandler[*categoryservice.GetCategoryQuery, *categorydomain.Category]
	listHandler shared.IQueryHandler[*categoryservice.ListCategoryQuery, []categorydomain.Category]
	updateHandler shared.ICommandHandler[*categoryservice.UpdateCategoryCommand]
	deleteHandler shared.ICommandHandler[*categoryservice.DeleteCategoryCommand]
}

func NewHTTPCategoryController(
	// svc categoryservice.ICategoryService, 
	createHandler shared.ICommandResultHandler[*categoryservice.CreateCategoryResultCommand, *uuid.UUID],
	getHandler shared.IQueryHandler[*categoryservice.GetCategoryQuery, *categorydomain.Category],
	listHandler shared.IQueryHandler[*categoryservice.ListCategoryQuery, []categorydomain.Category],
	updateHandler shared.ICommandHandler[*categoryservice.UpdateCategoryCommand],
	deleteHandler shared.ICommandHandler[*categoryservice.DeleteCategoryCommand]) *HTTPCategoryController {
	return &HTTPCategoryController{
		// svc: svc, 
		createHandler: createHandler,
		getHandler: getHandler, 
		listHandler: listHandler, 
		updateHandler: updateHandler,
		deleteHandler: deleteHandler,
	}
}

func (ctrl *HTTPCategoryController) SetupRoutes(v1 *gin.RouterGroup) {
	catGroup := v1.Group("/categories")
	catGroup.POST("", ctrl.CreateCategoryAPI())
	catGroup.GET("/:id", ctrl.GetCategoryByIdAPI())
	catGroup.GET("", ctrl.GetListCategoriesAPI())
	catGroup.PATCH("/:id", ctrl.UpdateCategoryAPI())
	catGroup.DELETE("/:id", ctrl.DeleteCategory())
}