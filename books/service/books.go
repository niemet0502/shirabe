package service

import (
	"github.com/niemet0502/shirabe/books/models"
	"github.com/niemet0502/shirabe/books/repository"
)

// the book package new the repository

type BookService struct {
	repo *repository.BookRepository
}

func NewBookService(repo *repository.BookRepository) *BookService {
	return &BookService{repo}
}

func (svc *BookService) GetBook(id int) (models.Book, error) {
	result, err := svc.repo.GetBook(id)

	if err != nil {
		return models.Book{}, err
	}

	return result, nil
}

func (svc *BookService) GetBooks(userId int) []models.Book {
	return svc.repo.GetBooksByUser(userId)
}

func (svc *BookService) CreateBook(newBook models.CreateBook) models.Book {
	return svc.repo.CreateBook(newBook)
}
