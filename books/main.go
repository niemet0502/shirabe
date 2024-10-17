package main

import (
	"fmt"
	"net"

	"github.com/niemet0502/shirabe/books/database"
	"github.com/niemet0502/shirabe/books/repository"
	"github.com/niemet0502/shirabe/books/service"

	pb "github.com/niemet0502/shirabe/books/book"
	p "github.com/niemet0502/shirabe/books/readingprogress"
	"github.com/niemet0502/shirabe/books/server"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	fmt.Println("Books service")

	// load env variable
	db := database.InitDb()

	// create the repository
	repo := repository.NewBookRepository(db)

	// create the service
	svc := service.NewBookService(repo)

	// create the server and start listening
	gs := grpc.NewServer()

	c := server.NewBook(svc)

	pb.RegisterBookServer(gs, c)

	// handle reading progress

	progressRepo := repository.NewReadingProgressRepository(db)

	progressSvc := service.NewReadingProgressService(progressRepo, svc)

	progressService := server.NewReadingProgressServer(progressSvc)

	p.RegisterReadingServiceServer(gs, progressService)

	reflection.Register(gs)

	l, err := net.Listen("tcp", fmt.Sprintf(":%d", 9002))

	if err != nil {
		fmt.Print("Unable to create listener")
	}

	gs.Serve(l)
}
