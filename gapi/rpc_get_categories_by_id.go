package gapi

import (
	"context"

	"github.com/Sotnasjeff/course-category-manager-api/pb"
)

func (server *Server) GetCategoriesById(ctx context.Context, req *pb.GetCategoriesByIdRequest) (*pb.GetCategoriesByIdResponse, error) {
	arg, err := server.store.Find(req.GetId())
	if err != nil {
		return nil, err
	}

	return &pb.GetCategoriesByIdResponse{
		Category: &pb.Category{
			Id:          arg.ID,
			Name:        arg.Name,
			Description: arg.Description,
		},
	}, nil
}
