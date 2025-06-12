package svc

import (
	modelGroup "im_message/app/group/model"
	modelMsg "im_message/app/msg/model"
	modelUser "im_message/app/user/model"
	"im_message/app/user/rpc/internal/config"
	"im_message/common/xlock"

	"github.com/redis/go-redis/v9"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config         config.Config
	UserModel      modelUser.UserModel
	GroupModel     modelGroup.GroupModel
	GroupUserModel modelGroup.GroupUserModel
	ChatMsgModel   modelMsg.ChatMsgModel
	Redis          *redis.Client
	Lock           *xlock.RedisLock
}

func NewServiceContext(c config.Config) *ServiceContext {
	conn := sqlx.NewMysql(c.Db.DataSource)

	// 初始化Redis客户端
	redisClient := redis.NewClient(&redis.Options{
		Addr:     c.CacheRedis[0].Host,
		Password: c.CacheRedis[0].Pass,
		DB:       0,
	})

	// 初始化分布式锁
	lock, err := xlock.NewRedisLock(redisClient)
	if err != nil {
		panic(err)
	}

	return &ServiceContext{
		Config:         c,
		UserModel:      modelUser.NewUserModel(conn, c.CacheRedis),
		GroupModel:     modelGroup.NewGroupModel(conn, c.CacheRedis),
		GroupUserModel: modelGroup.NewGroupUserModel(conn, c.CacheRedis),
		ChatMsgModel:   modelMsg.NewChatMsgModel(conn, c.CacheRedis),
		Redis:          redisClient,
		Lock:           lock,
	}
}
