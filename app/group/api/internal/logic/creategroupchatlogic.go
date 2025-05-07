package logic

import (
	"context"

	"im_message/app/group/api/internal/svc"
	"im_message/app/group/api/internal/types"
	"im_message/app/group/rpc/groupclient"
	"im_message/common/ctxdata"

	"github.com/jinzhu/copier"
	"github.com/zeromicro/go-zero/core/logx"
)

type CreateGroupChatLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCreateGroupChatLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateGroupChatLogic {
	return &CreateGroupChatLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateGroupChatLogic) CreateGroupChat(req *types.CreateGroupChatRequest) (*types.CreateGroupChatResponse, error) {
	uid := ctxdata.GetUidFromCtx(l.ctx)
	res, err := l.svcCtx.GroupRpc.CreateGroupChat(l.ctx, &groupclient.CreateGroupChatRequest{
		UserId:    uid,
		GroupName: req.GroupName,
	})
	if err != nil {
		return nil, err
	}
	resp := types.CreateGroupChatResponse{}
	copier.Copy(&resp, res)
	return &resp, nil
}
