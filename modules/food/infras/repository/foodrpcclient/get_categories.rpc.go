package foodrpcclient

import (
	"context"
	"fmt"
	fooddomain "vht-go/modules/food/domain"

	"github.com/google/uuid"
	"resty.dev/v3"
)

func (c *CategoryRPCClient) FindCategoriesByIds(ctx context.Context, ids []uuid.UUID) ([]fooddomain.FoodCategory, error) {
	fullURL := fmt.Sprintf("%s/get-categories", c.categoryServiceURI)

	var dataRPC struct {
		Data []fooddomain.FoodCategory `json:"data"`
	}

	_, err := resty.New().R().
		SetBody(map[string]any{
			"ids": ids,
		}).
		SetResult(&dataRPC).
		Post(fullURL)

	return dataRPC.Data, err
}

