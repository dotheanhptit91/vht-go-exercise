package restaurantservice

import (
	"context"
	restaurantdomain "vht-go/modules/restaurant/domain"

	"github.com/google/uuid"
)

type GetRestaurantQuery struct {
	Id int
}

type IGetRestaurantRepository interface {
	FindById(ctx context.Context, id int) (*restaurantdomain.Restaurant, error)
}

type IGetRestaurantCategory interface {
	FindCategoryById(ctx context.Context, id uuid.UUID) (*restaurantdomain.RestaurantCategory, error)
}

type GetRestaurantQueryHandler struct {
	categoryRepo IGetRestaurantCategory
	repo IGetRestaurantRepository
}

func NewGetRestaurantQueryHandler(repo IGetRestaurantRepository, categoryRepo IGetRestaurantCategory) *GetRestaurantQueryHandler {
	return &GetRestaurantQueryHandler{repo: repo, categoryRepo: categoryRepo}
}

func (h *GetRestaurantQueryHandler) Handle(ctx context.Context, query *GetRestaurantQuery) (*restaurantdomain.Restaurant, error) {
	restaurant, err := h.repo.FindById(ctx, query.Id)

	if err != nil {
		return nil, err
	}

	if restaurant != nil && restaurant.CategoryId != nil {
		category, err := h.categoryRepo.FindCategoryById(ctx, *restaurant.CategoryId)
		if err != nil {
			return nil, err
		}
		restaurant.Category = category
	}

	return restaurant, nil	
}