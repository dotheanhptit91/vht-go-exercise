package cmd

import (
	"context"
	"fmt"
	"log"
	"strconv"
	restaurantrepository "vht-go/modules/restaurant/infras/repository"
	"vht-go/shared"
	sharedcomponent "vht-go/shared/component"
	"vht-go/shared/component/pubsub"

	"github.com/spf13/cobra"
)

var decreaseLikedCountCmd = &cobra.Command{
	Use:   "decrease-liked-count",
	Short: "Decrease liked count",
	Long:  "Decrease liked count",
	Run: func(cmd *cobra.Command, args []string) {
		// Init service-context, you can put components as much as you can
		serviceCtx := newService()

		if err := serviceCtx.Load(); err != nil {
			log.Fatal(err)
		}

		ps := serviceCtx.MustGet(shared.KeyNatsPubSubComp).(pubsub.IPubSub)
		db := serviceCtx.MustGet(shared.KeyGormComp).(sharedcomponent.IGormComp).DB()

		ch, _ := ps.Subscribe(context.Background(), pubsub.Topic(shared.EvtRestaurantUnliked))

		repo := restaurantrepository.NewGORMRestaurantRepository(db)

		for msg := range ch {
			data := msg.Data().(map[string]interface{})
			restaurantIdStr := fmt.Sprintf("%v", data["restaurantId"])
			restaurantId, err := strconv.Atoi(restaurantIdStr)
			if err != nil {
				serviceCtx.Logger("increase-liked-count").Errorln("Invalid restaurantId", data)
				continue
			}
			// userId := data["userId"].(string)
			repo.DecreaseLikedCount(context.Background(), restaurantId)

			serviceCtx.Logger("decrease-liked-count").Debugln("Decreased liked count for restaurant", restaurantId)
		}
	},
}