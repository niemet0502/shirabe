package server

import (
	"github.com/niemet0502/shirabe/users/service"
	pb "github.com/niemet0502/shirabe/users/user"
)

type Server struct {
	pb.UnimplementedUserServiceServer
	svc *service.UserService
}

func NewServer(svc *service.UserService) *Server {
	return &Server{svc: svc}
}
