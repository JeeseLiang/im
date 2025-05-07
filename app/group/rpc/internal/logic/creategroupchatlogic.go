package logic

import (
	"context"
	"database/sql"

	"im_message/app/group/model"
	"im_message/app/group/rpc/internal/svc"
	modelMsg "im_message/app/msg/model"
	"im_message/common/utils"
	"im_message/common/xerr"
	"im_message/proto/group"

	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type CreateGroupChatLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateGroupChatLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateGroupChatLogic {
	return &CreateGroupChatLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CreateGroupChatLogic) CreateGroupChat(in *group.CreateGroupChatRequest) (*group.CreateGroupChatResponse, error) {
	// 使用事务创建一个group和添加自己为groupUser
	u1, _ := l.svcCtx.UserModel.FindOne(l.ctx, in.UserId)
	var groupId string
	err := l.svcCtx.GroupModel.Trans(l.ctx, func(ctx context.Context, session sqlx.Session) error {
		var err error
		groupId, err = l.svcCtx.GroupModel.CreateGroup(ctx, session, in.UserId, in.GroupName)
		if err != nil {
			return errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "CreateGroup error:%v", err)
		}
		groupUser1 := &model.GroupUser{
			GroupId:   groupId,
			UserId:    in.UserId,
			AliasName: sql.NullString{String: u1.NickName, Valid: true},
		}
		_, err = l.svcCtx.GroupUserModel.TransInsert(ctx, session, groupUser1)
		if err != nil {
			return errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "TransInsert error:%v", err)
		}
		chatMsg := &modelMsg.ChatMsg{
			GroupId:  groupId,
			SenderId: u1.Id,
			Type:     modelMsg.MsgTypeText,
			Content:  "成功创建群聊",
			Uuid:     utils.GenUuid(),
		}
		_, err = l.svcCtx.ChatMsgModel.TransInsert(ctx, session, chatMsg)
		if err != nil {
			return errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "TransInsert say hello error:%v", err)
		}
		return nil // commit
	})
	if err != nil {
		return nil, err
	}
	return &group.CreateGroupChatResponse{
		GroupId: groupId,
	}, nil
}
