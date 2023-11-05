package service

import (
	"context"

	"github.com/DevPio/grpc/internal/database"
	"github.com/DevPio/grpc/internal/pb"
)

type CategoryService struct {

	pb.UnimplementedCategoryServiceServer
	CategoryDb database.Category
}


func  NewCategoryService(categoryDb database.Category) (*CategoryService) {

	return &CategoryService{ 
		CategoryDb: categoryDb,
	}
}

func (c * CategoryService) CreateCategory(ctx context.Context, in *pb.CategoryRequest) (*pb.CategoryResponse, error) {

	category, err := c.CategoryDb.Create(in.Name, in.Description);

	if err != nil {
		return nil, err
	}


	categoryResponse := &pb.CategoryResponse{
		Category: &pb.Category{
			Id: category.ID,
			Name: category.Name,
			Description: category.Description,
		},

	}


	return categoryResponse, nil
}