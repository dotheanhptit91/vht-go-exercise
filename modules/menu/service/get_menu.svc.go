package menuservice

import (
	"context"
	menudomain "vht-go/modules/menu/domain"

	"github.com/google/uuid"
)

type GetMenuQuery struct {
	Id uuid.UUID
}

type IGetMenuRepository interface {
	FindById(ctx context.Context, id uuid.UUID) (*menudomain.Menu, error)
}

type IGetMenuFoods interface {
	FindFoodsByIds(ctx context.Context, ids []int) ([]menudomain.MenuFood, error)
}

type IGetMenuRestaurant interface {
	FindRestaurantById(ctx context.Context, id int) (*menudomain.MenuRestaurant, error)
}

type GetMenuQueryHandler struct {
	repo           IGetMenuRepository
	foodRepo       IGetMenuFoods
	restaurantRepo IGetMenuRestaurant
}

func NewGetMenuQueryHandler(
	repo IGetMenuRepository,
	foodRepo IGetMenuFoods,
	restaurantRepo IGetMenuRestaurant,
) *GetMenuQueryHandler {
	return &GetMenuQueryHandler{
		repo:           repo,
		foodRepo:       foodRepo,
		restaurantRepo: restaurantRepo,
	}
}

func (h *GetMenuQueryHandler) Handle(ctx context.Context, query *GetMenuQuery) (*menudomain.Menu, error) {
	menu, err := h.repo.FindById(ctx, query.Id)
	if err != nil {
		return nil, err
	}

	// Populate foods via RPC
	if menu != nil && len(menu.FoodIds) > 0 {
		foods, err := h.foodRepo.FindFoodsByIds(ctx, menu.FoodIds)
		if err == nil && len(foods) > 0 {
			menu.Foods = foods
		}
	}

	// Populate restaurant via RPC
	if menu != nil {
		restaurant, err := h.restaurantRepo.FindRestaurantById(ctx, menu.RestaurantId)
		if err == nil {
			menu.Restaurant = restaurant
		}
	}

	return menu, nil
}

