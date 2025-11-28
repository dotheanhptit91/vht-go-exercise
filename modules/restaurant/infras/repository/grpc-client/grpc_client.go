package restaurantgrpcclient

import (
	"context"
	"log"
	categorygrpc "vht-go/gen/proto/category"
	restaurantdomain "vht-go/modules/restaurant/domain"

	"github.com/google/uuid"
	"github.com/pkg/errors"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type RestaurantGrpcClient struct {
	categoryGrpcClient categorygrpc.CategoryServiceClient
}

func NewRestaurantGrpcClient(uri string) *RestaurantGrpcClient {
	conn, err := grpc.NewClient(uri, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalln("Failed to connect to gRPC server:", err)
	}

	return &RestaurantGrpcClient{categoryGrpcClient: categorygrpc.NewCategoryServiceClient(conn)}
}

func (c *RestaurantGrpcClient) FindCategoryById(ctx context.Context, id uuid.UUID) (*restaurantdomain.RestaurantCategory, error) {
	categoryReply, err := c.categoryGrpcClient.GetCategory(ctx, &categorygrpc.GetCategoryRequest{Id: id.String()})

	if err != nil {
		return nil, errors.WithStack(err)
	}

	category := &restaurantdomain.RestaurantCategory{
		Id:   uuid.MustParse(categoryReply.Category.Id),
		Name: categoryReply.Category.Name,
	}

	return category, nil
}