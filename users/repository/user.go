package repository

import (
	"errors"

	"github.com/niemet0502/shirabe/users/models"
	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db}
}

func (r *UserRepository) CreateUser(email, password string) (models.User, error) {
	user := models.User{Email: email, Password: password}

	if err := r.db.Create(&user).Error; err != nil {
		return models.User{}, errors.New("email already in use")
	}

	return user, nil
}
