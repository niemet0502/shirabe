package server

import (
	"context"

	"github.com/niemet0502/shirabe/users/models"
	"github.com/niemet0502/shirabe/users/service"
	pb "github.com/niemet0502/shirabe/users/user"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Server struct {
	pb.UnimplementedUserServiceServer
	svc *service.UserService
}

func mapModelToProto(user models.User) *pb.User {
	return &pb.User{
		Id:       int32(user.ID),
		Username: user.UserName,
		Email:    user.Email,
		Password: user.Password,
	}
}

func NewServer(svc *service.UserService) *Server {
	return &Server{svc: svc}
}

func (s *Server) CreateUser(ctx context.Context, rr *pb.CreateUserRequest) (*pb.CreateUserResponse, error) {

	user, err := s.svc.CreateUser(rr.GetEmail(), rr.GetPassword())

	if err != nil {
		return &pb.CreateUserResponse{}, status.Errorf(codes.AlreadyExists, "Email already exists")
	}

	return &pb.CreateUserResponse{User: mapModelToProto(user)}, nil
}
