package logic

import (
	"context"
	"fmt"
	"time"

	"im_message/app/user/model"
	"im_message/app/user/rpc/internal/svc"
	"im_message/common/xerr"
	"im_message/proto/user"

	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateOnlineStatusLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateOnlineStatusLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateOnlineStatusLogic {
	return &UpdateOnlineStatusLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateOnlineStatusLogic) UpdateOnlineStatus(in *user.UpdateOnlineStatusRequest) (*user.UpdateOnlineStatusResponse, error) {
	// 生成Redis key
	onlineKey := fmt.Sprintf("%s%d", model.UserOnlineKeyPrefix, in.UserId)
	lastOnlineKey := fmt.Sprintf("%s%d", model.UserLastOnlineKeyPrefix, in.UserId)

	// 获取当前时间戳
	now := time.Now().Unix()

	// 使用Redis Pipeline批量执行命令
	pipe := l.svcCtx.Redis.Pipeline()

	if in.IsOnline {
		// 用户上线
		// 设置在线状态，过期时间设为5分钟
		pipe.Set(l.ctx, onlineKey, "1", 5*time.Minute)
		// 更新最后在线时间
		pipe.Set(l.ctx, lastOnlineKey, now, 0)
	} else {
		// 用户下线
		pipe.Del(l.ctx, onlineKey)
		// 更新最后在线时间
		pipe.Set(l.ctx, lastOnlineKey, now, 0)
	}

	// 执行Pipeline
	_, err := pipe.Exec(l.ctx)
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR),
			"update online status failed, userId:%d, err:%v", in.UserId, err)
	}

	return &user.UpdateOnlineStatusResponse{}, nil
}
