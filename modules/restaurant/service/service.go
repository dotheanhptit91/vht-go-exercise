package restaurantservice

import (
	"context"
	restaurantdomain "vht-go/modules/restaurant/domain"
	restaurantdtos "vht-go/modules/restaurant/dtos"
)

type IRestaurantRepository interface {
	Insert(ctx context.Context, restaurant *restaurantdomain.Restaurant) error
	FindById(ctx context.Context, id int) (*restaurantdomain.Restaurant, error)
	FindAll(ctx context.Context, offset, limit int) ([]restaurantdomain.Restaurant, error)
	Count(ctx context.Context) (int64, error)
	FindWithFilters(ctx context.Context, ownerId *int, cityId *int, status *int, offset, limit int) ([]restaurantdomain.Restaurant, error)
	CountWithFilters(ctx context.Context, ownerId *int, cityId *int, status *int) (int64, error)
	Update(ctx context.Context, restaurant *restaurantdomain.Restaurant, id int) error
	Delete(ctx context.Context, id int) error
}

// Deprecated: Use CQRS handlers instead
type IRestaurantService interface {
	CreateNewRestaurant(ctx context.Context, dto *restaurantdtos.CreateRestaurantDTO) (newId int, err error)
	GetRestaurantById(ctx context.Context, id int) (*restaurantdomain.Restaurant, error)
	UpdateRestaurant(ctx context.Context, cmd *UpdateRestaurantCommand) error
	ListRestaurants(ctx context.Context, query *ListRestaurantQuery) (*ListRestaurantResult, error)
	DeleteRestaurant(ctx context.Context, id int) error
}
