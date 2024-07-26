package service

import (
	"github.com/niemet0502/shirabe/shelves/models"
	"github.com/niemet0502/shirabe/shelves/repository"
)

type BookShelfService struct {
	repo *repository.BookShelfRepository
}

func NewBookShelfService(repo *repository.BookShelfRepository) *BookShelfService {
	return &BookShelfService{repo}
}

func (svc *BookShelfService) AddBookToShelf(bookId, shelfId int) (models.BookShelf, error) {
	result, err := svc.repo.AddBookToShelf(int32(bookId), int32(shelfId))

	if err != nil {
		return models.BookShelf{}, err
	}

	return result, nil
}

func (svc *BookShelfService) RemoveBookFromShelf(bookId, shelfId int) error {
	return svc.repo.RemoveBookFromShelf(int32(bookId), int32(shelfId))
}

func (svc *BookShelfService) GetShelfBooks(bookId, shelfId int) []models.BookShelf {
	return svc.repo.GetShelfBooks(bookId, shelfId)

}
