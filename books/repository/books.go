package repository

import (
	"database/sql"
	"errors"
	"fmt"

	"github.com/niemet0502/shirabe/books/models"
)

// the book repository needs the db connection
type BookRepository struct {
	db    *sql.DB
	books []models.Book
}

func NewBookRepository(db *sql.DB) *BookRepository {
	return &BookRepository{
		db: db,
		books: []models.Book{
			{
				ID:              1,
				Title:           "Harry Potter",
				Author:          "JK Rowling",
				Genre:           "Science fiction",
				Status:          1, // Assuming Status is an integer value.
				TotalPages:      256,
				ReadingProgress: 0,
				UserId:          1,
			},
			{
				ID:              1,
				Title:           "Hunger games",
				Author:          "Collins",
				Genre:           "Science fiction",
				Status:          1, // Assuming Status is an integer value.
				TotalPages:      256,
				ReadingProgress: 0,
				UserId:          2,
			},
		},
	}
}

func (repo *BookRepository) CreateBook(newBook models.CreateBook) models.Book {
	bookToCreate := models.Book{
		ID:              len(repo.books) + 1,
		Title:           newBook.Title,
		Author:          newBook.Author,
		Genre:           newBook.Genre,
		UserId:          newBook.UserId,
		TotalPages:      newBook.TotalPages,
		ReadingProgress: 0,
		Status:          1,
	}
	repo.books = append(repo.books, bookToCreate)

	return bookToCreate
}

func (repo *BookRepository) GetBook(id int) (models.Book, error) {
	fmt.Print("Failed to fetch the book")

	for _, b := range repo.books {
		if b.ID == id {
			return b, nil
		}
	}

	return models.Book{}, errors.New("Book doesn't exist")
}

func (repo *BookRepository) GetBooksByUser(id int) []models.Book {
	var result []models.Book

	for _, b := range repo.books {
		if b.UserId == id {
			result = append(result, b)
		}
	}

	return result
}
