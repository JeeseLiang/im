package logic

import (
	"context"

	"im_message/app/user/api/internal/svc"
	"im_message/app/user/api/internal/types"
	"im_message/common/ctxdata"
	"im_message/proto/user"

	"github.com/jinzhu/copier"
	"github.com/zeromicro/go-zero/core/logx"
)

type ModifyPersonalInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewModifyPersonalInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ModifyPersonalInfoLogic {
	return &ModifyPersonalInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ModifyPersonalInfoLogic) ModifyPersonalInfo(req *types.ModifyPersonalInfoRequest) (resp *types.ModifyPersonalInfoResponse, err error) {
	uid := ctxdata.GetUidFromCtx(l.ctx)
	res, err := l.svcCtx.UserRpc.ModifyUserInfo(l.ctx, &user.ModifyUserInfoRequest{
		UserId:    uid,
		NickName:  req.NickName,
		Gender:    req.Gender,
		AvatarUrl: req.AvatarUrl,
	})
	copier.Copy(&resp, res)
	return
}
