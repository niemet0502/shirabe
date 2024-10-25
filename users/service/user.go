package service

import (
	"errors"

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

func (svc *UserService) GetUser(id int) (models.User, error) {

	result, err := svc.repo.GetUser(id)

	if err != nil {
		return models.User{}, errors.New("failed to get user")
	}

	return result, nil
}
