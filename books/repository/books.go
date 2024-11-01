package repository

import (
	"errors"
	"fmt"
	"strings"

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

	var maxID int64
	repo.db.Table("books").Select("MAX(id)").Row().Scan(&maxID)

	book := models.Book{
		Title:           newBook.Title,
		Author:          newBook.Author,
		Genre:           newBook.Genre,
		Slug:            fmt.Sprintf("%d.%s", maxID+1, strings.ReplaceAll(newBook.Title, " ", "_")),
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

func (repo *BookRepository) GetBookBySlug(slug string) (models.Book, error) {
	var book models.Book

	result := repo.db.Find(&book, "slug = ?", slug)

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

func (repo *BookRepository) SearchBooks(userId, status int, search string) []models.Book {
	var books []models.Book
	query := fmt.Sprintf("%%%s%%", search)

	db := repo.db.Where("user_id = ? AND (title LIKE ? OR author LIKE ?)", userId, query, query)

	if status != 0 {
		db = db.Where("status = ?", status)
	}

	db.Find(&books)

	return books
}

func (repo *BookRepository) GetBooksByShelf(ids []int32) []models.Book {
	var books []models.Book
	repo.db.Find(&books, ids)

	return books
}
