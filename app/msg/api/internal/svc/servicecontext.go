package svc

import (
	"im_message/app/group/rpc/groupclient"
	"im_message/app/msg/api/internal/config"
	"im_message/app/msg/rpc/messageclient"

	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config   config.Config
	MsgRpc   messageclient.MessageClient
	GroupRpc groupclient.GroupClient
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:   c,
		MsgRpc:   messageclient.NewMessageClient(zrpc.MustNewClient(c.MsgRpc)),
		GroupRpc: groupclient.NewGroupClient(zrpc.MustNewClient(c.GroupRpc)),
	}
}
