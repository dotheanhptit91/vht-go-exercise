package restaurantrpcclient

import (
	"context"
	"fmt"
	restaurantdomain "vht-go/modules/restaurant/domain"

	"github.com/google/uuid"
	"resty.dev/v3"
)

func (c *CategoryRPCClient) FindCategoriesByIds(ctx context.Context, ids []uuid.UUID) ([]restaurantdomain.RestaurantCategory, error) {
	fullURL := fmt.Sprintf("%s/get-categories", c.categoryServiceURI)

	var dataRPC struct {
		Data []restaurantdomain.RestaurantCategory `json:"data"`
	}

	_, err := resty.New().R().
		SetBody(map[string]any{
			"ids": ids,
		}).
		SetResult(&dataRPC).
		Post(fullURL)

	return dataRPC.Data, err
}