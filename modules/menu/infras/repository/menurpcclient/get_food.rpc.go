package menurpcclient

import (
	"context"
	"fmt"
	menudomain "vht-go/modules/menu/domain"

	"resty.dev/v3"
)

func (c *FoodRPCClient) FindFoodById(ctx context.Context, id int) (*menudomain.MenuFood, error) {
	fullURL := fmt.Sprintf("%s/%d", c.foodServiceURI, id)

	var dataRPC struct {
		Data menudomain.MenuFood `json:"data"`
	}

	_, err := resty.New().R().
		SetResult(&dataRPC).
		Get(fullURL)

	if err != nil {
		return nil, err
	}

	return &dataRPC.Data, nil
}

