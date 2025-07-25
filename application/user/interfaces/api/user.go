package main

import (
	"Ai-Novel/application/user/interfaces/api/internal/config"
	"Ai-Novel/application/user/interfaces/api/internal/handler"
	"Ai-Novel/application/user/interfaces/api/internal/svc"
	"Ai-Novel/common/zlog"
	"flag"
	"fmt"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
)

var configFile = flag.String("f", "etc/user-api.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)

	server := rest.MustNewServer(c.RestConf)
	defer server.Stop()

	ctx := svc.NewServiceContext(c)
	handler.RegisterHandlers(server, ctx)

	// 注册自定义日志
	zlog.InitLogger(c.ServiceConf)

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}
