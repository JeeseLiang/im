package logic

import (
	"context"
	"database/sql"

	"im_message/app/user/rpc/internal/svc"
	"im_message/proto/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type ModifyUserInfoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewModifyUserInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ModifyUserInfoLogic {
	return &ModifyUserInfoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ModifyUserInfoLogic) ModifyUserInfo(in *user.ModifyUserInfoRequest) (*user.ModifyUserInfoResponse, error) {
	userInfo, err := l.svcCtx.UserModel.FindOne(l.ctx, in.UserId)
	if err != nil {
		return nil, err
	}
	userInfo.NickName = in.NickName
	userInfo.Gender = in.Gender
	userInfo.AvatarUrl = sql.NullString{String: in.AvatarUrl, Valid: true}

	err = l.svcCtx.UserModel.Update(l.ctx, userInfo)
	if err != nil {
		return nil, err
	}
	return &user.ModifyUserInfoResponse{
		Success: true,
	}, nil
}
