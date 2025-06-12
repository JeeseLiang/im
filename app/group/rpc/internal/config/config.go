package config

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	zrpc.RpcServerConf
	Db struct {
		DataSource string
	}
	Cache      cache.CacheConf
	CacheRedis []struct {
		Host string
		Pass string
	}
	MqConf struct {
		Brokers []string
		Topic   string
	}
}
