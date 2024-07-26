package server

import (
	"context"

	"github.com/niemet0502/shirabe/shelves/models"
	"github.com/niemet0502/shirabe/shelves/service"
	pb "github.com/niemet0502/shirabe/shelves/shelve"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Server struct {
	pb.UnimplementedShelveServiceServer
	svc *service.ShelvesService
}

func mapShelfToShelfProto(shelf models.Shelf) *pb.Shelf {
	return &pb.Shelf{
		Id:     int32(shelf.ID),
		Name:   shelf.Name,
		UserId: shelf.UserId,
	}
}

func NewServer(svc *service.ShelvesService) *Server {
	return &Server{svc: svc}
}

func (s *Server) GetShelves(ctx context.Context, rr *pb.GetShelvesRequest) (*pb.GetShelvesResponse, error) {
	var list []*pb.Shelf

	result := s.svc.GetShelvesByUser(int(rr.GetUserId()))

	for _, s := range result {
		list = append(list, mapShelfToShelfProto(s))
	}
	return &pb.GetShelvesResponse{Shelves: list}, nil
}

func (s *Server) CreateShelf(ctx context.Context, rr *pb.CreateShelfRequest) (*pb.CreateShelfResponse, error) {
	result := s.svc.CreateShelf(rr.GetName(), int(rr.GetUserId()))
	return &pb.CreateShelfResponse{Shelf: mapShelfToShelfProto(result)}, nil
}

func (s *Server) RemoveShelf(ctx context.Context, rr *pb.RemoveShelfRequest) (*pb.RemoveShelfResponse, error) {

	if err := s.svc.RemoveShelf(int(rr.GetShelfId())); err != nil {
		return &pb.RemoveShelfResponse{}, status.Errorf(codes.NotFound, "Failed to remove the shelf")
	}
	return &pb.RemoveShelfResponse{Message: "Shelf successfully removed"}, nil
}
