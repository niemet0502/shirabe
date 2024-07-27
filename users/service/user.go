package service

import (
	"github.com/niemet0502/shirabe/users/models"
	"github.com/niemet0502/shirabe/users/repository"
)

type UserService struct {
	repo *repository.UserRepository
}

func NewUserService(repo *repository.UserRepository) *UserService {
	return &UserService{repo}
}

func (svc *UserService) CreateUser(email, password string) (models.User, error) {
	// encrypt the password
	user, err := svc.repo.CreateUser(email, password)

	if err != nil {
		return models.User{}, err
	}

	return user, nil

}
