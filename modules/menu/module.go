package menumodule

import (
	menucontroller "vht-go/modules/menu/infras/controller"
	menurepository "vht-go/modules/menu/infras/repository"
	"vht-go/modules/menu/infras/repository/menurpcclient"
	menuservice "vht-go/modules/menu/service"
	"vht-go/shared"
	sharedcomponent "vht-go/shared/component"

	"github.com/gin-gonic/gin"
	sctx "github.com/viettranx/service-context"
)

// SetupMenuModule initializes and registers the menu module
// Dependencies Injection
func SetupMenuModule(v1 *gin.RouterGroup, sctx sctx.ServiceContext) {
	appConfig := sctx.MustGet(sharedcomponent.AppConfigID).(sharedcomponent.IAppConfig)
	db := sctx.MustGet(shared.KeyGormComp).(sharedcomponent.IGormComp).DB()

	// 1. Initialize repository
	repo := menurepository.NewGORMMenuRepository(db)

	// 2. Initialize RPC clients
	foodRPCClient := menurpcclient.NewFoodRPCClient(appConfig.FoodServiceURI())
	restaurantRPCClient := menurpcclient.NewRestaurantRPCClient(appConfig.RestaurantServiceURI())

	// 3. Initialize handlers with repository and RPC clients
	createHandler := menuservice.NewCreateMenuResultCommandHandler(repo)
	getHandler := menuservice.NewGetMenuQueryHandler(repo, foodRPCClient, restaurantRPCClient)
	listHandler := menuservice.NewListMenuQueryHandler(repo, foodRPCClient, restaurantRPCClient)
	updateHandler := menuservice.NewUpdateMenuCommandHandler(repo)
	deleteHandler := menuservice.NewDeleteMenuCommandHandler(repo, repo)

	// 4. Initialize controller with handlers
	controller := menucontroller.NewHTTPMenuController(
		createHandler,
		getHandler,
		listHandler,
		updateHandler,
		deleteHandler,
	)

	// 5. Setup routes
	controller.SetupRoutes(v1)
}

