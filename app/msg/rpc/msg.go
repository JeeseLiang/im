package main

import (
	"flag"
	"fmt"

	"im_message/common/interceptor"

	"github.com/joho/godotenv"

	"im_message/app/msg/rpc/internal/config"
	"im_message/app/msg/rpc/internal/server"
	"im_message/app/msg/rpc/internal/svc"
	"im_message/proto/msg"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var configFile = flag.String("f", "etc/msg.yaml", "the config file")

func main() {
	// 读取.env
	err := godotenv.Load("../../../.env")
	if err != nil {
		logx.Infof("加载 .env 文件失败: %v", err)
	}
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c, conf.UseEnv())
	ctx := svc.NewServiceContext(c)
	svr := server.NewMessageClientServer(ctx)

	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		msg.RegisterMessageClientServer(grpcServer, svr)

		if c.Mode == service.DevMode || c.Mode == service.TestMode {
			reflection.Register(grpcServer)
		}
	})
	defer s.Stop()

	// 添加拦截器
	s.AddUnaryInterceptors(interceptor.LoggerInterceptor)
	// 禁用显示cpu
	logx.DisableStat()

	fmt.Printf("Starting rpc server at %s...\n", c.ListenOn)
	s.Start()
}
