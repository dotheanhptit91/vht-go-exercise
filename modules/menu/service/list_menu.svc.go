package menuservice

import (
	"context"
	menudomain "vht-go/modules/menu/domain"
	menudtos "vht-go/modules/menu/dtos"
)

type ListMenuQuery struct {
	DTO *menudtos.ListMenuDTO
}

type IListMenuRepository interface {
	FindAll(ctx context.Context, filter map[string]interface{}, limit, offset int) ([]menudomain.Menu, error)
	Count(ctx context.Context, filter map[string]interface{}) (int64, error)
}

type IListMenuFoods interface {
	FindFoodsByIds(ctx context.Context, ids []int) ([]menudomain.MenuFood, error)
}

type IListMenuRestaurants interface {
	FindRestaurantsByIds(ctx context.Context, ids []int) ([]menudomain.MenuRestaurant, error)
}

type ListMenuQueryHandler struct {
	repo           IListMenuRepository
	foodRepo       IListMenuFoods
	restaurantRepo IListMenuRestaurants
}

func NewListMenuQueryHandler(
	repo IListMenuRepository,
	foodRepo IListMenuFoods,
	restaurantRepo IListMenuRestaurants,
) *ListMenuQueryHandler {
	return &ListMenuQueryHandler{
		repo:           repo,
		foodRepo:       foodRepo,
		restaurantRepo: restaurantRepo,
	}
}

type ListMenuResult struct {
	Data   []menudomain.Menu `json:"data"`
	Paging interface{}       `json:"paging"`
}

func (h *ListMenuQueryHandler) Handle(ctx context.Context, query *ListMenuQuery) (*ListMenuResult, error) {
	query.DTO.Paging.Process()

	filter := make(map[string]interface{})
	if query.DTO.RestaurantId != nil {
		filter["restaurant_id"] = *query.DTO.RestaurantId
	}
	if query.DTO.Status != nil {
		filter["status"] = *query.DTO.Status
	}

	// Get total count
	total, err := h.repo.Count(ctx, filter)
	if err != nil {
		return nil, err
	}
	query.DTO.Paging.Total = total

	// Get menus
	offset := (query.DTO.Paging.Page - 1) * query.DTO.Paging.Limit
	menus, err := h.repo.FindAll(ctx, filter, query.DTO.Paging.Limit, offset)
	if err != nil {
		return nil, err
	}

	// Collect all unique food IDs and restaurant IDs
	foodIdsMap := make(map[int]bool)
	restaurantIdsMap := make(map[int]bool)

	for _, menu := range menus {
		for _, foodId := range menu.FoodIds {
			foodIdsMap[foodId] = true
		}
		restaurantIdsMap[menu.RestaurantId] = true
	}

	// Convert maps to slices
	var allFoodIds []int
	for foodId := range foodIdsMap {
		allFoodIds = append(allFoodIds, foodId)
	}

	var allRestaurantIds []int
	for restaurantId := range restaurantIdsMap {
		allRestaurantIds = append(allRestaurantIds, restaurantId)
	}

	// Fetch foods via RPC
	foodsMap := make(map[int]menudomain.MenuFood)
	if len(allFoodIds) > 0 {
		foods, err := h.foodRepo.FindFoodsByIds(ctx, allFoodIds)
		if err == nil {
			for _, food := range foods {
				foodsMap[food.Id] = food
			}
		}
	}

	// Fetch restaurants via RPC
	restaurantsMap := make(map[int]menudomain.MenuRestaurant)
	if len(allRestaurantIds) > 0 {
		restaurants, err := h.restaurantRepo.FindRestaurantsByIds(ctx, allRestaurantIds)
		if err == nil {
			for _, restaurant := range restaurants {
				restaurantsMap[restaurant.Id] = restaurant
			}
		}
	}

	// Populate menus with foods and restaurants
	for i := range menus {
		// Populate foods
		if len(menus[i].FoodIds) > 0 {
			var menuFoods []menudomain.MenuFood
			for _, foodId := range menus[i].FoodIds {
				if food, ok := foodsMap[foodId]; ok {
					menuFoods = append(menuFoods, food)
				}
			}
			if len(menuFoods) > 0 {
				menus[i].Foods = menuFoods
			}
		}

		// Populate restaurant
		if restaurant, ok := restaurantsMap[menus[i].RestaurantId]; ok {
			menus[i].Restaurant = &restaurant
		}
	}

	return &ListMenuResult{
		Data:   menus,
		Paging: query.DTO.Paging,
	}, nil
}

