package restaurantservice

import (
	"context"
	restaurantdomain "vht-go/modules/restaurant/domain"
	restaurantdtos "vht-go/modules/restaurant/dtos"

	"github.com/google/uuid"
)

type ListRestaurantQuery struct {
	DTO restaurantdtos.ListRestaurantDTO
}

type ListRestaurantResult struct {
	Data   []restaurantdomain.Restaurant `json:"data"`
	Paging struct {
		Page  int   `json:"page"`
		Limit int   `json:"limit"`
		Total int64 `json:"total"`
	} `json:"paging"`
}

type IListRestaurantRepository interface {
	FindAll(ctx context.Context, offset, limit int) ([]restaurantdomain.Restaurant, error)
	Count(ctx context.Context) (int64, error)
	FindWithFilters(ctx context.Context, ownerId *int, cityId *int, status *int, offset, limit int) ([]restaurantdomain.Restaurant, error)
	CountWithFilters(ctx context.Context, ownerId *int, cityId *int, status *int) (int64, error)
}

type IGetRestaurantsCategories interface {
	FindCategoriesByIds(ctx context.Context, ids []uuid.UUID) ([]restaurantdomain.RestaurantCategory, error)
}

type ListRestaurantQueryHandler struct {
	repo IListRestaurantRepository
	catRepo IGetRestaurantsCategories
}

func NewListRestaurantQueryHandler(repo IListRestaurantRepository, catRepo IGetRestaurantsCategories) *ListRestaurantQueryHandler {
	return &ListRestaurantQueryHandler{repo: repo, catRepo: catRepo}
}

func (h *ListRestaurantQueryHandler) Handle(ctx context.Context, query *ListRestaurantQuery) (*ListRestaurantResult, error) {
	// Process paging defaults
	query.DTO.Paging.Process()

	// Check if filters are applied
	hasFilters := query.DTO.OwnerId != nil || query.DTO.CityId != nil || query.DTO.Status != nil

	var restaurants []restaurantdomain.Restaurant
	var total int64
	var err error

	if hasFilters {
		// Query with filters
		restaurants, err = h.repo.FindWithFilters(
			ctx,
			query.DTO.OwnerId,
			query.DTO.CityId,
			query.DTO.Status,
			(query.DTO.Paging.Page-1)*query.DTO.Paging.Limit,
			query.DTO.Paging.Limit,
		)
		if err != nil {
			return nil, err
		}

		total, err = h.repo.CountWithFilters(ctx, query.DTO.OwnerId, query.DTO.CityId, query.DTO.Status)
		if err != nil {
			return nil, err
		}
	} else {
		// Query all
		restaurants, err = h.repo.FindAll(
			ctx,
			(query.DTO.Paging.Page-1)*query.DTO.Paging.Limit,
			query.DTO.Paging.Limit,
		)
		if err != nil {
			return nil, err
		}

		total, err = h.repo.Count(ctx)
		if err != nil {
			return nil, err
		}
	}

	catIds := make([]uuid.UUID, 0)

	for _, restaurant := range restaurants {
		if restaurant.CategoryId != nil {
			catIds = append(catIds, *restaurant.CategoryId)
		}
	}

	categories, err := h.catRepo.FindCategoriesByIds(ctx, catIds)
	if err != nil {
		return nil, err
	}

	catMap := make(map[uuid.UUID]restaurantdomain.RestaurantCategory)

	for _, category := range categories {
		catMap[category.Id] = category
	}

	for i, restaurant := range restaurants {
		if restaurant.CategoryId != nil {
			if category, exists := catMap[*restaurant.CategoryId]; exists {
				restaurants[i].Category = &category
			}
		}
	}

	result := &ListRestaurantResult{
		Data: restaurants,
	}
	result.Paging.Page = query.DTO.Paging.Page
	result.Paging.Limit = query.DTO.Paging.Limit
	result.Paging.Total = total

	return result, nil
}

