package main

import (
	"fmt"
	"net"

	"github.com/niemet0502/shirabe/users/database"
	"github.com/niemet0502/shirabe/users/repository"
	"github.com/niemet0502/shirabe/users/server"
	"github.com/niemet0502/shirabe/users/service"
	pb "github.com/niemet0502/shirabe/users/user"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	db := database.InitDb()

	repo := repository.NewUserRepository(db)

	svc := service.NewUserService(repo)

	gs := grpc.NewServer()

	c := server.NewServer(svc)

	pb.RegisterUserServiceServer(gs, c)

	reflection.Register(gs)

	l, err := net.Listen("tcp", fmt.Sprintf(":%d", 9004))

	if err != nil {
		fmt.Print("Unable to create listener")
	}

	gs.Serve(l)

}
