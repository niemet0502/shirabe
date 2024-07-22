package main

import (
	"fmt"
	"net"

	"github.com/niemet0502/shirabe/shelves/database"
	"github.com/niemet0502/shirabe/shelves/repository"
	"github.com/niemet0502/shirabe/shelves/server"
	"github.com/niemet0502/shirabe/shelves/service"
	pb "github.com/niemet0502/shirabe/shelves/shelve"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	// init the db
	db := database.InitDb()

	// create repository
	repo := repository.NewShelvesRepository(db)

	// create service
	svc := service.NewShelvesService(repo)

	// create grpc server
	gs := grpc.NewServer()

	// create new Shelves grpc server
	c := server.NewServer(svc)

	pb.RegisterShelveServiceServer(gs, c)

	reflection.Register(gs)

	l, err := net.Listen("tcp", fmt.Sprintf(":%d", 9003))

	if err != nil {
		fmt.Print("Unable to create listener")
	}

	gs.Serve(l)

}
