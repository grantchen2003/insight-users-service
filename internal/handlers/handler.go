package handler

import (
	"context"
	"log"

	db "github.com/grantchen2003/insight/users/internal/database"
	pb "github.com/grantchen2003/insight/users/internal/protobufs"
	"google.golang.org/protobuf/types/known/emptypb"
)

type UsersServiceHandler struct {
	pb.UsersServiceServer
}

func (u *UsersServiceHandler) CreateUser(ctx context.Context, req *pb.CreateUserRequest) (*pb.CreateUsersResponse, error) {
	log.Println("received CreateUser request")

	database := db.GetSingletonInstance()

	user, err := database.GetUserBySessionId(req.SessionId)
	if err != nil {
		return nil, err
	}

	if user != nil {
		return &pb.CreateUsersResponse{UserId: user.Id}, nil
	}

	userId, err := database.SaveUser(req.SessionId)
	if err != nil {
		return nil, err
	}

	return &pb.CreateUsersResponse{UserId: userId}, nil
}

func (u *UsersServiceHandler) InitializeUser(ctx context.Context, req *pb.InitializeUserRequest) (*emptypb.Empty, error) {
	log.Println("received InitializeUser request")

	database := db.GetSingletonInstance()

	if err := database.SetUserIsInitialized(req.UserId, true); err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}
