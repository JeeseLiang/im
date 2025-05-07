package logic

import (
	"context"

	"im_message/app/group/rpc/internal/svc"
	"im_message/proto/group"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddGroupChatLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddGroupChatLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddGroupChatLogic {
	return &AddGroupChatLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *AddGroupChatLogic) AddGroupChat(in *group.AddGroupChatRequest) (*group.AddGroupChatResponse, error) {
	// todo: add your logic here and delete this line

	return &group.AddGroupChatResponse{}, nil
}
