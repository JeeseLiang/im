package logic

import (
	"context"

	"im_message/app/msg/api/internal/svc"
	"im_message/app/msg/api/internal/types"
	"im_message/common/ctxdata"
	"im_message/proto/msg"

	"github.com/jinzhu/copier"
	"github.com/zeromicro/go-zero/core/logx"
)

type PullLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewPullLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PullLogic {
	return &PullLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *PullLogic) Pull(req *types.PullRequest) (*types.PullResponse, error) {
	uid := ctxdata.GetUidFromCtx(l.ctx)
	var pbPullRequest msg.PullRequest
	copier.Copy(&pbPullRequest, req)
	pbPullRequest.UserId = uid
	pbPullResponse, err := l.svcCtx.MsgRpc.Pull(l.ctx, &pbPullRequest)
	if err != nil {
		return nil, err
	}
	var resp types.PullResponse
	copier.Copy(&resp, pbPullResponse)
	return &resp, nil
}
