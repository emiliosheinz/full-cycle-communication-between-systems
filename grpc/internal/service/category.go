package service

import (
	"context"

	"github.com/emiliosheinz/full-cycle-cbs-grpc/internal/database"
	"github.com/emiliosheinz/full-cycle-cbs-grpc/internal/pb"
)

type CategoryService struct {
	pb.UnimplementedCategoryServiceServer
	CategoryDb database.Category
}

func NewCategoryService(categoryDb database.Category) *CategoryService {
	return &CategoryService{CategoryDb: categoryDb}
}

func (c *CategoryService) CreateCategory(ctx context.Context, input *pb.CreateCategoryRequest) (*pb.CategoryResponse, error) {
	category, err := c.CategoryDb.Create(input.Name, input.Description)
	if err != nil {
		return nil, err
	}
	responseCategory := &pb.Category{
		Id:          category.ID,
		Name:        category.Name,
		Description: category.Description,
	}
	return &pb.CategoryResponse{
		Category: responseCategory,
	}, nil
}
