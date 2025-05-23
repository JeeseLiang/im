// Code generated by goctl. DO NOT EDIT.
// goctl 1.8.2
// Source: user.proto

package userclient

import (
	"context"

	"im_message/proto/user"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	GetOnlineStatusRequest     = user.GetOnlineStatusRequest
	GetOnlineStatusResponse    = user.GetOnlineStatusResponse
	LoginRequest               = user.LoginRequest
	LoginResponse              = user.LoginResponse
	ModifyUserInfoRequest      = user.ModifyUserInfoRequest
	ModifyUserInfoResponse     = user.ModifyUserInfoResponse
	PersonalInfoRequest        = user.PersonalInfoRequest
	PersonalInfoResponse       = user.PersonalInfoResponse
	RegisterRequest            = user.RegisterRequest
	RegisterResponse           = user.RegisterResponse
	ResetPasswordRequest       = user.ResetPasswordRequest
	ResetPasswordResponse      = user.ResetPasswordResponse
	UpdateOnlineStatusRequest  = user.UpdateOnlineStatusRequest
	UpdateOnlineStatusResponse = user.UpdateOnlineStatusResponse

	UserClient interface {
		Login(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*LoginResponse, error)
		Register(ctx context.Context, in *RegisterRequest, opts ...grpc.CallOption) (*RegisterResponse, error)
		PersonalInfo(ctx context.Context, in *PersonalInfoRequest, opts ...grpc.CallOption) (*PersonalInfoResponse, error)
		ResetPassword(ctx context.Context, in *ResetPasswordRequest, opts ...grpc.CallOption) (*ResetPasswordResponse, error)
		UpdateOnlineStatus(ctx context.Context, in *UpdateOnlineStatusRequest, opts ...grpc.CallOption) (*UpdateOnlineStatusResponse, error)
		GetOnlineStatus(ctx context.Context, in *GetOnlineStatusRequest, opts ...grpc.CallOption) (*GetOnlineStatusResponse, error)
		ModifyUserInfo(ctx context.Context, in *ModifyUserInfoRequest, opts ...grpc.CallOption) (*ModifyUserInfoResponse, error)
	}

	defaultUserClient struct {
		cli zrpc.Client
	}
)

func NewUserClient(cli zrpc.Client) UserClient {
	return &defaultUserClient{
		cli: cli,
	}
}

func (m *defaultUserClient) Login(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*LoginResponse, error) {
	client := user.NewUserClientClient(m.cli.Conn())
	return client.Login(ctx, in, opts...)
}

func (m *defaultUserClient) Register(ctx context.Context, in *RegisterRequest, opts ...grpc.CallOption) (*RegisterResponse, error) {
	client := user.NewUserClientClient(m.cli.Conn())
	return client.Register(ctx, in, opts...)
}

func (m *defaultUserClient) PersonalInfo(ctx context.Context, in *PersonalInfoRequest, opts ...grpc.CallOption) (*PersonalInfoResponse, error) {
	client := user.NewUserClientClient(m.cli.Conn())
	return client.PersonalInfo(ctx, in, opts...)
}

func (m *defaultUserClient) ResetPassword(ctx context.Context, in *ResetPasswordRequest, opts ...grpc.CallOption) (*ResetPasswordResponse, error) {
	client := user.NewUserClientClient(m.cli.Conn())
	return client.ResetPassword(ctx, in, opts...)
}

func (m *defaultUserClient) UpdateOnlineStatus(ctx context.Context, in *UpdateOnlineStatusRequest, opts ...grpc.CallOption) (*UpdateOnlineStatusResponse, error) {
	client := user.NewUserClientClient(m.cli.Conn())
	return client.UpdateOnlineStatus(ctx, in, opts...)
}

func (m *defaultUserClient) GetOnlineStatus(ctx context.Context, in *GetOnlineStatusRequest, opts ...grpc.CallOption) (*GetOnlineStatusResponse, error) {
	client := user.NewUserClientClient(m.cli.Conn())
	return client.GetOnlineStatus(ctx, in, opts...)
}

func (m *defaultUserClient) ModifyUserInfo(ctx context.Context, in *ModifyUserInfoRequest, opts ...grpc.CallOption) (*ModifyUserInfoResponse, error) {
	client := user.NewUserClientClient(m.cli.Conn())
	return client.ModifyUserInfo(ctx, in, opts...)
}
