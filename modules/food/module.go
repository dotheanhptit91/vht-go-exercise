package foodmodule

import (
	foodcontroller "vht-go/modules/food/infras/controller"
	"vht-go/modules/food/infras/controller/foodrpcserver"
	foodrepository "vht-go/modules/food/infras/repository"
	"vht-go/modules/food/infras/repository/foodrpcclient"
	foodservice "vht-go/modules/food/service"
	"vht-go/shared"
	sharedcomponent "vht-go/shared/component"

	"github.com/gin-gonic/gin"
	sctx "github.com/viettranx/service-context"
)

// Dependencies Injection
func SetupFoodModule(v1 *gin.RouterGroup, sctx sctx.ServiceContext) {
	appConfig := sctx.MustGet(sharedcomponent.AppConfigID).(sharedcomponent.IAppConfig)
	db := sctx.MustGet(shared.KeyGormComp).(sharedcomponent.IGormComp).DB()

	// 1. Initialize repository
	repo := foodrepository.NewGORMFoodRepository(db)

	// 2. Initialize RPC clients
	categoryRPCClient := foodrpcclient.NewCategoryRPCClient(appConfig.CategoryServiceURI())
	restaurantRPCClient := foodrpcclient.NewRestaurantRPCClient(appConfig.RestaurantServiceURI())

	// 3. Initialize handlers with repository and RPC clients
	createHandler := foodservice.NewCreateFoodResultCommandHandler(repo)
	getHandler := foodservice.NewGetFoodQueryHandler(repo, categoryRPCClient, restaurantRPCClient)
	listHandler := foodservice.NewListFoodQueryHandler(repo, categoryRPCClient, restaurantRPCClient)
	updateHandler := foodservice.NewUpdateFoodCommandHandler(repo)
	deleteHandler := foodservice.NewDeleteFoodCommandHandler(repo)

	// 4. Initialize controller with handlers
	controller := foodcontroller.NewHTTPFoodController(
		createHandler,
		getHandler,
		listHandler,
		updateHandler,
		deleteHandler,
	)

	// 5. Setup routes
	controller.SetupRoutes(v1)

	// 6. Setup RPC server
	rpcServer := foodrpcserver.NewFoodRPCServer(db)
	rpcServer.SetupRouter(v1)
}

