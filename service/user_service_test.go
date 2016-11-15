package service

import (
	"testing"
	"golang.org/x/net/context"
	sdk "github.com/LoginGRPC/login-contractor"
	"github.com/stretchr/testify/assert"
)

func TestServer_Ping(t *testing.T) {
	server := &Server{}
	res, err := server.Ping(context.TODO(), &sdk.PingRequest{
		UserName:"Parmatma",
	})

	assert.NotNil(t,res)
	assert.Nil(t,err)

	assert.Equal(t,res.Message,"Hello Parmatma");
}

func TestServer_GetUserProfile(t *testing.T) {
	server := &Server{}
	res, err := server.GetUserProfile(context.TODO(), &sdk.LoginRequest{
		UserId:"1223",
		Password:"Parmatma",
	})

	assert.NotNil(t,res)
	assert.Nil(t,err)

	assert.Equal(t,res.User,&sdk.UserProfile{
		UserId:"1223",
		UserName:"Parmatma",
		ContactNumber:"1213245",

	});

}