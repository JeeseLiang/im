package svc

import (
	"time"

	modelGroup "im_message/app/group/model"
	"im_message/app/group/rpc/internal/config"
	modelMsg "im_message/app/msg/model"
	modelUser "im_message/app/user/model"
	"im_message/common/xlock"

	"github.com/redis/go-redis/v9"
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
	Redis          *redis.Client
	Lock           *xlock.RedisLock
}

func NewServiceContext(c config.Config) *ServiceContext {
	sqlConn := sqlx.NewMysql(c.Db.DataSource)

	// 初始化Redis客户端
	redisClient := redis.NewClient(&redis.Options{
		Addr:     c.Cache[0].Host,
		Password: c.Cache[0].Pass,
		DB:       0,
	})

	// 初始化分布式锁
	lock, err := xlock.NewRedisLock(redisClient)
	if err != nil {
		panic(err)
	}

	mqWriter := &kafka.Writer{
		Addr:         kafka.TCP(c.MqConf.Brokers...),
		Topic:        c.MqConf.Topic,
		BatchTimeout: time.Millisecond * 20,
	}
	return &ServiceContext{
		Config:         c,
		GroupModel:     modelGroup.NewGroupModel(sqlConn, c.Cache, lock),
		GroupUserModel: modelGroup.NewGroupUserModel(sqlConn, c.Cache),
		UserModel:      modelUser.NewUserModel(sqlConn, c.Cache),
		ChatMsgModel:   modelMsg.NewChatMsgModel(sqlConn, c.Cache),
		MqWriter:       mqWriter,
		Redis:          redisClient,
		Lock:           lock,
	}
}
