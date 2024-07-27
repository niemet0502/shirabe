package service

import (
	"github.com/niemet0502/shirabe/users/models"
	"github.com/niemet0502/shirabe/users/repository"
	"github.com/niemet0502/shirabe/users/utils"
)

type UserService struct {
	repo *repository.UserRepository
}

func NewUserService(repo *repository.UserRepository) *UserService {
	return &UserService{repo}
}

func (svc *UserService) CreateUser(email, password string) (models.User, error) {
	// encrypt the password

	hash, _ := utils.HashPassword(password)

	user, err := svc.repo.CreateUser(email, hash)

	if err != nil {
		return models.User{}, err
	}

	return user, nil

}
