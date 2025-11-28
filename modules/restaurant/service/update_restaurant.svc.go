package restaurantservice

import (
	"context"
	"strings"
	"time"
	restaurantdomain "vht-go/modules/restaurant/domain"
	restaurantdtos "vht-go/modules/restaurant/dtos"
)

type UpdateRestaurantCommand struct {
	Id   int
	Data restaurantdtos.UpdateRestaurantDTO
}

type IUpdateRestaurantRepository interface {
	FindById(ctx context.Context, id int) (*restaurantdomain.Restaurant, error)
	Update(ctx context.Context, restaurant *restaurantdomain.Restaurant, id int) error
}

type UpdateRestaurantCommandHandler struct {
	repo IUpdateRestaurantRepository
}

func NewUpdateRestaurantCommandHandler(repo IUpdateRestaurantRepository) *UpdateRestaurantCommandHandler {
	return &UpdateRestaurantCommandHandler{repo: repo}
}

func (h *UpdateRestaurantCommandHandler) Handle(ctx context.Context, cmd *UpdateRestaurantCommand) error {
	// Fetch existing restaurant
	oldRestaurant, err := h.repo.FindById(ctx, cmd.Id)
	if err != nil {
		return err
	}

	currentTime := time.Now().UTC()

	// Build updated entity with existing values
	restaurant := &restaurantdomain.Restaurant{
		Id:               cmd.Id,
		OwnerId:          oldRestaurant.OwnerId,
		Name:             oldRestaurant.Name,
		Addr:             oldRestaurant.Addr,
		CityId:           oldRestaurant.CityId,
		Lat:              oldRestaurant.Lat,
		Lng:              oldRestaurant.Lng,
		ShippingFeePerKm: oldRestaurant.ShippingFeePerKm,
		Status:           oldRestaurant.Status,
		CreatedAt:        oldRestaurant.CreatedAt,
		UpdatedAt:        &currentTime,
	}

	// Apply changes with validation
	if cmd.Data.Name != nil {
		name := strings.TrimSpace(*cmd.Data.Name)
		if name != "" {
			restaurant.Name = name
		}
	}

	if cmd.Data.Addr != nil {
		addr := strings.TrimSpace(*cmd.Data.Addr)
		if addr != "" {
			restaurant.Addr = addr
		}
	}

	if cmd.Data.CityId != nil {
		restaurant.CityId = cmd.Data.CityId
	}

	if cmd.Data.Lat != nil {
		restaurant.Lat = cmd.Data.Lat
	}

	if cmd.Data.Lng != nil {
		restaurant.Lng = cmd.Data.Lng
	}

	if cmd.Data.ShippingFeePerKm != nil {
		if *cmd.Data.ShippingFeePerKm >= 0 {
			restaurant.ShippingFeePerKm = *cmd.Data.ShippingFeePerKm
		}
	}

	if cmd.Data.Status != nil {
		restaurant.Status = *cmd.Data.Status
	}

	// Persist changes
	return h.repo.Update(ctx, restaurant, cmd.Id)
}

