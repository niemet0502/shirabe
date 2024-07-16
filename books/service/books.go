package service

import (
	"github.com/niemet0502/shirabe/books/repository"
)

// the book package new the repository

type BookService struct {
	repo *repository.BookRepository
}

func NewBookService(repo *repository.BookRepository) *BookService {
	return &BookService{repo}
}

func (svc *BookService) GetBook() {
	svc.repo.GetBook()
}
