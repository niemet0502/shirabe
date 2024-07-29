package services

import (
	"context"
	"errors"

	"github.com/niemet0502/shirabe/bff/models"
	pb "github.com/niemet0502/shirabe/books/book"
)

type BookService struct {
	client pb.BookClient
}

func mapProtosToBookModel(pb *pb.BookEntity) models.Book {
	return models.Book{
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
		PublicatedAt:    pb.GetPublicatedAt(),
		StartReadingAt:  pb.GetStartReadingAt(),
		FinishReadingAt: pb.GetFinishReadingAt(),
	}

}

func mapBookModelToProto(book models.Book) *pb.BookEntity {
	return &pb.BookEntity{
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
		PublicatedAt:    book.PublicatedAt,
		StartReadingAt:  book.StartReadingAt,
		FinishReadingAt: book.FinishReadingAt,
	}
}

func NewBookService(client pb.BookClient) *BookService {
	return &BookService{client}
}

func (s *BookService) GetBook(id int) (models.Book, error) {

	rr := &pb.GetBookRequest{
		BookId: int32(id),
	}

	resp, err := s.client.GetBook(context.Background(), rr)

	if err != nil {
		return models.Book{}, err
	}

	return mapProtosToBookModel(resp.GetBook()), nil
}

func (s *BookService) GetBooks(userId int) ([]models.Book, error) {

	rr := &pb.BooksRequest{
		UserId: int32(userId),
	}

	resp, err := s.client.GetBooks(context.Background(), rr)

	if err != nil {
		return []models.Book{}, errors.New("failed to fetch books")
	}

	var books []models.Book

	for _, b := range resp.GetBooks() {

		books = append(books, mapProtosToBookModel(b))
	}

	return books, nil

}

func (s *BookService) BooksSearch(userId, status int, search string) ([]models.Book, error) {

	rr := &pb.SearchBooksRequest{
		UserId: int32(userId),
		Status: int32(status),
		Search: search,
	}

	resp, err := s.client.SearchBooks(context.Background(), rr)

	if err != nil {
		return []models.Book{}, errors.New("failed to fetch books")
	}

	var books []models.Book

	for _, b := range resp.GetBooks() {

		books = append(books, mapProtosToBookModel(b))
	}

	return books, nil

}

func (s *BookService) CreateBook(toCreate models.CreateBook) (models.Book, error) {

	rr := &pb.CreateBookRequest{
		Title:        toCreate.Title,
		Author:       toCreate.Author,
		Genre:        toCreate.Genre,
		TotalPages:   int32(toCreate.TotalPages),
		UserId:       int32(toCreate.UserId),
		Cover:        toCreate.Cover,
		Description:  toCreate.Description,
		PublicatedAt: toCreate.PublicatedAt,
	}

	resp, err := s.client.CreateBook(context.Background(), rr)

	if err != nil {
		return models.Book{}, errors.New("failed to create the book")
	}

	return mapProtosToBookModel(resp.GetBook()), nil
}

func (s *BookService) UpdateBook(book models.Book) (models.Book, error) {

	res, err := s.client.UpdateBook(context.Background(), mapBookModelToProto(book))

	if err != nil {
		return models.Book{}, err
	}

	return mapProtosToBookModel(res.GetBook()), nil
}
