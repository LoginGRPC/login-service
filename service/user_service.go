package service

import (
	"golang.org/x/net/context"

	sdk "github.com/LoginGRPC/login-contractor"
)

type Server struct{

}

func (s *Server) Ping(ctx context.Context, in *sdk.PingRequest) (*sdk.PingResponse, error) {
	return &sdk.PingResponse{Message: "Hello " + in.UserName}, nil
}

func (s *Server) GetUserProfile(ctx context.Context, in *sdk.PingRequest) (*sdk.PingResponse, error) {
	return &sdk.LoginResponse{
		Result: true,
		User: *sdk.UserProfile{
			"123", "Parmatma", 123456789,
		},
	}, nil
}

