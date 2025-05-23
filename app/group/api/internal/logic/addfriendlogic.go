package logic

import (
	"context"

	"im_message/app/group/api/internal/svc"
	"im_message/app/group/api/internal/types"
	"im_message/common/ctxdata"
	"im_message/proto/group"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddFriendLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAddFriendLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddFriendLogic {
	return &AddFriendLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AddFriendLogic) AddFriend(req *types.AddFriendRequest) (*types.AddFriendResponse, error) {
	uid := ctxdata.GetUidFromCtx(l.ctx)
	_, err := l.svcCtx.GroupRpc.AddFriend(l.ctx, &group.AddFriendRequest{
		FromUid: uid,
		ToUid:   req.UserId,
	})
	if err != nil {
		return nil, err
	}
	return nil, nil
}
