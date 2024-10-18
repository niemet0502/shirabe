package services

import (
	"context"
	"errors"

	"github.com/niemet0502/shirabe/bff/models"
	pb "github.com/niemet0502/shirabe/books/readingprogress"
)

type ReadingProgressService struct {
	cc pb.ReadingServiceClient
}

func NewReadingService(cc pb.ReadingServiceClient) *ReadingProgressService {
	return &ReadingProgressService{cc}
}

func mapProgressProtoToEntity(proto *pb.ReadingProgress) models.ReadingProgress {
	return models.ReadingProgress{
		ID:        uint(proto.GetId()),
		UserId:    int(proto.GetUserId()),
		BookId:    int(proto.GetBookId()),
		CreatedAt: proto.GetCreatedAt(),
	}
}

func (svc *ReadingProgressService) Create(toCreate models.CreateReadingProgress) (models.ReadingProgress, error) {

	rr, err := svc.cc.CreateReadingProgress(context.Background(), &pb.CreateReadingProgressRequest{
		BookId:      int32(toCreate.BookId),
		UserId:      int32(toCreate.UserId),
		PagesNumber: int32(toCreate.PagesNumber),
	})

	if err != nil {
		return models.ReadingProgress{}, errors.New("failed to create reading progress")
	}

	return mapProgressProtoToEntity(rr), nil
}

func (svc *ReadingProgressService) GetByBookId(bookId int) ([]models.ReadingProgress, error) {
	var result []models.ReadingProgress

	rr, err := svc.cc.GetReadingProgress(context.Background(), &pb.GetReadingProgressRequest{BookId: int32(bookId)})

	if err != nil {
		return result, errors.New(err.Error())
	}

	for _, progress := range rr.ReadingProgresses {
		result = append(result, mapProgressProtoToEntity(progress))
	}

	return result, nil
}
