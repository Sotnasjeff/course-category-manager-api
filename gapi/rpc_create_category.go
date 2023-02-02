package gapi

import (
	"context"
	"io"

	"github.com/Sotnasjeff/course-category-manager-api/pb"
)

func (server *Server) CreateCategory(ctx context.Context, req *pb.CreateCategoryRequest) (*pb.CreateCategoryResponse, error) {
	arg, err := server.store.Create(req.GetName(), req.GetDescription())
	if err != nil {
		return nil, err
	}

	return &pb.CreateCategoryResponse{
		Category: &pb.Category{
			Id:          arg.ID,
			Name:        arg.Name,
			Description: arg.Description,
		},
	}, nil

}

func (server *Server) CreateCategoryStream(stream pb.CategoryService_CreateCategoryStreamServer) error {
	categories := &pb.ListCategoryResponse{}

	for {
		category, err := stream.Recv()

		if err == io.EOF {
			return stream.SendAndClose(categories)
		}
		if err != nil {
			return err
		}

		categoryResult, err := server.store.Create(category.GetName(), category.GetDescription())
		if err != nil {
			return err
		}

		categories.Categories = append(categories.Categories, &pb.Category{
			Id:          categoryResult.ID,
			Name:        categoryResult.Name,
			Description: categoryResult.Description,
		})
	}
}

func (server *Server) CreateCategoryBidirectionalStream(stream pb.CategoryService_CreateCategoryBidirectionalStreamServer) error {
	for {
		category, err := stream.Recv()

		if err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}

		categoryResult, err := server.store.Create(category.GetName(), category.GetDescription())
		if err != nil {
			return err
		}

		err = stream.Send(&pb.CreateCategoryResponse{
			Category: &pb.Category{
				Id:          categoryResult.ID,
				Name:        categoryResult.Name,
				Description: category.Description,
			},
		})

		if err != nil {
			return err
		}

	}
}
