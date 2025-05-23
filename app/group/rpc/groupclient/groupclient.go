// Code generated by goctl. DO NOT EDIT.
// goctl 1.8.2
// Source: group.proto

package groupclient

import (
	"context"

	"im_message/proto/group"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	AddFriendRequest             = group.AddFriendRequest
	AddFriendResponse            = group.AddFriendResponse
	AddGroupChatRequest          = group.AddGroupChatRequest
	AddGroupChatResponse         = group.AddGroupChatResponse
	CreateGroupChatRequest       = group.CreateGroupChatRequest
	CreateGroupChatResponse      = group.CreateGroupChatResponse
	FriendListRequest            = group.FriendListRequest
	FriendListResponse           = group.FriendListResponse
	GroupUserListRequest         = group.GroupUserListRequest
	GroupUserListResponse        = group.GroupUserListResponse
	HandleFriendRequest          = group.HandleFriendRequest
	HandleFriendResponse         = group.HandleFriendResponse
	MessageGroupInfo             = group.MessageGroupInfo
	MessageGroupInfoListRequest  = group.MessageGroupInfoListRequest
	MessageGroupInfoListResponse = group.MessageGroupInfoListResponse
	UserGroupListRequest         = group.UserGroupListRequest
	UserGroupListResponse        = group.UserGroupListResponse

	GroupClient interface {
		AddFriend(ctx context.Context, in *AddFriendRequest, opts ...grpc.CallOption) (*AddFriendResponse, error)
		HandleFriend(ctx context.Context, in *HandleFriendRequest, opts ...grpc.CallOption) (*HandleFriendResponse, error)
		GroupUserList(ctx context.Context, in *GroupUserListRequest, opts ...grpc.CallOption) (*GroupUserListResponse, error)
		UserGroupList(ctx context.Context, in *UserGroupListRequest, opts ...grpc.CallOption) (*UserGroupListResponse, error)
		MessageGroupInfoList(ctx context.Context, in *MessageGroupInfoListRequest, opts ...grpc.CallOption) (*MessageGroupInfoListResponse, error)
		AddGroupChat(ctx context.Context, in *AddGroupChatRequest, opts ...grpc.CallOption) (*AddGroupChatResponse, error)
		CreateGroupChat(ctx context.Context, in *CreateGroupChatRequest, opts ...grpc.CallOption) (*CreateGroupChatResponse, error)
		GetFriendListByUserId(ctx context.Context, in *FriendListRequest, opts ...grpc.CallOption) (*FriendListResponse, error)
	}

	defaultGroupClient struct {
		cli zrpc.Client
	}
)

func NewGroupClient(cli zrpc.Client) GroupClient {
	return &defaultGroupClient{
		cli: cli,
	}
}

func (m *defaultGroupClient) AddFriend(ctx context.Context, in *AddFriendRequest, opts ...grpc.CallOption) (*AddFriendResponse, error) {
	client := group.NewGroupClientClient(m.cli.Conn())
	return client.AddFriend(ctx, in, opts...)
}

func (m *defaultGroupClient) HandleFriend(ctx context.Context, in *HandleFriendRequest, opts ...grpc.CallOption) (*HandleFriendResponse, error) {
	client := group.NewGroupClientClient(m.cli.Conn())
	return client.HandleFriend(ctx, in, opts...)
}

func (m *defaultGroupClient) GroupUserList(ctx context.Context, in *GroupUserListRequest, opts ...grpc.CallOption) (*GroupUserListResponse, error) {
	client := group.NewGroupClientClient(m.cli.Conn())
	return client.GroupUserList(ctx, in, opts...)
}

func (m *defaultGroupClient) UserGroupList(ctx context.Context, in *UserGroupListRequest, opts ...grpc.CallOption) (*UserGroupListResponse, error) {
	client := group.NewGroupClientClient(m.cli.Conn())
	return client.UserGroupList(ctx, in, opts...)
}

func (m *defaultGroupClient) MessageGroupInfoList(ctx context.Context, in *MessageGroupInfoListRequest, opts ...grpc.CallOption) (*MessageGroupInfoListResponse, error) {
	client := group.NewGroupClientClient(m.cli.Conn())
	return client.MessageGroupInfoList(ctx, in, opts...)
}

func (m *defaultGroupClient) AddGroupChat(ctx context.Context, in *AddGroupChatRequest, opts ...grpc.CallOption) (*AddGroupChatResponse, error) {
	client := group.NewGroupClientClient(m.cli.Conn())
	return client.AddGroupChat(ctx, in, opts...)
}

func (m *defaultGroupClient) CreateGroupChat(ctx context.Context, in *CreateGroupChatRequest, opts ...grpc.CallOption) (*CreateGroupChatResponse, error) {
	client := group.NewGroupClientClient(m.cli.Conn())
	return client.CreateGroupChat(ctx, in, opts...)
}

func (m *defaultGroupClient) GetFriendListByUserId(ctx context.Context, in *FriendListRequest, opts ...grpc.CallOption) (*FriendListResponse, error) {
	client := group.NewGroupClientClient(m.cli.Conn())
	return client.GetFriendListByUserId(ctx, in, opts...)
}
