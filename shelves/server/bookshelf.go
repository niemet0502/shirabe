package server

import (
	"context"

	"github.com/niemet0502/shirabe/shelves/models"
	"github.com/niemet0502/shirabe/shelves/service"
	pb "github.com/niemet0502/shirabe/shelves/shelve"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type BookShelfServer struct {
	pb.UnimplementedBookShelfServiceServer

	bookSvc *service.BookShelfService
}

func mapModelToProto(bookShelf models.BookShelf) *pb.BookShelf {
	return &pb.BookShelf{Id: int32(bookShelf.ID), BookId: bookShelf.BookId, ShelfId: bookShelf.ShelfId}
}

func NewBookShelfServer(bookSvc *service.BookShelfService) *BookShelfServer {
	return &BookShelfServer{bookSvc: bookSvc}
}

func (s *BookShelfServer) AddBook(ctx context.Context, rr *pb.CreateBookShelf) (*pb.BookShelfResponse, error) {

	result, err := s.bookSvc.AddBookToShelf(int(rr.GetBookId()), int(rr.GetShelfId()))

	if err != nil {
		return &pb.BookShelfResponse{}, status.Errorf(codes.Canceled, "failed to add the book")
	}

	return &pb.BookShelfResponse{Bookshelf: mapModelToProto(result)}, nil
}

func (s *BookShelfServer) SearchBookShelf(ctx context.Context, rr *pb.BookShelf) (*pb.BookShelvesResponse, error) {
	result := s.bookSvc.GetShelfBooks(int(rr.GetBookId()), int(rr.GetShelfId()))

	var protos []*pb.BookShelf

	for _, b := range result {
		protos = append(protos, mapModelToProto(b))
	}

	return &pb.BookShelvesResponse{BookShelves: protos}, nil
}
