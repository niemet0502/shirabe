package server

import (
	"context"
	"time"

	protos "github.com/niemet0502/shirabe/books/book"
	"github.com/niemet0502/shirabe/books/models"
	"github.com/niemet0502/shirabe/books/service"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type bookServer struct {
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
		Cover:           book.Cover,
		Description:     book.Description,
		PublicatedAt:    book.PublicatedAt.String(),
		StartReadingAt:  book.StartReadingAt.String(),
		FinishReadingAt: book.FinishReadingAt.String(),
	}
}

func mapProtosToBookModel(pb *protos.BookEntity) models.Book {
	book := models.Book{
		ID:              uint(pb.GetId()),
		Title:           pb.GetTitle(),
		Author:          pb.GetAuthor(),
		Genre:           pb.GetGenre(),
		Status:          int(pb.GetStatus()),
		TotalPages:      int(pb.GetTotalPages()),
		ReadingProgress: int(pb.GetReadingProgress()),
		UserId:          int(pb.GetUserId()),
		Cover:           pb.GetCover(),
		Description:     pb.GetDescription(),
	}
	book.PublicatedAt, _ = time.Parse("02-01-2006", pb.GetPublicatedAt())
	book.StartReadingAt, _ = time.Parse("02-01-2006", pb.GetStartReadingAt())
	book.FinishReadingAt, _ = time.Parse("02-01-2006", pb.GetFinishReadingAt())

	return book
}

func NewBook(svc *service.BookService) *bookServer {
	return &bookServer{svc: svc}
}

func (b *bookServer) GetBook(ctx context.Context, rr *protos.GetBookRequest) (*protos.GetBookResponse, error) {
	result, err := b.svc.GetBook(int(rr.GetBookId()))

	if err != nil {
		return nil, status.Errorf(codes.NotFound, "Book not found")
	}
	return &protos.GetBookResponse{Book: mapBookToProtosBookEntity(result)}, nil
}

func (b *bookServer) CreateBook(ctx context.Context, rr *protos.CreateBookRequest) (*protos.CreateBookResponse, error) {
	createBookDto := models.CreateBook{
		Title:       rr.GetTitle(),
		Author:      rr.GetAuthor(),
		Genre:       rr.GetGenre(),
		TotalPages:  int(rr.GetTotalPages()),
		UserId:      int(rr.GetUserId()),
		Cover:       rr.GetCover(),
		Description: rr.GetDescription(),
	}

	createBookDto.PublicatedAt, _ = time.Parse("02-01-2006", rr.GetPublicatedAt())

	result := b.svc.CreateBook(createBookDto)

	return &protos.CreateBookResponse{Book: mapBookToProtosBookEntity(result)}, nil
}

func (b *bookServer) GetBooks(ctx context.Context, rr *protos.BooksRequest) (*protos.BooksResponse, error) {
	books := b.svc.GetBooks(int(rr.GetUserId()))

	var protoBooks []*protos.BookEntity
	for _, book := range books {
		protoBooks = append(protoBooks, mapBookToProtosBookEntity(book))
	}

	return &protos.BooksResponse{Books: protoBooks}, nil
}

func (b *bookServer) UpdateBook(ctx context.Context, rr *protos.BookEntity) (*protos.CreateBookResponse, error) {
	result := b.svc.UpdateBook(mapProtosToBookModel(rr))

	return &protos.CreateBookResponse{Book: mapBookToProtosBookEntity(result)}, nil
}

func (b *bookServer) SearchBooks(ctx context.Context, rr *protos.SearchBooksRequest) (*protos.BooksResponse, error) {

	books := b.svc.SearchBooks(int(rr.GetUserId()), int(rr.GetStatus()), rr.GetSearch())

	var protoBooks []*protos.BookEntity
	for _, book := range books {
		protoBooks = append(protoBooks, mapBookToProtosBookEntity(book))
	}

	return &protos.BooksResponse{Books: protoBooks}, nil
}

func (b *bookServer) GetBooksByShelf(ctx context.Context, rr *protos.GetBooksByShelfRequest) (*protos.BooksResponse, error) {

	books := b.svc.GetBooksByShelf(rr.GetIds())

	var protoBooks []*protos.BookEntity
	for _, book := range books {
		protoBooks = append(protoBooks, mapBookToProtosBookEntity(book))
	}

	return &protos.BooksResponse{Books: protoBooks}, nil
}
