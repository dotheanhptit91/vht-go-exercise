package categorygrpcserver

import (
	"context"
	"errors"
	categorygrpc "vht-go/gen/proto/category"
	categorydomain "vht-go/modules/category/domain"
	"vht-go/shared"

	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type IGetCategoryRepository interface {
	FindById(ctx context.Context, id uuid.UUID) (*categorydomain.Category, error)
}

type CategoryGrpcServer struct {
	repo IGetCategoryRepository
}

func NewCategoryGrpcServer(repo IGetCategoryRepository) *CategoryGrpcServer {
	return &CategoryGrpcServer{repo: repo}
}

func (s *CategoryGrpcServer) GetCategory(ctx context.Context, req *categorygrpc.GetCategoryRequest) (*categorygrpc.GetCategoryReply, error) {
	strId := req.Id

	if strId == "" {
		return nil, status.Errorf(codes.InvalidArgument, "id is required")
	}

	id, err := uuid.Parse(strId)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "invalid uuid format")
	}

	category, err := s.repo.FindById(ctx, id)
	if err != nil {
		if errors.Is(err, shared.ErrDataNotFound) {
			return nil, status.Errorf(codes.NotFound, "category not found")
		}

		return nil, status.Errorf(codes.Internal, "failed to get category")
	}

	grpcCategory := &categorygrpc.Category{
		Id:          category.Id.String(),
		Name:        category.Name,
		Description: category.Description,
		Status:      int32(category.Status),
	}

	return &categorygrpc.GetCategoryReply{Category: grpcCategory}, nil
}