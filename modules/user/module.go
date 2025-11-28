package usermodule

import (
	"vht-go/middleware"
	usercontroller "vht-go/modules/user/infras/controller"
	userrepository "vht-go/modules/user/infras/repository"
	userservice "vht-go/modules/user/service"
	"vht-go/shared"
	sharedcomponent "vht-go/shared/component"

	"github.com/gin-gonic/gin"
	sctx "github.com/viettranx/service-context"
)

func SetupUserModule(v1 *gin.RouterGroup, sctx sctx.ServiceContext, jwtComponent userservice.IJWTComponent, middlewareProvider middleware.IMiddlewareProvider) {
	db := sctx.MustGet(shared.KeyGormComp).(sharedcomponent.IGormComp).DB()

	repo := userrepository.NewGORMUserRepository(db)
	registerUserHandler := userservice.NewRegisterUserCommandHandler(repo)
	loginUserHandler := userservice.NewLoginUserCommandHandler(repo, jwtComponent)

	controller := usercontroller.NewHTTPUserController(
		registerUserHandler,
		loginUserHandler,
	)

	controller.SetupRouter(v1, middlewareProvider)
}