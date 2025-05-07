package logic

import (
	"context"
	"strconv"
	"strings"

	"im_message/app/group/rpc/internal/svc"
	"im_message/proto/group"

	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/mr"
)

type GetFriendListByUserIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetFriendListByUserIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetFriendListByUserIdLogic {
	return &GetFriendListByUserIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetFriendListByUserIdLogic) GetFriendListByUserId(in *group.FriendListRequest) (*group.FriendListResponse, error) {
	// 1. 从group_user表中查询userId的groupId列表
	// 2. 从group表中筛选出type和status为1的群组
	// 3. 从筛选出的群组中查询好友列表并组装返回(mapreduce)
	groupIds, err := l.svcCtx.GroupUserModel.FindGroupIdListByUserId(l.ctx, in.UserId)
	if err != nil {
		return nil, err
	}

	friendsList, err := mr.MapReduce[string, string, []int64](func(source chan<- string) {
		for _, id := range groupIds {
			source <- id
		}
	}, func(item string, writer mr.Writer[string], cancel func(error)) {
		arr := strings.Split(item, "_") // 筛选掉群聊
		if len(arr) != 2 {
			return
		}
		group, err := l.svcCtx.GroupModel.FindOne(l.ctx, item)
		if err != nil || (group.Type != 1 || group.Status != 1) { // 筛选掉单聊的无效群组
			return
		}
		writer.Write(group.Id)
	}, func(pipe <-chan string, writer mr.Writer[[]int64], cancel func(error)) {
		var res []int64
		for item := range pipe {
			uid, err := strconv.ParseInt(item, 10, 64)
			if err != nil {
				continue
			}
			res = append(res, uid)
		}
		writer.Write(res)
	})
	if err != nil {
		return nil, err
	}

	return &group.FriendListResponse{
		List: friendsList,
	}, nil
}
