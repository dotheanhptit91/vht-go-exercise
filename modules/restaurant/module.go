package restaurantmodule

import (
	"os"
	restaurantcontroller "vht-go/modules/restaurant/infras/controller"
	"vht-go/modules/restaurant/infras/controller/restaurantrpcserver"
	restaurantrepository "vht-go/modules/restaurant/infras/repository"
	restaurantgrpcclient "vht-go/modules/restaurant/infras/repository/grpc-client"
	"vht-go/modules/restaurant/infras/repository/restaurantrpcclient"
	restaurantservice "vht-go/modules/restaurant/service"
	"vht-go/shared"
	sharedcomponent "vht-go/shared/component"

	"github.com/gin-gonic/gin"
	sctx "github.com/viettranx/service-context"
)

// Dependencies Injection
func SetupRestaurantModule(rg *gin.RouterGroup, sctx sctx.ServiceContext) {
	appConfig := sctx.MustGet(sharedcomponent.AppConfigID).(sharedcomponent.IAppConfig)
	db := sctx.MustGet(shared.KeyGormComp).(sharedcomponent.IGormComp).DB()

	rstGRPCClient := restaurantgrpcclient.NewRestaurantGrpcClient(os.Getenv("CATEGORY_GRPC_URI"))


	// 1. Initialize repository
	repo := restaurantrepository.NewGORMRestaurantRepository(db)
	rpcClient := restaurantrpcclient.NewCategoryRPCClient(appConfig.CategoryServiceURI())

	categoryCachedRPCClient := restaurantrpcclient.NewGetCategoryCachedRPCClient(
		rstGRPCClient,
		sctx.MustGet(shared.KeyRedisComp).(sharedcomponent.IRedisComp),
	)

	// 2. Initialize handlers with repository
	createHandler := restaurantservice.NewCreateRestaurantResultCommandHandler(repo)
	getHandler := restaurantservice.NewGetRestaurantQueryHandler(repo, categoryCachedRPCClient)
	listHandler := restaurantservice.NewListRestaurantQueryHandler(repo, rpcClient)
	updateHandler := restaurantservice.NewUpdateRestaurantCommandHandler(repo)
	deleteHandler := restaurantservice.NewDeleteRestaurantCommandHandler(repo, repo)

	// 3. Initialize controller with handlers
	controller := restaurantcontroller.NewHTTPRestaurantController(
		createHandler,
		getHandler,
		listHandler,
		updateHandler,
		deleteHandler,
	)

	// 4. Setup routes
	controller.SetupRoutes(rg)

	// 5. Setup RPC server
	rpcServer := restaurantrpcserver.NewRestaurantRPCServer(db)
	rpcServer.SetupRouter(rg)
}

