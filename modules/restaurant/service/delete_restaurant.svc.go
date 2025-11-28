package restaurantservice

import (
	"context"
	restaurantdomain "vht-go/modules/restaurant/domain"
)

type DeleteRestaurantCommand struct {
	Id int
}

type IDeleteRestaurantFindRepository interface {
	FindById(ctx context.Context, id int) (*restaurantdomain.Restaurant, error)
}

type IDeleteRestaurantRepository interface {
	Delete(ctx context.Context, id int) error
}

type DeleteRestaurantCommandHandler struct {
	findRepo   IDeleteRestaurantFindRepository
	deleteRepo IDeleteRestaurantRepository
}

func NewDeleteRestaurantCommandHandler(findRepo IDeleteRestaurantFindRepository, deleteRepo IDeleteRestaurantRepository) *DeleteRestaurantCommandHandler {
	return &DeleteRestaurantCommandHandler{
		findRepo:   findRepo,
		deleteRepo: deleteRepo,
	}
}

func (h *DeleteRestaurantCommandHandler) Handle(ctx context.Context, cmd *DeleteRestaurantCommand) error {
	// Check if restaurant exists
	_, err := h.findRepo.FindById(ctx, cmd.Id)
	if err != nil {
		return err
	}

	// Delete restaurant
	return h.deleteRepo.Delete(ctx, cmd.Id)
}

