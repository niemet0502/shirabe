package repository

import (
	"github.com/niemet0502/shirabe/books/models"
	"gorm.io/gorm"
)

type ReadingProgressRepository struct {
	db *gorm.DB
}

func NewReadingProgressRepository(db *gorm.DB) *ReadingProgressRepository {
	return &ReadingProgressRepository{db}
}

func (repo *ReadingProgressRepository) Create(toCreate models.CreateReadingProgress) models.ReadingProgress {
	readingProgress := models.ReadingProgress{
		UserId:      toCreate.UserId,
		BookId:      toCreate.BookId,
		PagesNumber: toCreate.PagesNumber,
	}

	repo.db.Create(&readingProgress)

	return models.ReadingProgress(readingProgress)
}

func (repo *ReadingProgressRepository) GetByBooks(bookId int) []models.ReadingProgress {
	var readingProgresses []models.ReadingProgress

	_ = repo.db.Where("book_id = ? ", bookId).Order("id desc").Find(&readingProgresses)

	return readingProgresses
}
