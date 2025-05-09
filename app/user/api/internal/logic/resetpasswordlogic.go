package logic

import (
	"context"

	"im_message/app/user/api/internal/svc"
	"im_message/app/user/api/internal/types"
	"im_message/app/user/rpc/userclient"

	"github.com/zeromicro/go-zero/core/logx"
)

type ResetPasswordLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewResetPasswordLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ResetPasswordLogic {
	return &ResetPasswordLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ResetPasswordLogic) ResetPassword(req *types.ResetPasswordRequest) (resp *types.ResetPasswordResponse, err error) {
	resetResp, err := l.svcCtx.UserRpc.ResetPassword(l.ctx, &userclient.ResetPasswordRequest{
		Email:       req.Email,
		OldPassword: req.OldPassword,
		NewPassword: req.NewPassword,
	})
	if err != nil {
		return nil, err
	}

	return &types.ResetPasswordResponse{
		Success: resetResp.Success,
	}, nil
}
