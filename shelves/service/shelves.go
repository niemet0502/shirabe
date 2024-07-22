package service

import (
	"github.com/niemet0502/shirabe/shelves/repository"
)

type ShelvesService struct {
	repo *repository.ShelvesRepository
}

func NewShelvesService(repo *repository.ShelvesRepository) *ShelvesService {
	return &ShelvesService{repo}
}
