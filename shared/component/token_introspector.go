package sharedcomponent

import (
	userdomain "vht-go/modules/user/domain"
	"vht-go/shared"

	"github.com/pkg/errors"
	"gorm.io/gorm"
)

type tokenIntrospector struct {
	jwtComponent *JwtComp
	db           *gorm.DB
}

func NewTokenIntrospector(jwtComponent *JwtComp, db *gorm.DB) *tokenIntrospector {
	return &tokenIntrospector{jwtComponent: jwtComponent, db: db}
}

func (t *tokenIntrospector) Introspect(token string) (shared.Requester, error) {
	subject, err := t.jwtComponent.Validate(token)

	if err != nil {
		return nil, err
	}

	var user userdomain.User

	if err := t.db.Where("id = ?", subject).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("user not found")
		}
		return nil, errors.WithStack(err)
	}

	if user.IsDeleted() {
		return nil, errors.New("user is deleted")
	}

	user.Mask()

	return &user, nil
}