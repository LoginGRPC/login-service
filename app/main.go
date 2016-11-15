package main

import (
	"log"
	"net"

	"golang.org/x/net/context"
	"google.golang.org/grpc"

	sdk "github.com/LoginGRPC/login-contractor"
	"github.com/facebookgo/inject"
)

const (
	port = ":5005"
)

type Server struct{

}

func (s *Server) Ping(ctx context.Context, in *sdk.PingRequest) (*sdk.PingResponse, error) {
	return &sdk.PingResponse{Message: "Hello " + in.UserName}, nil
}

func (s *Server) GetUserProfile(ctx context.Context, in *sdk.PingRequest) (*sdk.PingResponse, error) {
	return &sdk.LoginResponse{Result: true, User: *sdk.UserProfile{
		"123", "Parmatma", 123456789,
	}}, nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	sdk.RegisterUserServiceServer(s, &Server{})

	var g inject.Graph
	var userService Server
	err = g.Provide(
		&inject.Object{Value: &userService},
	)

	if err = g.Populate(); err != nil {
		panic(err)
	}

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
