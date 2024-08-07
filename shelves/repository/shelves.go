package repository

import (
	"errors"
	"fmt"

	"gorm.io/gorm"

	"github.com/niemet0502/shirabe/shelves/models"
)

type ShelvesRepository struct {
	db *gorm.DB
}

func NewShelvesRepository(db *gorm.DB) *ShelvesRepository {
	return &ShelvesRepository{db}
}

func (r *ShelvesRepository) GetShelvesByUser(userId int) []models.Shelf {
	var shelves []models.Shelf

	err := r.db.Find(&shelves, "user_id = ?", userId).Error

	if err != nil {
		fmt.Printf("Unable to find shelves")
	}

	return shelves
}

func (r *ShelvesRepository) CreateShelf(name string, userId int) models.Shelf {
	shelf := models.Shelf{Name: name, UserId: int32(userId)}

	r.db.Create(&shelf)

	return shelf
}

func (r *ShelvesRepository) RemoveShelf(shelfId int) error {
	var shelf models.Shelf

	if err := r.db.First(&shelf, shelfId).Error; err != nil {
		return errors.New("shelf not found")
	}

	r.db.Delete(&models.Shelf{}, shelfId)

	return nil
}

func (r *ShelvesRepository) GetShelf(shelfId int) (models.Shelf, error) {
	var shelf models.Shelf

	if err := r.db.First(&shelf, shelfId).Error; err != nil {
		return models.Shelf{}, errors.New("shelf not found")
	}

	return shelf, nil
}
