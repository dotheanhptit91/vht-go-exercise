package userservice

import (
	"context"
	"errors"
	userdomain "vht-go/modules/user/domain"
	userdto "vht-go/modules/user/dto"
	"vht-go/shared"

	"golang.org/x/crypto/bcrypt"
)

type LoginUserCommand struct {
	DTO *userdto.LoginUserDTO
}

type ILoginUserRepository interface {
	GetUserByEmail(ctx context.Context, email string) (*userdomain.User, error)
}

type IJWTComponent interface {
	IssueToken(ctx context.Context, userID string) (string, error)
}

type LoginUserCommandHandler struct {
	repo         ILoginUserRepository
	jwtComponent IJWTComponent
}

func NewLoginUserCommandHandler(repo ILoginUserRepository, jwtComponent IJWTComponent) *LoginUserCommandHandler {
	return &LoginUserCommandHandler{repo: repo, jwtComponent: jwtComponent}
}

func (h *LoginUserCommandHandler) Handle(ctx context.Context, cmd *LoginUserCommand) (*userdto.LoginResponseDTO, error) {
	dto := cmd.DTO

	if err := dto.Validate(); err != nil {
		return nil, shared.ErrBadRequest.WithError(err.Error())
	}

	user, err := h.repo.GetUserByEmail(ctx, dto.Email)

	if err != nil {
		if errors.Is(err, shared.ErrDataNotFound) {
			return nil, shared.ErrBadRequest.WithError(userdomain.ErrUserNotFound)
		}
		return nil, shared.ErrInternalServerError.WithWrap(err).WithDebug(err.Error())
	}

	if user.IsDeleted() {
		return nil, shared.ErrBadRequest.WithError(userdomain.ErrUserEmailAlreadyDeleted)
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(user.Salt+dto.Password)); err != nil {
		return nil, shared.ErrBadRequest.WithError(userdomain.ErrUserEmailAndPasswordInvalid)
	}

	token, err := h.jwtComponent.IssueToken(ctx, user.Id.String())
	if err != nil {
		return nil, shared.ErrInternalServerError.WithWrap(err).WithDebug(err.Error())
	}

	return &userdto.LoginResponseDTO{Token: token}, nil
}