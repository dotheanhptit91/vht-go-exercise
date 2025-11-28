package menuservice

import (
	"context"
	menudomain "vht-go/modules/menu/domain"

	"github.com/google/uuid"
)

type DeleteMenuCommand struct {
	Id uuid.UUID
}

type IDeleteMenuFindRepository interface {
	FindById(ctx context.Context, id uuid.UUID) (*menudomain.Menu, error)
}

type IDeleteMenuRepository interface {
	SoftDelete(ctx context.Context, id uuid.UUID) error
}

type DeleteMenuCommandHandler struct {
	findRepo   IDeleteMenuFindRepository
	deleteRepo IDeleteMenuRepository
}

func NewDeleteMenuCommandHandler(
	findRepo IDeleteMenuFindRepository,
	deleteRepo IDeleteMenuRepository,
) *DeleteMenuCommandHandler {
	return &DeleteMenuCommandHandler{
		findRepo:   findRepo,
		deleteRepo: deleteRepo,
	}
}

func (h *DeleteMenuCommandHandler) Handle(ctx context.Context, cmd *DeleteMenuCommand) error {
	// Verify menu exists
	_, err := h.findRepo.FindById(ctx, cmd.Id)
	if err != nil {
		return err
	}

	// Soft delete (set status to 0)
	if err := h.deleteRepo.SoftDelete(ctx, cmd.Id); err != nil {
		return err
	}

	return nil
}

