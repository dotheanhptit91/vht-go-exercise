package menurpcclient

import (
	"context"
	"fmt"
	menudomain "vht-go/modules/menu/domain"

	"resty.dev/v3"
)

func (c *RestaurantRPCClient) FindRestaurantById(ctx context.Context, id int) (*menudomain.MenuRestaurant, error) {
	fullURL := fmt.Sprintf("%s/%d", c.restaurantServiceURI, id)

	var dataRPC struct {
		Data menudomain.MenuRestaurant `json:"data"`
	}

	_, err := resty.New().R().
		SetResult(&dataRPC).
		Get(fullURL)

	if err != nil {
		return nil, err
	}

	return &dataRPC.Data, nil
}

