package main

import (
	"Ai-Novel/common/websocketx"
	"Ai-Novel/common/zlog"
	"flag"
	"fmt"
	"net/http"

	"Ai-Novel/application/ai/interfaces/api/internal/config"
	"Ai-Novel/application/ai/interfaces/api/internal/handler"
	"Ai-Novel/application/ai/interfaces/api/internal/svc"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
)

var configFile = flag.String("f", "etc/ai-api.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)

	// 跨域设置
	server := rest.MustNewServer(c.RestConf, rest.WithCustomCors(nil, func(w http.ResponseWriter) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Headers", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS, PATCH")
		w.Header().Set("Access-Control-Expose-Headers", "Content-Length, Content-Type, Access-Control-Allow-Origin, Access-Control-Allow-Headers")
		w.Header().Set("Access-Control-Allow-Credentials", "true")
	}, "*"))
	defer server.Stop()

	// 启动 websocket 服务
	websocketx.WebsocketManager = websocketx.NewClientManager()
	go websocketx.WebsocketManager.Start()

	ctx := svc.NewServiceContext(c)
	handler.RegisterHandlers(server, ctx)

	// 注册自定义日志
	zlog.InitLogger(c.ServiceConf)

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}
