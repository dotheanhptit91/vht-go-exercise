package cmd

import (
	"context"
	"log"
	"net/http"
	"os"
	"strconv"
	"vht-go/middleware"
	categorymodule "vht-go/modules/category"
	foodmodule "vht-go/modules/food"
	restaurantmodule "vht-go/modules/restaurant"
	restaurantrepository "vht-go/modules/restaurant/infras/repository"
	rstlikemodule "vht-go/modules/restaurantlike"
	usermodule "vht-go/modules/user"
	"vht-go/shared"
	"vht-go/shared/asyncjob"
	sharedcomponent "vht-go/shared/component"
	"vht-go/shared/component/pubsub"

	"github.com/gin-gonic/gin"
	"github.com/spf13/cobra"
	sctx "github.com/viettranx/service-context"
)

func newService() sctx.ServiceContext {
	serviceCtx := sctx.NewServiceContext(
		sctx.WithComponent(sharedcomponent.NewAppConfig()),
		sctx.WithComponent(sharedcomponent.NewGormComp(shared.KeyGormComp)),
		sctx.WithComponent(pubsub.NewPubSub(shared.KeyLocalPubSubComp)),
		sctx.WithComponent(pubsub.NewNatsComp(shared.KeyNatsPubSubComp)),
		sctx.WithComponent(sharedcomponent.NewRedisComp(shared.KeyRedisComp)),
		sctx.WithComponent(sharedcomponent.NewGrpcServerComp(shared.KeyGrpcServerComp)),
	)

	return serviceCtx
}

var rootCmd = &cobra.Command{
	Use:   "app",
	Short: "Run application",
	Long:  "Run application with all dependencies",
	Run: func(cmd *cobra.Command, args []string) {
		// Init service-context, you can put components as much as you can
		serviceCtx := newService()

		// Load() will iterate registered components
		// It does parse flags and make some configurations if you defined
		// in Activate() method of the components
		if err := serviceCtx.Load(); err != nil {
			log.Fatal(err)
		}

		expIn := 60 * 60 * 24 * 7

		if os.Getenv("JWT_EXP_IN") != "" {
			expInInt, err := strconv.Atoi(os.Getenv("JWT_EXP_IN"))
			if err != nil {
				log.Fatalln("failed to convert JWT_EXP_IN to int")
			}
			expIn = expInInt
		}

		db := serviceCtx.MustGet(shared.KeyGormComp).(sharedcomponent.IGormComp).DB()

		jwtComponent := sharedcomponent.NewJwtComp(os.Getenv("JWT_SECRET_KEY"), expIn)

		tokenIntrospector := sharedcomponent.NewTokenIntrospector(jwtComponent, db)
		middlewareProvider := middleware.NewMiddlewareProvider(tokenIntrospector)

		// Create a Gin router with default middleware (logger and recovery)
		r := gin.Default()

		// r.Use(myLogger())
		r.Use(middleware.RecoverMiddleware())

		v1 := r.Group("/v1")
		categorymodule.SetupCategoryModule(v1, serviceCtx)
		restaurantmodule.SetupRestaurantModule(v1, serviceCtx)
		foodmodule.SetupFoodModule(v1, serviceCtx)
		usermodule.SetupUserModule(v1, serviceCtx, jwtComponent, middlewareProvider)
		rstlikemodule.SetupRestaurantLikeModule(v1, serviceCtx, middlewareProvider)

		// Define a simple GET endpoint
		r.GET("/ping", func(c *gin.Context) {
			// panic("error in ping")
			// Return JSON response
			c.JSON(http.StatusOK, gin.H{
				"message": "pong",
			})
		})

		// runConsumers(serviceCtx)

		// Start server on port 8080 (default)
		// Server will listen on 0.0.0.0:8080 (localhost:8080 on Windows)
		grpcServer := serviceCtx.MustGet(shared.KeyGrpcServerComp).(sharedcomponent.IGrpcServerComp)
		grpcServer.Serve()
		r.Run()
	},
}

func Execute() {
	rootCmd.AddCommand(outenvCmd)
	rootCmd.AddCommand(increaseLikedCountCmd)
	rootCmd.AddCommand(decreaseLikedCountCmd)

	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}

// var consumer = &cobra.Command{
// 	Use:   "consumer",
// 	Short: "Run consumer",
// 	Long:  "Run consumer with all dependencies",
// 	Run: func(cmd *cobra.Command, args []string) {

// 	},
// }

func runConsumers(serviceCtx sctx.ServiceContext) {

	ps := serviceCtx.MustGet(shared.KeyLocalPubSubComp).(pubsub.IPubSub)
	db := serviceCtx.MustGet(shared.KeyGormComp).(sharedcomponent.IGormComp).DB()

	chUnliked, _ := ps.Subscribe(context.Background(), pubsub.Topic(shared.EvtRestaurantUnliked))

	repo := restaurantrepository.NewGORMRestaurantRepository(db)

	go func() {
		serviceCtx.Logger("consumers").Debugln("Running decrease liked count consumer")
		defer shared.RecoverApp()

		for msg := range chUnliked {
			data := msg.Data().(map[string]interface{})
			restaurantId := data["restaurantId"].(int)
			// userId := data["userId"].(string)

			job := asyncjob.NewJob(func(ctx context.Context) error {
				return repo.DecreaseLikedCount(ctx, restaurantId)
			})

			if err := asyncjob.NewGroup(false, job).Run(context.Background()); err != nil {
				serviceCtx.Logger("consumers").Errorln("Error running decrease liked count job", err)
			}

			serviceCtx.Logger("consumers").Debugln("Decreased liked count for restaurant", restaurantId)
		}
	}()

	chLiked, _ := ps.Subscribe(context.Background(), pubsub.Topic(shared.EvtRestaurantLiked))

	go func() {
		serviceCtx.Logger("consumers").Debugln("Running increase liked count consumer")
		defer shared.RecoverApp()

		for msg := range chLiked {
			data := msg.Data().(map[string]interface{})
			restaurantId := data["restaurantId"].(int)
			// userId := data["userId"].(string)
			job := asyncjob.NewJob(func(ctx context.Context) error {
				return repo.IncreaseLikedCount(ctx, restaurantId)
			})

			if err := asyncjob.NewGroup(false, job).Run(context.Background()); err != nil {
				serviceCtx.Logger("consumers").Errorln("Error running increase liked count job", err)
			}

			serviceCtx.Logger("consumers").Debugln("Increased liked count for restaurant", restaurantId)
		}
	}()
}