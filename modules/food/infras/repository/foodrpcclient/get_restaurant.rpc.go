package foodrpcclient

import (
	"context"
	"fmt"
	fooddomain "vht-go/modules/food/domain"

	"resty.dev/v3"
)

func (c *RestaurantRPCClient) FindRestaurantById(ctx context.Context, id int) (*fooddomain.FoodRestaurant, error) {
	fullURL := fmt.Sprintf("%s/get-restaurant", c.restaurantServiceURI)

	var dataRPC struct {
		Data fooddomain.FoodRestaurant `json:"data"`
	}

	_, err := resty.New().R().
		SetBody(map[string]any{
			"id": id,
		}).
		SetResult(&dataRPC).
		Post(fullURL)

	if err != nil {
		return nil, err
	}

	return &dataRPC.Data, nil
}

