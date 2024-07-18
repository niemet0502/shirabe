package server

import (
	"context"

	protos "github.com/niemet0502/shirabe/books/book"
	"github.com/niemet0502/shirabe/books/models"
	"github.com/niemet0502/shirabe/books/service"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type server struct {
	protos.UnimplementedBookServer
	svc *service.BookService
}

func mapBookToProtosBookEntity(book models.Book) *protos.BookEntity {
	return &protos.BookEntity{
		Id:              int32(book.ID),
		Title:           book.Title,
		Author:          book.Author,
		Genre:           book.Genre,
		Status:          int32(book.Status),
		TotalPages:      int32(book.TotalPages),
		ReadingProgress: int32(book.ReadingProgress),
		UserId:          int32(book.UserId),
	}
}

func mapProtoBookEntityToBook(proto *protos.BookEntity) models.Book {
	return models.Book{
		ID:         int(proto.Id),
		Title:      proto.Title,
		Author:     proto.Author,
		Genre:      proto.Genre,
		TotalPages: int(proto.TotalPages),
		UserId:     int(proto.UserId),
	}
}

func NewBook(svc *service.BookService) *server {
	return &server{svc: svc}
}

func (b *server) GetBook(ctx context.Context, rr *protos.GetBookRequest) (*protos.GetBookResponse, error) {
	result, err := b.svc.GetBook(int(rr.GetBookId()))

	if err != nil {
		return nil, status.Errorf(codes.NotFound, "Book not found")
	}
	return &protos.GetBookResponse{Book: mapBookToProtosBookEntity(result)}, nil
}

func (b *server) CreateBook(ctx context.Context, rr *protos.CreateBookRequest) (*protos.CreateBookResponse, error) {
	createBookDto := models.CreateBook{
		Title:      rr.GetTitle(),
		Author:     rr.GetAuthor(),
		Genre:      rr.GetGenre(),
		TotalPages: int(rr.GetTotalPages()),
		UserId:     int(rr.GetUserId()),
	}
	result := b.svc.CreateBook(createBookDto)

	return &protos.CreateBookResponse{Book: mapBookToProtosBookEntity(result)}, nil
}

// func (b *server) UpdateBook(ctx context.Context)

func (b *server) GetBooks(ctx context.Context, rr *protos.BooksRequest) (*protos.BooksResponse, error) {
	books := b.svc.GetBooks(int(rr.GetUserId()))

	var protoBooks []*protos.BookEntity
	for _, book := range books {
		protoBooks = append(protoBooks, mapBookToProtosBookEntity(book))
	}

	return &protos.BooksResponse{Books: protoBooks}, nil
}
