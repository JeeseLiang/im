package logic

import (
	"context"

	"im_message/app/user/rpc/internal/svc"
	"im_message/common/xcrypt"
	"im_message/common/xerr"
	"im_message/proto/user"

	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
)

type ResetPasswordLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewResetPasswordLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ResetPasswordLogic {
	return &ResetPasswordLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ResetPasswordLogic) ResetPassword(in *user.ResetPasswordRequest) (*user.ResetPasswordResponse, error) {
	userModel, err := l.svcCtx.UserModel.FindOneByEmail(l.ctx, in.Email)
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DATA_EXIST),
			"User is not exist, err: %s", err)
	}

	ok := xcrypt.PasswordVerify(in.OldPassword, userModel.Password)
	if !ok {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DATA_EXIST),
			"user  password error, err :%s", err)
	}
	hashedPwd, err := xcrypt.PasswordHash(in.NewPassword)
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrMsg("密码处理失败"),
			"user hash error, err: %s", err)
	}
	userModel.Password = hashedPwd
	err = l.svcCtx.UserModel.Update(l.ctx, userModel)
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrMsg("更新密码失败"),
			"user password update error, err: %s", err)
	}
	return &user.ResetPasswordResponse{
		Success: true,
	}, nil
}
