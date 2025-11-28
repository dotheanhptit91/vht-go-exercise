package foodrpcclient

import (
	"context"
	"fmt"
	fooddomain "vht-go/modules/food/domain"

	"resty.dev/v3"
)

func (c *RestaurantRPCClient) FindRestaurantsByIds(ctx context.Context, ids []int) ([]fooddomain.FoodRestaurant, error) {
	fullURL := fmt.Sprintf("%s/get-restaurants", c.restaurantServiceURI)

	var dataRPC struct {
		Data []fooddomain.FoodRestaurant `json:"data"`
	}

	_, err := resty.New().R().
		SetBody(map[string]any{
			"ids": ids,
		}).
		SetResult(&dataRPC).
		Post(fullURL)

	return dataRPC.Data, err
}

