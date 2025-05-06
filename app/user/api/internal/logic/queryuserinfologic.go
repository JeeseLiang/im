package logic

import (
	"context"

	"im_message/app/user/api/internal/svc"
	"im_message/app/user/api/internal/types"
	"im_message/app/user/rpc/userclient"

	"github.com/jinzhu/copier"
	"github.com/zeromicro/go-zero/core/logx"
)

type QueryUserInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewQueryUserInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryUserInfoLogic {
	return &QueryUserInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *QueryUserInfoLogic) QueryUserInfo(req *types.QueryUserInfoRequest) (*types.PersonalInfoResponse, error) {
	personInfo, err := l.svcCtx.UserRpc.PersonalInfo(l.ctx, &userclient.PersonalInfoRequest{
		Id: req.UserId,
	})
	if err != nil {
		return nil, err
	}
	var resp types.PersonalInfoResponse
	copier.Copy(&resp, personInfo)
	return &resp, nil
}
