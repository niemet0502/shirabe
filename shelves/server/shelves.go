package server

import (
	"context"
	"fmt"

	"github.com/niemet0502/shirabe/shelves/service"
	pb "github.com/niemet0502/shirabe/shelves/shelve"
)

type Server struct {
	pb.UnimplementedShelveServiceServer
	svc *service.ShelvesService
}

func NewServer(svc *service.ShelvesService) *Server {
	return &Server{svc: svc}
}

func (s *Server) GetShelves(ctx context.Context, rr *pb.GetShelvesRequest) (*pb.GetShelvesResponse, error) {
	var list []*pb.Shelf

	fmt.Print("fetching shelves")
	return &pb.GetShelvesResponse{Shelves: list}, nil
}
