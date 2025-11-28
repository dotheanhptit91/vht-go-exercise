package rstlikeservice

import (
	"context"
	"errors"
	"log"
	rstlikedomain "vht-go/modules/restaurantlike/domain"
	"vht-go/shared"
	"vht-go/shared/component/pubsub"

	"github.com/google/uuid"
)

type UnlikeRestaurantCommand struct {
	RestaurantId int              `json:"restaurantId"`
	Requester    shared.Requester `json:"-"`
}

type IUnlikeRestaurantRepository interface {
	GetRestaurantLike(ctx context.Context, restaurantId int, userId uuid.UUID) (*rstlikedomain.RestaurantLike, error)
	DeleteRestaurantLike(ctx context.Context, restaurantId int, userId uuid.UUID) error
}

// type IDecreaseLikedCountRepository interface {
// 	DecreaseLikedCount(ctx context.Context, restaurantId int) error
// }

type IEventPublisher interface {
	Publish(ctx context.Context, topic pubsub.Topic, data *pubsub.Message) error
}

type UnlikeRestaurantCommandHandler struct {
	repo IUnlikeRestaurantRepository
	ps   IEventPublisher
}

func NewUnlikeRestaurantCommandHandler(
	repo IUnlikeRestaurantRepository,
	// decreaseLikedCountRepo IDecreaseLikedCountRepository,
	ps IEventPublisher,
) *UnlikeRestaurantCommandHandler {
	return &UnlikeRestaurantCommandHandler{
		repo: repo,
		// decreaseLikedCountRepo: decreaseLikedCountRepo,
		ps: ps,
	}
}

func (h *UnlikeRestaurantCommandHandler) Handle(ctx context.Context, cmd *UnlikeRestaurantCommand) error {
	_, err := h.repo.GetRestaurantLike(ctx, cmd.RestaurantId, cmd.Requester.Subject())

	if err != nil {
		if errors.Is(err, shared.ErrDataNotFound) {
			return shared.ErrNotFound.WithWrap(err).WithError(rstlikedomain.ErrRestaurantHasNotLiked)
		}

		return shared.ErrInternalServerError.WithWrap(err).WithDebug(err.Error())
	}

	if err := h.repo.DeleteRestaurantLike(ctx, cmd.RestaurantId, cmd.Requester.Subject()); err != nil {
		return shared.ErrInternalServerError.WithWrap(err).WithDebug(err.Error())
	}

	// Event data and publish to pubsub
	evtData := rstlikedomain.RestaurantEvent{
		RestaurantId: cmd.RestaurantId,
		UserId:       cmd.Requester.Subject(),
	}
	
	if err := h.ps.Publish(ctx, pubsub.Topic(shared.EvtRestaurantUnliked), pubsub.NewMessage(evtData)); err != nil {
		log.Println("error publish unliked event", err)
	}

	return nil
}