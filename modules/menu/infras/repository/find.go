package menurepository

import (
	"context"
	"errors"
	menudomain "vht-go/modules/menu/domain"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

func (repo *GORMMenuRepository) FindById(ctx context.Context, id uuid.UUID) (*menudomain.Menu, error) {
	var menu menudomain.Menu
	
	if err := repo.db.WithContext(ctx).Where("id = ?", id).First(&menu).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New(menudomain.ErrMenuNotFound)
		}
		return nil, err
	}
	
	return &menu, nil
}

func (repo *GORMMenuRepository) FindAll(ctx context.Context, filter map[string]interface{}, limit, offset int) ([]menudomain.Menu, error) {
	var menus []menudomain.Menu
	
	query := repo.db.WithContext(ctx)
	
	// Apply filters
	if restaurantId, ok := filter["restaurant_id"].(int); ok {
		query = query.Where("restaurant_id = ?", restaurantId)
	}
	
	if status, ok := filter["status"].(int); ok {
		query = query.Where("status = ?", status)
	}
	
	if err := query.Limit(limit).Offset(offset).Order("created_at DESC").Find(&menus).Error; err != nil {
		return nil, err
	}
	
	return menus, nil
}

func (repo *GORMMenuRepository) Count(ctx context.Context, filter map[string]interface{}) (int64, error) {
	var count int64
	
	query := repo.db.WithContext(ctx).Model(&menudomain.Menu{})
	
	// Apply filters
	if restaurantId, ok := filter["restaurant_id"].(int); ok {
		query = query.Where("restaurant_id = ?", restaurantId)
	}
	
	if status, ok := filter["status"].(int); ok {
		query = query.Where("status = ?", status)
	}
	
	if err := query.Count(&count).Error; err != nil {
		return 0, err
	}
	
	return count, nil
}

