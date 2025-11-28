package menurpcclient

import (
	"context"
	"fmt"
	menudomain "vht-go/modules/menu/domain"

	"resty.dev/v3"
)

func (c *FoodRPCClient) FindFoodsByIds(ctx context.Context, ids []int) ([]menudomain.MenuFood, error) {
	if len(ids) == 0 {
		return []menudomain.MenuFood{}, nil
	}

	fullURL := fmt.Sprintf("%s/get-foods", c.foodServiceURI)

	var dataRPC struct {
		Data []menudomain.MenuFood `json:"data"`
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

