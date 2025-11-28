package rstlikeservice

import (
	"context"
	"errors"
	"log"
	"time"
	rstlikedomain "vht-go/modules/restaurantlike/domain"
	"vht-go/shared"
	"vht-go/shared/component/pubsub"

	"github.com/google/uuid"
)

type LikeRestaurantCommand struct {
	RestaurantId int              `json:"restaurantId"`
	Requester    shared.Requester `json:"-"`
}

type IGetRestaurantRepository interface {
	GetRestaurant(ctx context.Context, id int) (*rstlikedomain.Restaurant, error)
}

type ILikeRestaurantRepository interface {
	GetRestaurantLike(ctx context.Context, restaurantId int, userId uuid.UUID) (*rstlikedomain.RestaurantLike, error)
	InsertRestaurantLike(ctx context.Context, restaurantLike *rstlikedomain.RestaurantLike) error
}

// type IIncreaseLikedCountRepository interface {
// 	IncreaseLikedCount(ctx context.Context, restaurantId int) error
// }

type LikeRestaurantCommandHandler struct {
	restaurantRepo IGetRestaurantRepository
	repo           ILikeRestaurantRepository
	ps             IEventPublisher
	// increaseLikedCountRepo IIncreaseLikedCountRepository
}

func NewLikeRestaurantCommandHandler(
	restaurantRepo IGetRestaurantRepository,
	repo ILikeRestaurantRepository,
	// increaseLikedCountRepo IIncreaseLikedCountRepository,
	ps IEventPublisher,
) *LikeRestaurantCommandHandler {
	return &LikeRestaurantCommandHandler{
		restaurantRepo: restaurantRepo,
		repo:           repo,
		// increaseLikedCountRepo: increaseLikedCountRepo,
		ps: ps,
	}
}

func (h *LikeRestaurantCommandHandler) Handle(ctx context.Context, cmd *LikeRestaurantCommand) error {
	restaurant, err := h.restaurantRepo.GetRestaurant(ctx, cmd.RestaurantId)

	if err != nil {
		if errors.Is(err, shared.ErrDataNotFound) {
			return shared.ErrNotFound.WithWrap(err).WithError(rstlikedomain.ErrRestaurantFound)
		}

		return shared.ErrInternalServerError.WithWrap(err).WithDebug(err.Error())
	}

	if restaurant.IsDeleted() {
		return shared.ErrNotFound.WithWrap(err).WithError(rstlikedomain.ErrRestaurantFound)
	}

	restaurantLike, err := h.repo.GetRestaurantLike(ctx, cmd.RestaurantId, cmd.Requester.Subject())

	if err != nil && !errors.Is(err, shared.ErrDataNotFound) {
		return shared.ErrInternalServerError.WithWrap(err).WithDebug(err.Error())
	}

	if restaurantLike != nil {
		return shared.ErrBadRequest.WithError(rstlikedomain.ErrRestaurantHasAlreadyLiked)
	}

	newRestaurantLike := &rstlikedomain.RestaurantLike{
		RestaurantId: cmd.RestaurantId,
		UserId:       cmd.Requester.Subject(),
		CreatedAt:    time.Now().UTC(),
	}

	if err := h.repo.InsertRestaurantLike(ctx, newRestaurantLike); err != nil {
		return shared.ErrInternalServerError.WithWrap(err).WithDebug(err.Error())
	}

	// Event data and publish to pubsub
	evtData := rstlikedomain.RestaurantEvent{
		RestaurantId: cmd.RestaurantId,
		UserId:       cmd.Requester.Subject(),
	}

	if err := h.ps.Publish(ctx, pubsub.Topic(shared.EvtRestaurantLiked), pubsub.NewMessage(evtData)); err != nil {
		log.Println("error publish liked event", err)
	}

	return nil
}