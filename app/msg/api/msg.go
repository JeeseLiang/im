package main

import (
	"flag"
	"fmt"
	"net/http"

	"im_message/app/msg/api/internal/config"
	"im_message/app/msg/api/internal/handler"
	"im_message/app/msg/api/internal/logic"
	"im_message/app/msg/api/internal/svc"
	"im_message/common/xresp"

	"github.com/joho/godotenv"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/rest"
)

var configFile = flag.String("f", "etc/msg.yaml", "the config file")

func main() {
	// 读取.env
	err := godotenv.Load("../../../.env")
	if err != nil {
		logx.Errorf("加载 .env 文件失败: %v", err)
	}
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c, conf.UseEnv())

	ctx := svc.NewServiceContext(c)
	server := rest.MustNewServer(c.RestConf)
	defer server.Stop()

	handler.RegisterHandlers(server, ctx)
	logx.DisableStat() // 禁用显示cpu
	// http升级为websocket
	server.AddRoute(
		rest.Route{
			Method: http.MethodGet,
			Path:   "/ws",
			Handler: func(w http.ResponseWriter, r *http.Request) {
				err := logic.ServeWs(ctx, w, r)
				xresp.Response(r, w, nil, err)
			},
		},
		rest.WithJwt(ctx.Config.JwtAuth.AccessSecret),
	)
	// 开启协程, 专门从MQ中获取消息, 发给对应的群
	go logic.ConsumeMsgFromMQ(ctx)

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}
