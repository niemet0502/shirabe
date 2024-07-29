package services

import (
	"context"

	"github.com/niemet0502/shirabe/bff/models"
	pb "github.com/niemet0502/shirabe/shelves/shelve"
)

type ShelvesService struct {
	cc pb.ShelveServiceClient
}

func mapShelPrototoModel(pb *pb.Shelf) models.Shelf {
	return models.Shelf{
		ID:     int(pb.GetId()),
		Name:   pb.GetName(),
		UserId: int(pb.GetUserId()),
	}
}

func NewShelvesService(cc pb.ShelveServiceClient) *ShelvesService {
	return &ShelvesService{cc}
}

func (svc *ShelvesService) GetShelves(userId int) ([]models.Shelf, error) {

	rr := &pb.GetShelvesRequest{UserId: int32(userId)}
	res, err := svc.cc.GetShelves(context.Background(), rr)

	if err != nil {
		return []models.Shelf{}, err
	}

	var shelves []models.Shelf

	for _, s := range res.GetShelves() {
		shelves = append(shelves, mapShelPrototoModel(s))
	}

	return shelves, nil
}

func (svc *ShelvesService) CreateShelf(createShelf models.CreateShelf) (models.Shelf, error) {

	rr := &pb.CreateShelfRequest{
		Name:   createShelf.Name,
		UserId: int32(createShelf.UserId),
	}

	res, err := svc.cc.CreateShelf(context.Background(), rr)

	if err != nil {
		return models.Shelf{}, err
	}

	return mapShelPrototoModel(res.GetShelf()), nil

}

func (svc *ShelvesService) GetShelf(shelfId int) (models.Shelf, error) {

	res, err := svc.cc.GetShelf(context.Background(), &pb.RemoveShelfRequest{
		ShelfId: int32(shelfId),
	})

	if err != nil {
		return models.Shelf{}, nil
	}

	return mapShelPrototoModel(res.GetShelf()), nil
}

func (svc *ShelvesService) RemoveShelf(shelfId int) (string, error) {
	res, err := svc.cc.RemoveShelf(context.Background(), &pb.RemoveShelfRequest{
		ShelfId: int32(shelfId),
	})

	if err != nil {
		return "", nil
	}

	return res.GetMessage(), nil
}

func (svc *ShelvesService) UpdateShelf(shelf models.Shelf) (models.Shelf, error) {
	res, err := svc.cc.UpdateShelf(context.Background(), &pb.UpdateShelfRequest{
		Id:     int32(shelf.ID),
		Name:   shelf.Name,
		UserId: int32(shelf.UserId),
	})

	if err != nil {
		return models.Shelf{}, nil
	}

	return mapShelPrototoModel(res.GetShelf()), nil
}
