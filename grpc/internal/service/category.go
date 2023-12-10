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

func (c *CategoryService) CreateCategory(ctx context.Context, input *pb.CreateCategoryRequest) (*pb.Category, error) {
	category, err := c.CategoryDb.Create(input.Name, input.Description)
	if err != nil {
		return nil, err
	}
	return &pb.Category{
		Id:          category.ID,
		Name:        category.Name,
		Description: category.Description,
	}, nil
}

func (c *CategoryService) ListCategories(ctx context.Context, input *pb.Blank) (*pb.CategoryList, error) {
	categories, err := c.CategoryDb.FindAll()
	if err != nil {
		return nil, err
	}
	var categoryList []*pb.Category
	for _, category := range categories {
		categoryList = append(categoryList, &pb.Category{
			Id:          category.ID,
			Name:        category.Name,
			Description: category.Description,
		})
	}
	return &pb.CategoryList{Categories: categoryList}, nil
}

func (c *CategoryService) GetCategory(ctx context.Context, input *pb.GetCategoryRequest) (*pb.Category, error) {
	category, err := c.CategoryDb.Find(input.Id)
	if err != nil {
		return nil, err
	}
	return &pb.Category{
		Id:          category.ID,
		Name:        category.Name,
		Description: category.Description,
	}, nil
}
