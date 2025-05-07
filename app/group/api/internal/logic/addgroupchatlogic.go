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

type AddGroupChatLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAddGroupChatLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddGroupChatLogic {
	return &AddGroupChatLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AddGroupChatLogic) AddGroupChat(req *types.AddGroupChatRequest) (*types.AddGroupChatResponse, error) {
	uid := ctxdata.GetUidFromCtx(l.ctx)
	res, err := l.svcCtx.GroupRpc.AddGroupChat(l.ctx, &groupclient.AddGroupChatRequest{
		FromUid: uid,
		ToUid:   req.ToUid,
		GroupId: req.GroupId,
	})
	if err != nil {
		return nil, err
	}
	resp := types.AddGroupChatResponse{}
	copier.Copy(&resp, res)
	return &resp, nil
}
