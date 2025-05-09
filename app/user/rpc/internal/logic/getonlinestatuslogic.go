package logic

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"im_message/app/user/model"
	"im_message/app/user/rpc/internal/svc"
	"im_message/common/xerr"
	"im_message/proto/user"

	"github.com/pkg/errors"
	"github.com/redis/go-redis/v9"
	"github.com/zeromicro/go-zero/core/logx"
)

type GetOnlineStatusLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetOnlineStatusLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetOnlineStatusLogic {
	return &GetOnlineStatusLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetOnlineStatusLogic) GetOnlineStatus(in *user.GetOnlineStatusRequest) (*user.GetOnlineStatusResponse, error) {
	// 生成Redis key
	onlineKey := fmt.Sprintf("%s%d", model.UserOnlineKeyPrefix, in.UserId)
	lastOnlineKey := fmt.Sprintf("%s%d", model.UserLastOnlineKeyPrefix, in.UserId)

	// 使用Redis Pipeline批量执行命令
	pipe := l.svcCtx.Redis.Pipeline()

	// 获取在线状态
	onlineCmd := pipe.Get(l.ctx, onlineKey)
	// 获取最后在线时间
	lastOnlineCmd := pipe.Get(l.ctx, lastOnlineKey)

	// 执行Pipeline
	_, err := pipe.Exec(l.ctx)
	if err != nil && err != redis.Nil {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR),
			"get online status failed, userId:%d, err:%v", in.UserId, err)
	}

	// 处理在线状态
	isOnline := false
	if onlineCmd.Val() == "1" {
		isOnline = true
	}

	// 处理最后在线时间
	lastOnlineTime := time.Now().Unix()
	if lastOnlineCmd.Val() != "" {
		lastOnlineTime, err = strconv.ParseInt(lastOnlineCmd.Val(), 10, 64)
		if err != nil {
			return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR),
				"parse last online time failed, userId:%d, err:%v", in.UserId, err)
		}
	}

	// 返回结果
	return &user.GetOnlineStatusResponse{
		IsOnline:       isOnline,
		LastOnlineTime: lastOnlineTime,
	}, nil
}
