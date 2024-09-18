package services

import (
	"context"
	"errors"
	"fmt"
	"log"
	"mime/multipart"

	"github.com/niemet0502/shirabe/bff/models"
	pb "github.com/niemet0502/shirabe/books/book"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
)

type BookService struct {
	client pb.BookClient
	s3svc  *s3.S3
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

func NewBookService(client pb.BookClient, s3svc *s3.S3) *BookService {
	return &BookService{client, s3svc}
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

func (s *BookService) CreateBook(toCreate models.CreateBook, file multipart.File, fileName string) (models.Book, error) {
	name := fmt.Sprintf("%d-%s-%s", toCreate.UserId, toCreate.Title, fileName)

	rr := &pb.CreateBookRequest{
		Title:        toCreate.Title,
		Author:       toCreate.Author,
		Genre:        toCreate.Genre,
		TotalPages:   int32(toCreate.TotalPages),
		UserId:       int32(toCreate.UserId),
		Cover:        name,
		Description:  toCreate.Description,
		PublicatedAt: toCreate.PublicatedAt,
	}

	// upload file
	if file != nil {
		input := &s3.PutObjectInput{
			Bucket: aws.String("shirabebookscover"),
			Key:    aws.String(name),
			Body:   file,
		}

		_, err := s.s3svc.PutObject(input)

		if err != nil {
			log.Printf("Failed to upload the file")
		}

		key := fmt.Sprintf("https://shirabebookscover.s3.eu-north-1.amazonaws.com/%s", name)

		rr.Cover = key

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
