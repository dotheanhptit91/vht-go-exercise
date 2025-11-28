package categorymodule

import (
	categorygrpc "vht-go/gen/proto/category"
	categorycontroller "vht-go/modules/category/infras/controller"
	"vht-go/modules/category/infras/controller/categoryrpcserver"
	categorygrpcserver "vht-go/modules/category/infras/controller/grpc-server"
	categoryrepository "vht-go/modules/category/infras/repository"
	categoryservice "vht-go/modules/category/service"
	"vht-go/shared"
	sharedcomponent "vht-go/shared/component"

	"github.com/gin-gonic/gin"
	sctx "github.com/viettranx/service-context"
	"google.golang.org/grpc"
)

// Dependencies Injection
func SetupCategoryModule(v1 *gin.RouterGroup, sctx sctx.ServiceContext) {
	db := sctx.MustGet(shared.KeyGormComp).(sharedcomponent.IGormComp).DB()

	repo := categoryrepository.NewGORMCategoryRepository(db)
	// service := categoryservice.NewCategoryService(repo)
	createHandler := categoryservice.NewCreateCategoryResultCommandHandler(repo)
	getHandler := categoryservice.NewGetCategoryQueryHandler(repo)
	listHandler := categoryservice.NewListCategoryQueryHandler(repo)
	updateHandler := categoryservice.NewUpdateCategoryCommandHandler(repo)
	deleteHandler := categoryservice.NewDeleteCategoryCommandHandler(repo, repo)
	
	controller := categorycontroller.NewHTTPCategoryController(
		createHandler, 
		getHandler, 
		listHandler, 
		updateHandler,
		deleteHandler)

	controller.SetupRoutes(v1)

	rpcServer := categoryrpcserver.NewCategoryRPCServer(db)
	rpcServer.SetupRouter(v1)

	// GRPC Server Registers
	grpcServerComp := categorygrpcserver.NewCategoryGrpcServer(repo)

	grpcServer := sctx.MustGet(shared.KeyGrpcServerComp).(sharedcomponent.IGrpcServerComp)
	grpcServer.Register(func(s *grpc.Server) {
		categorygrpc.RegisterCategoryServiceServer(s, grpcServerComp)
	})
}