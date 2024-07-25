package repository

import (
	"errors"

	"gorm.io/gorm"

	"github.com/niemet0502/shirabe/books/models"
)

// the book repository needs the db connection
type BookRepository struct {
	db *gorm.DB
}

func NewBookRepository(db *gorm.DB) *BookRepository {
	return &BookRepository{db}
}

func (repo *BookRepository) CreateBook(newBook models.CreateBook) models.Book {
	book := models.Book{
		Title:           newBook.Title,
		Author:          newBook.Author,
		Genre:           newBook.Genre,
		Status:          1,
		TotalPages:      newBook.TotalPages,
		ReadingProgress: 0,
		UserId:          newBook.UserId,
		Cover:           newBook.Cover,
		Description:     newBook.Description,
		PublicatedAt:    newBook.PublicatedAt,
	}

	repo.db.Create(&book)

	return models.Book(book)
}

func (repo *BookRepository) GetBook(id int) (models.Book, error) {
	var book models.Book

	result := repo.db.First(&book, id)

	if result.Error != nil {
		return models.Book{}, errors.New("failed to fetch the book")
	}

	return book, nil
}

func (repo *BookRepository) GetBooksByUser(id int) []models.Book {
	var result []models.Book

	repo.db.Find(&result, "user_id = ?", id)

	return result
}

func (repo *BookRepository) UpdateBook(book models.Book) models.Book {
	repo.db.Save(&book)

	return book

}
