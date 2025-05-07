package logic

import (
	"context"

	"im_message/app/group/model"
	"im_message/app/group/rpc/internal/svc"
	"im_message/proto/group"

	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/mr"
)

type AddGroupChatLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddGroupChatLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddGroupChatLogic {
	return &AddGroupChatLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 添加群组成员(是好友则无需同意，不是好友无法邀请)
func (l *AddGroupChatLogic) AddGroupChat(in *group.AddGroupChatRequest) (*group.AddGroupChatResponse, error) {
	// 1. 查询FromUid的好友列表
	// 2. 筛选出可以被添加的ToUid
	// 3. 添加群成员
	friendsList, err := NewGetFriendListByUserIdLogic(l.ctx, l.svcCtx).GetFriendListByUserId(&group.FriendListRequest{UserId: in.FromUid})
	if err != nil {
		return nil, err
	}
	mp := map[int64]struct{}{}
	for _, uid := range in.ToUid {
		mp[uid] = struct{}{}
	}
	allowedUids := make([]int64, 0)
	for _, uid := range friendsList.List {
		if _, ok := mp[uid]; ok {
			allowedUids = append(allowedUids, uid)
		}
	}
	cnt, err := mr.MapReduce[int64, struct{}, int64](func(source chan<- int64) {
		for _, uid := range allowedUids {
			source <- uid
		}
	}, func(item int64, writer mr.Writer[struct{}], cancel func(error)) {
		_, err := l.svcCtx.GroupUserModel.Insert(l.ctx, &model.GroupUser{
			GroupId: in.GroupId,
			UserId:  item,
		})
		if err != nil {
			return
		}
		writer.Write(struct{}{})
	}, func(pipe <-chan struct{}, writer mr.Writer[int64], cancel func(error)) {
		count := 0
		for range pipe {
			count++
		}
		writer.Write(int64(count))
	})
	if err != nil {
		return nil, err
	}
	return &group.AddGroupChatResponse{
		Cnt: cnt,
	}, nil
}
