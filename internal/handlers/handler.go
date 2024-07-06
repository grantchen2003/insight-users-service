package handler

import (
	"context"
	"log"

	db "github.com/grantchen2003/insight/users/internal/database"
	pb "github.com/grantchen2003/insight/users/internal/protobufs"
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
		log.Println(user)
		return &pb.CreateUsersResponse{UserId: user.Id}, nil
	}

	userId, err := database.SaveUser(req.SessionId)
	if err != nil {
		return nil, err
	}

	return &pb.CreateUsersResponse{UserId: userId}, nil
}
