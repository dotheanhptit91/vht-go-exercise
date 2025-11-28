package restaurantservice

import (
	"context"
	"log"
	"time"
	restaurantdomain "vht-go/modules/restaurant/domain"
	restaurantdtos "vht-go/modules/restaurant/dtos"

	"github.com/google/uuid"
)

type CreateRestaurantResultCommand struct {
	DTO *restaurantdtos.CreateRestaurantDTO
}

type ICreateRestaurantRepository interface {
	Insert(ctx context.Context, restaurant *restaurantdomain.Restaurant) error
}

type CreateRestaurantResultCommandHandler struct {
	repo ICreateRestaurantRepository
}

func NewCreateRestaurantResultCommandHandler(repo ICreateRestaurantRepository) *CreateRestaurantResultCommandHandler {
	return &CreateRestaurantResultCommandHandler{repo: repo}
}

func (h *CreateRestaurantResultCommandHandler) Handle(ctx context.Context, cmd *CreateRestaurantResultCommand) (int, error) {
	if err := cmd.DTO.Validate(); err != nil {
		return 0, err
	}

    var catId uuid.UUID

	if cmd.DTO.CategoryId != nil {
		var err error
		catId, err = uuid.Parse(*cmd.DTO.CategoryId)
		if err != nil {
			log.Println(err)
		}
	}
	currentTime := time.Now().UTC()

	restaurant := &restaurantdomain.Restaurant{
		OwnerId:          cmd.DTO.OwnerId,
		CategoryId:       &catId,
		Name:             cmd.DTO.Name,
		Addr:             cmd.DTO.Addr,
		CityId:           cmd.DTO.CityId,
		Lat:              cmd.DTO.Lat,
		Lng:              cmd.DTO.Lng,
		ShippingFeePerKm: cmd.DTO.ShippingFeePerKm,
		Status:           cmd.DTO.Status,
		CreatedAt:        &currentTime,
		UpdatedAt:        &currentTime,
	}

	if err := h.repo.Insert(ctx, restaurant); err != nil {
		return 0, err
	}

	return restaurant.Id, nil
}

