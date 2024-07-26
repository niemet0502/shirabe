package repository

import (
	"errors"
	"fmt"

	"github.com/niemet0502/shirabe/shelves/models"
	"gorm.io/gorm"
)

type BookShelfRepository struct {
	db *gorm.DB
}

func NewBookShelfRepository(db *gorm.DB) *BookShelfRepository {
	return &BookShelfRepository{db}
}

func (r *BookShelfRepository) AddBookToShelf(bookId, shelfId int32) (models.BookShelf, error) {
	bs := models.BookShelf{BookId: bookId, ShelfId: shelfId}

	if err := r.db.Create(&bs).Error; err != nil {
		return models.BookShelf{}, errors.New("failed to add book to the shelf")
	}

	return bs, nil
}

func (r *BookShelfRepository) RemoveBookFromShelf(bookId, shelfId int32) error {
	bs := models.BookShelf{BookId: bookId, ShelfId: shelfId}

	if err := r.db.Delete(&bs).Error; err != nil {
		return errors.New("failed to remove book from shelf")
	}

	return nil

}

func (r *BookShelfRepository) GetShelfBooks(bookId, shelfId int) []models.BookShelf {
	var bookShelves []models.BookShelf

	query := r.db.Model(&models.BookShelf{})

	if bookId != 0 {
		fmt.Printf("tt")
		query = query.Where("book_id = ?", bookId)
	}

	if shelfId != 0 {
		fmt.Printf("tt")
		query = query.Where("shelf_id = ?", shelfId)
	}

	query.Find(&bookShelves)

	return bookShelves
}
