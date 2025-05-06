package logic

import (
	"context"

	"im_message/app/group/rpc/internal/svc"
	"im_message/common/xerr"
	"im_message/proto/group"

	"github.com/pkg/errors"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserGroupListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUserGroupListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserGroupListLogic {
	return &UserGroupListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UserGroupListLogic) UserGroupList(in *group.UserGroupListRequest) (*group.UserGroupListResponse, error) {
	groupIdList, err := l.svcCtx.GroupUserModel.FindGroupIdListByUserId(l.ctx, in.UserId)
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "UserGroupList failed, userId: %v, err: %v", in.UserId, err)
	}
	return &group.UserGroupListResponse{List: groupIdList}, nil
}
