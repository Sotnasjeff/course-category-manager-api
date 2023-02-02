package gapi

import (
	"context"

	"github.com/Sotnasjeff/course-category-manager-api/pb"
)

func (server *Server) ListCategories(ctx context.Context, p *pb.Blank) (*pb.ListCategoryResponse, error) {
	categories, err := server.store.FindAll()
	if err != nil {
		return nil, err
	}

	var categoriesResponse []*pb.Category

	for _, category := range categories {
		categoryResponse := &pb.Category{
			Id:          category.ID,
			Name:        category.Name,
			Description: category.Description,
		}

		categoriesResponse = append(categoriesResponse, categoryResponse)
	}

	return &pb.ListCategoryResponse{
		Categories: categoriesResponse,
	}, nil

}
