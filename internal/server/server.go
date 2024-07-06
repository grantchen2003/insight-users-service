package server

import (
	"log"
	"net"

	handler "github.com/grantchen2003/insight/users/internal/handlers"
	pb "github.com/grantchen2003/insight/users/internal/protobufs"
	"google.golang.org/grpc"
)

type Server struct {
	grpcServer *grpc.Server
}

func NewServer() *Server {
	grpcServer := grpc.NewServer()

	pb.RegisterUsersServiceServer(
		grpcServer, &handler.UsersServiceHandler{},
	)

	return &Server{grpcServer: grpcServer}
}

func (server Server) Start(address string) error {
	listener, err := net.Listen("tcp", address)

	if err != nil {
		log.Fatalf("failed to listen: %v", err)
		return err
	}

	log.Printf("server listening on %s", address)

	if err := server.grpcServer.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %v", err)
		return err
	}

	return nil
}
