package server

import (
	"context"

	protos "github.com/niemet0502/shirabe/books/book"
	"github.com/niemet0502/shirabe/books/service"
)

type server struct {
	protos.UnimplementedBookServer
	svc *service.BookService
}

func NewBook(svc *service.BookService) *server {
	return &server{svc: svc}
}

func (b *server) GetBook(ctx context.Context, rr *protos.GetBookRequest) (*protos.GetBookResponse, error) {
	b.svc.GetBook()
	return &protos.GetBookResponse{Msg: "Testing"}, nil
}
