package main

import (
	"log"
	"net"

	"google.golang.org/grpc"

	sdk "github.com/LoginGRPC/login-contractor"
	"github.com/facebookgo/inject"
	"yadavparmatma/login-service/service"
)

const (
	port = ":5005"
)

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	sdk.RegisterUserServiceServer(s, &service.Server{})

	var g inject.Graph
	var userService service.Server
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
