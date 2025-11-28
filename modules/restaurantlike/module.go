package rstlikemodule

import (
	"vht-go/middleware"
	rstlikecontroller "vht-go/modules/restaurantlike/infras/controller"
	rstlikerepository "vht-go/modules/restaurantlike/infras/repository"
	rstlikeservice "vht-go/modules/restaurantlike/service"
	"vht-go/shared"
	sharedcomponent "vht-go/shared/component"
	"vht-go/shared/component/pubsub"

	"github.com/gin-gonic/gin"
	sctx "github.com/viettranx/service-context"
)

func SetupRestaurantLikeModule(v1 *gin.RouterGroup, sctx sctx.ServiceContext, middlewareProvider middleware.IMiddlewareProvider) {
	db := sctx.MustGet(shared.KeyGormComp).(sharedcomponent.IGormComp).DB()
	// ps := sctx.MustGet(shared.KeyLocalPubSubComp).(pubsub.IPubSub)
	natsPS := sctx.MustGet(shared.KeyNatsPubSubComp).(pubsub.IPubSub)

	restaurantRepository := rstlikerepository.NewGORMRestaurantRepository(db)
	repo := rstlikerepository.NewGORMRestaurantLikeRepository(db)

	// updateCountersRepository := repository.NewGORMRestaurantRepository(db)
	likeRestaurantCommandHandler := rstlikeservice.NewLikeRestaurantCommandHandler(restaurantRepository, repo, natsPS)
	unlikeRestaurantCommandHandler := rstlikeservice.NewUnlikeRestaurantCommandHandler(repo, natsPS)

	ctrl := rstlikecontroller.NewHTTPRestaurantLikeController(likeRestaurantCommandHandler, unlikeRestaurantCommandHandler)

	ctrl.SetupRouter(v1, middlewareProvider)
}