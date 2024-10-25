package services

import (
	"context"
	"errors"

	"github.com/niemet0502/shirabe/bff/models"
	pb "github.com/niemet0502/shirabe/users/user"
)

type UserService struct {
	cc pb.UserServiceClient
}

func NewUserService(cc pb.UserServiceClient) *UserService {
	return &UserService{cc}
}

func (svc *UserService) CreateUser(createUser models.CreateUser) (models.User, error) {

	rr := &pb.CreateUserRequest{
		Email:    createUser.Email,
		Password: createUser.Password,
	}

	res, err := svc.cc.CreateUser(context.Background(), rr)

	if err != nil {
		return models.User{}, err
	}

	return models.User{ID: uint(res.GetUser().Id), UserName: res.GetUser().GetUsername(), Email: res.GetUser().Email}, nil

}

func (svc *UserService) GetUser(id int) (models.User, error) {
	rr := &pb.GetUserRequest{
		UserId: int32(id),
	}

	res, err := svc.cc.GetUser(context.Background(), rr)

	if err != nil {
		return models.User{}, errors.New("User not found")
	}

	return models.User{ID: uint(res.GetUser().Id), UserName: res.GetUser().GetUsername(), Email: res.GetUser().Email}, nil
}
