package service

import (
	"errors"
	"time"

	"github.com/niemet0502/shirabe/books/models"
	"github.com/niemet0502/shirabe/books/repository"
)

type ReadingProgressService struct {
	repo *repository.ReadingProgressRepository

	bookService *BookService
}

func NewReadingProgressService(repo *repository.ReadingProgressRepository, bookService *BookService) *ReadingProgressService {
	return &ReadingProgressService{repo, bookService}
}

// Create

func (svc *ReadingProgressService) Create(toCreate models.CreateReadingProgress) (models.ReadingProgress, error) {
	// get the book details first by id
	book, er := svc.bookService.GetBook(toCreate.BookId)

	if er != nil {
		return models.ReadingProgress{}, errors.New("failed to fetch the book")
	}

	if toCreate.PagesNumber == book.TotalPages {
		book.FinishReadingAt = time.Now()
	}

	book.ReadingProgress = toCreate.PagesNumber

	// update the books details

	svc.bookService.UpdateBook(book)

	// create the reading progress entry (only if the update was successfull)

	return svc.repo.Create(toCreate), nil
}

func (svc *ReadingProgressService) GetByBookId(bookId int) []models.ReadingProgress {
	return svc.repo.GetByBooks(bookId)
}
