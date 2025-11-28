package restaurantrpcclient

import (
	"context"
	"fmt"
	restaurantdomain "vht-go/modules/restaurant/domain"

	"github.com/google/uuid"
	"resty.dev/v3"
)

func (c *CategoryRPCClient) FindCategoryById(ctx context.Context, id uuid.UUID) (*restaurantdomain.RestaurantCategory, error) {
	fullURL := fmt.Sprintf("%s/get-category", c.categoryServiceURI)

	var dataRPC struct {
		Data restaurantdomain.RestaurantCategory `json:"data"`
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