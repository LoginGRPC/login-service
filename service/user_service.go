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

func (s *Server) GetUserProfile(ctx context.Context, in *sdk.LoginRequest) (*sdk.LoginResponse, error) {
	return &sdk.LoginResponse{
		Result: true,
		User: &sdk.UserProfile{
			UserId:"1223",
			UserName:"Parmatma",
			ContactNumber:"1213245",

		},
	}, nil
}

