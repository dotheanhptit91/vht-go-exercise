package foodservice

import (
	"context"

	fooddomain "vht-go/modules/food/domain"

	"github.com/google/uuid"
)

type GetFoodQuery struct {
	Id int
}

type IGetFoodRepository interface {
	FindById(ctx context.Context, id int) (*fooddomain.Food, error)
}

type IGetFoodCategory interface {
	FindCategoryById(ctx context.Context, id uuid.UUID) (*fooddomain.FoodCategory, error)
}

type IGetFoodRestaurant interface {
	FindRestaurantById(ctx context.Context, id int) (*fooddomain.FoodRestaurant, error)
}

type GetFoodQueryHandler struct {
	repo           IGetFoodRepository
	categoryRepo   IGetFoodCategory
	restaurantRepo IGetFoodRestaurant
}

func NewGetFoodQueryHandler(repo IGetFoodRepository, categoryRepo IGetFoodCategory, restaurantRepo IGetFoodRestaurant) *GetFoodQueryHandler {
	return &GetFoodQueryHandler{repo: repo, categoryRepo: categoryRepo, restaurantRepo: restaurantRepo}
}

func (h *GetFoodQueryHandler) Handle(ctx context.Context, query *GetFoodQuery) (*fooddomain.Food, error) {
	food, err := h.repo.FindById(ctx, query.Id)

	if err != nil {
		return nil, err
	}

	// Populate category if exists
	if food != nil && food.CategoryId != nil {
		category, err := h.categoryRepo.FindCategoryById(ctx, *food.CategoryId)
		if err == nil {
			food.Category = category
		}
	}

	// Populate restaurant
	if food != nil {
		restaurant, err := h.restaurantRepo.FindRestaurantById(ctx, food.RestaurantId)
		if err == nil {
			food.Restaurant = restaurant
		}
	}

	return food, nil
}

