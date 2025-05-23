package logic

import (
	"context"
	"time"

	"im_message/app/user/model"
	"im_message/app/user/rpc/internal/svc"
	"im_message/common/xcrypt"
	"im_message/common/xerr"
	"im_message/common/xjwt"
	"im_message/proto/user"

	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
)

type LoginLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *LoginLogic) Login(in *user.LoginRequest) (*user.LoginResponse, error) {
	// 查询用户是否存在
	userModel, err := l.svcCtx.UserModel.FindOneByEmail(l.ctx, in.Email)
	if err != nil {
		if err == model.ErrNotFound {
			return nil, errors.Wrapf(xerr.NewErrCode(xerr.NO_DATA),
				"user login email not exist, email:%s,err:%v", in.Email, err)
		}
		return nil, errors.Wrapf(xerr.NewErrMsg("unknown"),
			"user login email unknown, email:%s,err:%v", in.Email, err)
	}
	// 判断密码是否正确
	ok := xcrypt.PasswordVerify(in.Password, userModel.Password)
	if !ok {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.UNAUTHORIZED),
			"user login password error, password:%s", in.Password, err)
	}
	// 生成jwt
	accessSecret := l.svcCtx.Config.JwtAuth.AccessSecret
	accessExpire := l.svcCtx.Config.JwtAuth.AccessExpire
	accessToken, err := xjwt.GetJwtToken(accessSecret, time.Now().Unix(), accessExpire, userModel.Id)
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrMsg("generate token fail"),
			"getJwtToken err userId:%d, err:%v", userModel.Id, err)
	}
	return &user.LoginResponse{
		AccessToken:  accessToken,
		AccessExpire: accessExpire,
	}, nil
}
