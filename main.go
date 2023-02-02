package main

import (
	"database/sql"
	"log"
	"net"

	"github.com/Sotnasjeff/course-category-manager-api/db"
	"github.com/Sotnasjeff/course-category-manager-api/gapi"
	"github.com/Sotnasjeff/course-category-manager-api/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	database, err := sql.Open("sqlite3", "./db.sqlite")
	if err != nil {
		log.Fatal("Couldn't connect to db")
	}
	defer database.Close()

	categoryDB := db.NewCategory(database)
	runGRPCServer(*categoryDB)
}

func runGRPCServer(store db.Category) {
	server, err := gapi.NewServer(store)
	if err != nil {
		log.Fatal("Cannot create server")
	}

	grpcServer := grpc.NewServer()
	pb.RegisterCategoryServiceServer(grpcServer, server)
	reflection.Register(grpcServer)

	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatal("Couldn't connect to the port")
	}

	err = grpcServer.Serve(lis)

	if err != nil {
		log.Fatal("Couldn't connect to the grpc Server")
	}

}
