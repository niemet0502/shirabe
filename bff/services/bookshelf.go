package services

import (
	"context"

	"github.com/niemet0502/shirabe/bff/models"
	bpb "github.com/niemet0502/shirabe/books/book"
	spb "github.com/niemet0502/shirabe/shelves/shelve"
)

type BookShelfService struct {
	bsc spb.BookShelfServiceClient
	sc  spb.ShelveServiceClient
	bc  bpb.BookClient
}

func NewBoolShelfService(bsc spb.BookShelfServiceClient,
	sc spb.ShelveServiceClient,
	bc bpb.BookClient) *BookShelfService {
	return &BookShelfService{bsc, sc, bc}
}

func (svc *BookShelfService) GetBooksByShelf(shelfId int) ([]models.Book, error) {
	var books []models.Book

	rr := &spb.BookShelf{ShelfId: int32(shelfId), BookId: 0}

	res, err := svc.bsc.SearchBookShelf(context.Background(), rr)

	if err != nil {
		return books, err
	}

	var bookIds []int32

	for _, s := range res.GetBookShelves() {
		bookIds = append(bookIds, s.BookId)
	}

	// fetch books

	result, error := svc.bc.GetBooksByShelf(context.Background(), &bpb.GetBooksByShelfRequest{Ids: bookIds})

	if error != nil {
		return books, error
	}

	// return books

	for _, b := range result.GetBooks() {
		books = append(books, mapProtosToBookModel(b))
	}

	return books, nil

}
func (svc *BookShelfService) GetShelvesByBook() {}
func (svc *BookShelfService) AddBookToShelf(shelfId, bookId int) (models.BookShelf, error) {

	res, err := svc.bsc.AddBook(context.Background(), &spb.CreateBookShelf{
		BookId:  int32(bookId),
		ShelfId: int32(shelfId),
	})

	if err != nil {
		return models.BookShelf{}, err
	}

	return models.BookShelf{
		ID:      int(res.GetBookshelf().GetId()),
		BookId:  int(res.GetBookshelf().GetBookId()),
		ShelfId: int(res.GetBookshelf().ShelfId),
	}, nil
}
func (svc *BookShelfService) RemoveBookFromShel(bookId, shelfId int) (string, error) {

	res, err := svc.bsc.RemoveBook(context.Background(), &spb.CreateBookShelf{
		BookId:  int32(bookId),
		ShelfId: int32(shelfId),
	})

	if err != nil {
		return "", err
	}
	return res.GetMessage(), nil
}
