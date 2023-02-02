package gapi

import (
	"github.com/Sotnasjeff/course-category-manager-api/db"
	"github.com/Sotnasjeff/course-category-manager-api/pb"
)

type Server struct {
	pb.UnimplementedCategoryServiceServer
	store db.Category
}

func NewServer(store db.Category) (*Server, error) {
	return &Server{
		store: store,
	}, nil
}
