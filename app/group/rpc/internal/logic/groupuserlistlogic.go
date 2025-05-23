package logic

import (
	"context"

	"im_message/app/group/rpc/internal/svc"
	"im_message/common/xerr"
	"im_message/proto/group"

	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
)

type GroupUserListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGroupUserListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GroupUserListLogic {
	return &GroupUserListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 获取组内的所有用户
func (l *GroupUserListLogic) GroupUserList(in *group.GroupUserListRequest) (*group.GroupUserListResponse, error) {
	groupId := in.GroupId
	groupUsers, err := l.svcCtx.GroupUserModel.FindUserIdListByGroupId(l.ctx, groupId)
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "find group users error, groupId:%s, err:%v", groupId, err)
	}
	return &group.GroupUserListResponse{
		List: groupUsers,
	}, nil
}
