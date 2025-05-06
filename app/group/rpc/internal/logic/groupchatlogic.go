package logic

import (
	"context"

	"im_message/app/group/rpc/internal/svc"
	"im_message/proto/group"

	"github.com/zeromicro/go-zero/core/logx"
)

type GroupChatLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGroupChatLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GroupChatLogic {
	return &GroupChatLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GroupChatLogic) GroupChat(in *group.GroupChatRequest) (*group.GroupChatResponse, error) {
	// todo: add your logic here and delete this line

	return &group.GroupChatResponse{}, nil
}
