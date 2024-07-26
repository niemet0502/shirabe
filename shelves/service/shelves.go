package service

import (
	"github.com/niemet0502/shirabe/shelves/models"
	"github.com/niemet0502/shirabe/shelves/repository"
)

type ShelvesService struct {
	repo *repository.ShelvesRepository
}

func NewShelvesService(repo *repository.ShelvesRepository) *ShelvesService {
	return &ShelvesService{repo}
}

func (svc *ShelvesService) GetShelvesByUser(userId int) []models.Shelf {
	return svc.repo.GetShelvesByUser(userId)
}

func (svc *ShelvesService) CreateShelf(name string, userId int) models.Shelf {
	return svc.repo.CreateShelf(name, userId)
}

func (svc *ShelvesService) RemoveShelf(shelfId int) error {
	return svc.repo.RemoveShelf(shelfId)
}
