package main

import (
	"flag"
	"fmt"

	"im_message/app/user/api/internal/config"
	"im_message/app/user/api/internal/handler"
	"im_message/app/user/api/internal/svc"

	"github.com/joho/godotenv"
	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/rest"
)

var configFile = flag.String("f", "etc/user-api.yaml", "the config file")

func main() {
	// 读取.env
	err := godotenv.Load("../../../.env")
	if err != nil {
		logx.Infof("加载 .env 文件失败: %v", err)
	}

	// fmt.Println(os.Getenv("MYSQL_PASSWORD"), "\n")

	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c, conf.UseEnv())

	ctx := svc.NewServiceContext(c)
	server := rest.MustNewServer(c.RestConf)
	defer server.Stop()

	handler.RegisterHandlers(server, ctx)

	logx.DisableStat()
	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}
