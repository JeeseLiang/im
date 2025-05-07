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

// 添加群组成员(是好友则无需同意，不是好友无法邀请)
func (l *AddGroupChatLogic) AddGroupChat(in *group.AddGroupChatRequest) (*group.AddGroupChatResponse, error) {
	// 1. 查询FromUid的好友列表
	// 2. 筛选出可以被添加的ToUid
	// 3. 添加群成员
	return &group.AddGroupChatResponse{}, nil
}
