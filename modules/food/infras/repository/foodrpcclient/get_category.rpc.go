package foodrpcclient

import (
	"context"
	"fmt"
	fooddomain "vht-go/modules/food/domain"

	"github.com/google/uuid"
	"resty.dev/v3"
)

func (c *CategoryRPCClient) FindCategoryById(ctx context.Context, id uuid.UUID) (*fooddomain.FoodCategory, error) {
	fullURL := fmt.Sprintf("%s/get-category", c.categoryServiceURI)

	var dataRPC struct {
		Data fooddomain.FoodCategory `json:"data"`
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

