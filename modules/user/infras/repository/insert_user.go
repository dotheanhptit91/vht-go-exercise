package userrepository

import (
	"context"
	userdomain "vht-go/modules/user/domain"

	"github.com/pkg/errors"
)

func (repo *gormUserRepository) InsertUser(ctx context.Context, user *userdomain.User) error {
	if err := repo.db.Create(user).Error; err != nil {
		return errors.WithStack(err)
	}

	return nil
}