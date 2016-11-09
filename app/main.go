package main

import (
	"log"
	"net"

	"golang.org/x/net/context"
	"google.golang.org/grpc"

	sdk "github.com/LoginGRPC/login-contractor"
)

const (
	port = ":5005"
)

type server struct{}


func (s *server) Ping(ctx context.Context, in *sdk.PingRequest) (*sdk.PingResponse, error) {
	return &sdk.PingResponse{Message: "Hello " + in.UserName}, nil
}

func (s *server) GetUserProfile(ctx context.Context, in *sdk.PingRequest) (*sdk.PingResponse, error) {
	return &sdk.LoginResponse{Result: true, User:*sdk.UserProfile{
		"123","Parmatma",123456789,
	}}, nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	s.RegisterService(*sdk.UserService{}, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
