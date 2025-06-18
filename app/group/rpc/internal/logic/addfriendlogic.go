package logic

import (
	"context"
	"strings"
	"time"

	"im_message/app/group/model"
	"im_message/app/group/rpc/internal/svc"
	"im_message/common/biz"
	"im_message/common/xerr"
	"im_message/common/xlock"
	"im_message/common/xmq"
	"im_message/proto/group"
	"im_message/proto/msg"

	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
)

type AddFriendLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddFriendLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddFriendLogic {
	return &AddFriendLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 添加好友
func (l *AddFriendLogic) AddFriend(in *group.AddFriendRequest) (*group.AddFriendResponse, error) {
	fromUid := in.FromUid
	toUid := in.ToUid

	// 创建分布式锁
	lock, err := xlock.NewRedisLock(l.svcCtx.Redis)
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.SERVER_ERROR), "获取锁失败: %v", err)
	}
	lockKey := biz.GetGroupId(fromUid, toUid)
	lockValue := time.Now().String()

	// 尝试获取锁，设置10秒超时
	err = lock.Acquire(l.ctx, lockKey, lockValue, 10*time.Second)
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.SERVER_ERROR), "获取锁失败: %v", err)
	}
	defer lock.Release(l.ctx, lockKey, lockValue)

	// 检查是否已经是好友
	groupId := biz.GetGroupId(fromUid, toUid)
	existingGroup, err := l.svcCtx.GroupModel.FindOne(l.ctx, groupId)
	if err == nil && existingGroup.Status == model.GroupStatusYes {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.CLIENT_ERROR), "已经是好友关系")
	}

	// 查询这两个用户的nickName
	u1, err := l.svcCtx.UserModel.FindOne(l.ctx, fromUid)
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "AddFriend query user failed, fromUid: %v", fromUid)
	}
	u2, err := l.svcCtx.UserModel.FindOne(l.ctx, toUid)
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "AddFriend query user failed, toUid: %v", toUid)
	}
	groupName := strings.Join([]string{u1.NickName, u2.NickName}, ", ")

	// 创建一个group
	_, err = l.svcCtx.GroupModel.Insert(l.ctx, &model.Group{
		Id:     groupId,
		Name:   groupName,
		Type:   1,
		Status: model.GroupStatusNo,
	})
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "AddFriend insert group failed: %v", err)
	}

	// 加好友请求消息 放入消息队列
	chatMsg := &msg.ChatMsg{
		GroupId:    groupId,
		SenderId:   fromUid,
		Type:       0,
		Content:    "请求加你为好友",
		Uuid:       biz.GetUuid(),
		CreateTime: time.Now().UnixMilli(),
	}
	err = xmq.PushToMq(l.ctx, l.svcCtx.MqWriter, chatMsg)
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.SERVER_ERROR), "AddFriend push to mq failed: %v", err)
	}

	return &group.AddFriendResponse{
		GroupId: groupId,
	}, nil
}
