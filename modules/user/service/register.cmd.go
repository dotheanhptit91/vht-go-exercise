package userservice

import (
	"context"
	"crypto/rand"
	"errors"
	"math/big"
	"time"
	userdomain "vht-go/modules/user/domain"
	userdto "vht-go/modules/user/dto"
	"vht-go/shared"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type RegisterUserCommand struct {
	DTO *userdto.RegisterUserDTO
}

type IRegisterUserRepository interface {
	GetUserByEmail(ctx context.Context, email string) (*userdomain.User, error)
	InsertUser(ctx context.Context, user *userdomain.User) error
}

type RegisterUserCommandHandler struct {
	repo IRegisterUserRepository
}

func NewRegisterUserCommandHandler(repo IRegisterUserRepository) *RegisterUserCommandHandler {
	return &RegisterUserCommandHandler{repo: repo}
}

func (h *RegisterUserCommandHandler) Handle(ctx context.Context, cmd *RegisterUserCommand) (*uuid.UUID, error) {
	dto := cmd.DTO

	if err := dto.Validate(); err != nil {
		return nil, shared.ErrBadRequest.WithError(err.Error())
	}

	user, err := h.repo.GetUserByEmail(ctx, dto.Email)
	if err != nil && !errors.Is(err, shared.ErrDataNotFound) {
		return nil, shared.ErrInternalServerError.WithWrap(err).WithDebug(err.Error())
	}

	if user != nil {
		if user.IsDeleted() {
			return nil, shared.ErrBadRequest.WithError(userdomain.ErrUserEmailAlreadyDeleted)
		}

		return nil, shared.ErrBadRequest.WithError(userdomain.ErrUserEmailAlreadyExists)
	}

	newId, err := uuid.NewV7()
	if err != nil {
		return nil, shared.ErrInternalServerError.WithWrap(err).WithDebug(err.Error())
	}

	salt := randomString(32)
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(salt+dto.Password), 10)
	if err != nil {
		return nil, shared.ErrInternalServerError.WithWrap(err).WithDebug(err.Error())
	}

	user = &userdomain.User{
		Id:        newId,
		Email:     dto.Email,
		Salt:      salt,
		Password:  string(hashedPassword),
		LastName:  dto.LastName,
		FirstName: dto.FirstName,
		Status:    1,
		Role:      shared.RoleUser,
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
	}

	if err := h.repo.InsertUser(ctx, user); err != nil {
		return nil, shared.ErrInternalServerError.WithWrap(err).
			WithDebug(err.Error()).WithError(userdomain.ErrCannotRegisterUser)
	}

	return &newId, nil
}

func randomString(length int) string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

	b := make([]byte, length)
	charsetLen := big.NewInt(int64(len(charset)))

	for i := range b {
		n, err := rand.Int(rand.Reader, charsetLen)
		if err != nil {
			panic(err)
		}
		b[i] = charset[n.Int64()%int64(len(charset))]
	}

	return string(b)
}