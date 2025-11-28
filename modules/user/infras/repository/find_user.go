package userrepository

import (
	"context"
	userdomain "vht-go/modules/user/domain"
	"vht-go/shared"

	"github.com/pkg/errors"
	"gorm.io/gorm"
)

func (repo *gormUserRepository) GetUserByEmail(ctx context.Context, email string) (*userdomain.User, error) {
	var user userdomain.User

	if err := repo.db.Where("email = ?", email).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, shared.ErrDataNotFound
		}
		return nil, errors.WithStack(err)
	}

	return &user, nil
}