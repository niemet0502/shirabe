package server

import (
	"context"

	"github.com/niemet0502/shirabe/books/models"
	protos "github.com/niemet0502/shirabe/books/readingprogress"
	"github.com/niemet0502/shirabe/books/service"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type readingProgressServer struct {
	protos.UnimplementedReadingServiceServer
	svc *service.ReadingProgressService
}

func NewReadingProgressServer(svc *service.ReadingProgressService) *readingProgressServer {
	return &readingProgressServer{svc: svc}
}

func (r *readingProgressServer) CreateReadingProgress(ctx context.Context, rr *protos.CreateReadingProgressRequest) (*protos.ReadingProgress, error) {
	readingProgress, err := r.svc.Create(models.CreateReadingProgress{
		UserId:      int(rr.GetUserId()),
		BookId:      int(rr.GetBookId()),
		PagesNumber: int(rr.GetPagesNumber()),
	})

	if err != nil {
		return &protos.ReadingProgress{}, status.Errorf(codes.FailedPrecondition, "Failed to create the reading progress")
	}

	return &protos.ReadingProgress{
		Id:        int32(readingProgress.ID),
		UserId:    int32(readingProgress.UserId),
		BookId:    int32(readingProgress.BookId),
		CreatedAt: readingProgress.CreatedAt.String(),
	}, nil
}

func (r *readingProgressServer) GetReadingProgress(ctx context.Context, rr *protos.GetReadingProgressRequest) (*protos.ReadingProgresses, error) {
	readingProgresses := r.svc.GetByBookId(int(rr.GetBookId()))

	var result []*protos.ReadingProgress

	for _, progress := range readingProgresses {
		result = append(result, &protos.ReadingProgress{
			Id:          int32(progress.ID),
			UserId:      int32(progress.UserId),
			BookId:      int32(progress.BookId),
			PagesNumber: int32(progress.PagesNumber),
			CreatedAt:   progress.CreatedAt.String(),
		})
	}

	return &protos.ReadingProgresses{ReadingProgresses: result}, nil
}
