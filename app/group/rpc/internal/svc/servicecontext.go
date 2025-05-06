package svc

import (
	"time"

	modelGroup "im_message/app/group/model"
	"im_message/app/group/rpc/internal/config"
	modelMsg "im_message/app/msg/model"
	modelUser "im_message/app/user/model"

	"github.com/segmentio/kafka-go"

	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config         config.Config
	GroupModel     modelGroup.GroupModel
	GroupUserModel modelGroup.GroupUserModel
	UserModel      modelUser.UserModel
	ChatMsgModel   modelMsg.ChatMsgModel
	MqWriter       *kafka.Writer
}

func NewServiceContext(c config.Config) *ServiceContext {
	sqlConn := sqlx.NewMysql(c.Db.DataSource)
	mqWriter := &kafka.Writer{
		Addr:         kafka.TCP(c.MqConf.Brokers...),
		Topic:        c.MqConf.Topic,
		BatchTimeout: time.Millisecond * 20,
	}
	return &ServiceContext{
		Config:         c,
		GroupModel:     modelGroup.NewGroupModel(sqlConn, c.Cache),
		GroupUserModel: modelGroup.NewGroupUserModel(sqlConn, c.Cache),
		UserModel:      modelUser.NewUserModel(sqlConn, c.Cache),
		ChatMsgModel:   modelMsg.NewChatMsgModel(sqlConn, c.Cache),
		MqWriter:       mqWriter,
	}
}
