package service

import (
	"io"
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


func (c * CategoryService) ListCategories(ctx context.Context, in *pb.Blank) (*pb.CategoryListCategories, error) {

	categorys, err := c.CategoryDb.FindAll();

	if err != nil {
		return nil, err
	}

	categories := []*pb.Category{}

	

	for _, category := range categorys {
		

		categories = append(categories, &pb.Category{
			Id:category.ID,
			Name:category.Name,
			Description:category.Description,
		})
	}



	return &pb.CategoryListCategories{
		Categories:categories,
	}, nil
}

func (c * CategoryService) SearchCategory(ctx context.Context, in *pb.Query) (*pb.CategoryListCategories, error) {

	
	categories, err := c.CategoryDb.Search(in.Search)

	if err != nil {
		return nil, err
	}


	categoriesR := []*pb.Category{}

	for _, category := range categories {

		categoriesR = append(categoriesR, &pb.Category{
			Id:category.ID,
			Name:category.Name,
			Description:category.Description,
		})
	}

	return &pb.CategoryListCategories{
		Categories:categoriesR,
	}, nil

}


func (c * CategoryService) CreateCategoryStream(stream pb.CategoryService_CreateCategoryStreamServer)  error {

	categories := &pb.CategoryListCategories{}

	for {
		category, err := stream.Recv();

		if err == io.EOF {
			return stream.SendAndClose(categories)
		}

		if err != nil {
			return err
		}

		categoryResult, err := c.CategoryDb.Create(category.Name, category.Description)

		if err != nil {
			return err
		}

		categories.Categories = append(categories.Categories, &pb.Category {
			Id: categoryResult.ID,
			Name: categoryResult.Name,
			Description: categoryResult.Description,
		})

	}

}

func (c * CategoryService) CreateCategoryStreamBiderectinal(stream pb.CategoryService_CreateCategoryStreamBiderectinalServer)  error {

	categories := &pb.CategoryListCategories{}

	for {
		category, err := stream.Recv();

		

		if err != nil {
			return err
		}

		categoryResult, err := c.CategoryDb.Create(category.Name, category.Description)

		if err != nil {
			return err
		}

		categories.Categories = append(categories.Categories,&pb.Category{
			Id:categoryResult.ID,
			Name:categoryResult.Name,
			Description:categoryResult.Description,
		})
		err = stream.Send(categories)
		if err != nil {
			return err
		}
	}

}