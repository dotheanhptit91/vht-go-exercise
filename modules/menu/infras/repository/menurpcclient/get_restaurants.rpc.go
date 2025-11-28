package menurpcclient

import (
	"context"
	"fmt"
	menudomain "vht-go/modules/menu/domain"

	"resty.dev/v3"
)

func (c *RestaurantRPCClient) FindRestaurantsByIds(ctx context.Context, ids []int) ([]menudomain.MenuRestaurant, error) {
	if len(ids) == 0 {
		return []menudomain.MenuRestaurant{}, nil
	}

	fullURL := fmt.Sprintf("%s/get-restaurants", c.restaurantServiceURI)

	var dataRPC struct {
		Data []menudomain.MenuRestaurant `json:"data"`
	}

	_, err := resty.New().R().
		SetBody(map[string]any{
			"ids": ids,
		}).
		SetResult(&dataRPC).
		Post(fullURL)

	if err != nil {
		return nil, err
	}

	return dataRPC.Data, nil
}

