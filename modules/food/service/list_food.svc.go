package foodservice

import (
	"context"

	fooddomain "vht-go/modules/food/domain"
	"vht-go/shared"

	"github.com/google/uuid"
)

type ListFoodQuery struct {
	RestaurantId *int
	CategoryId   *string
	Status       *int
	Paging       *shared.Paging
}

type IListFoodRepository interface {
	FindAll(ctx context.Context, filter map[string]interface{}, limit, offset int) ([]fooddomain.Food, error)
	Count(ctx context.Context, filter map[string]interface{}) (int64, error)
}

type IGetFoodCategories interface {
	FindCategoriesByIds(ctx context.Context, ids []uuid.UUID) ([]fooddomain.FoodCategory, error)
}

type IGetFoodRestaurants interface {
	FindRestaurantsByIds(ctx context.Context, ids []int) ([]fooddomain.FoodRestaurant, error)
}

type ListFoodQueryHandler struct {
	repo           IListFoodRepository
	categoryRepo   IGetFoodCategories
	restaurantRepo IGetFoodRestaurants
}

func NewListFoodQueryHandler(repo IListFoodRepository, categoryRepo IGetFoodCategories, restaurantRepo IGetFoodRestaurants) *ListFoodQueryHandler {
	return &ListFoodQueryHandler{repo: repo, categoryRepo: categoryRepo, restaurantRepo: restaurantRepo}
}

func (h *ListFoodQueryHandler) Handle(ctx context.Context, query *ListFoodQuery) ([]fooddomain.Food, error) {
	query.Paging.Process()

	filter := make(map[string]interface{})
	if query.RestaurantId != nil {
		filter["restaurant_id"] = *query.RestaurantId
	}
	if query.CategoryId != nil {
		filter["category_id"] = *query.CategoryId
	}
	if query.Status != nil {
		filter["status"] = *query.Status
	}

	// Get total count
	total, err := h.repo.Count(ctx, filter)
	if err != nil {
		return nil, err
	}
	query.Paging.Total = total

	// Get foods
	offset := (query.Paging.Page - 1) * query.Paging.Limit
	foods, err := h.repo.FindAll(ctx, filter, query.Paging.Limit, offset)
	if err != nil {
		return nil, err
	}

	// Collect category IDs and restaurant IDs
	catIds := make([]uuid.UUID, 0)
	rstIds := make([]int, 0)
	rstIdMap := make(map[int]bool)

	for _, food := range foods {
		if food.CategoryId != nil {
			catIds = append(catIds, *food.CategoryId)
		}
		if !rstIdMap[food.RestaurantId] {
			rstIds = append(rstIds, food.RestaurantId)
			rstIdMap[food.RestaurantId] = true
		}
	}

	// Fetch categories
	var catMap map[uuid.UUID]fooddomain.FoodCategory
	if len(catIds) > 0 {
		categories, err := h.categoryRepo.FindCategoriesByIds(ctx, catIds)
		if err == nil {
			catMap = make(map[uuid.UUID]fooddomain.FoodCategory)
			for _, category := range categories {
				catMap[category.Id] = category
			}
		}
	}

	// Fetch restaurants
	var rstMap map[int]fooddomain.FoodRestaurant
	if len(rstIds) > 0 {
		restaurants, err := h.restaurantRepo.FindRestaurantsByIds(ctx, rstIds)
		if err == nil {
			rstMap = make(map[int]fooddomain.FoodRestaurant)
			for _, restaurant := range restaurants {
				rstMap[restaurant.RestaurantId] = restaurant
			}
		}
	}

	// Populate foods with category and restaurant data
	for i := range foods {
		if foods[i].CategoryId != nil {
			if category, exists := catMap[*foods[i].CategoryId]; exists {
				foods[i].Category = &category
			}
		}
		if restaurant, exists := rstMap[foods[i].RestaurantId]; exists {
			foods[i].Restaurant = &restaurant
		}
	}

	return foods, nil
}

