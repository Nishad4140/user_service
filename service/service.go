package service

import (
	"context"
	"fmt"

	"github.com/Nishad4140/proto_files/pb"
	"github.com/Nishad4140/user_service/adapter"
	"github.com/Nishad4140/user_service/entities"
	"github.com/Nishad4140/user_service/helper"
	"github.com/opentracing/opentracing-go"
)

var (
	Tracer opentracing.Tracer
)

func RetrieveTracer(tr opentracing.Tracer) {
	Tracer = tr
}

type UserService struct {
	Adapter adapter.UserInterface
	pb.UnimplementedUserServiceServer
}

func NewUserService(adapter adapter.UserInterface) *UserService {
	return &UserService{
		Adapter: adapter,
	}
}

func (u *UserService) UserSignUp(ctx context.Context, req *pb.UserSignUpRequest) (*pb.UserResponse, error) {
	if req.Name == "" {
		return nil, fmt.Errorf("name cant be empty")
	} else if req.Email == "" {
		return nil, fmt.Errorf("email cant be empty")
	} else if req.Password == "" {
		return nil, fmt.Errorf("password cant be empty")
	}

	hashed, err := helper.HashPassword(req.Password)
	if err != nil {
		return nil, err
	}
	req.Password = hashed

	ogreq := entities.User{
		Name:     req.Name,
		Email:    req.Email,
		Password: req.Password,
	}

	res, err := u.Adapter.UserSignUp(ogreq)
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}

	return &pb.UserResponse{
		Id:    uint32(res.Id),
		Name:  res.Name,
		Email: res.Email,
	}, nil
}
