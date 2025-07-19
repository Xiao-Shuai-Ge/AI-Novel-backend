package main

import (
	"Ai-Novel/common/zlog"
	"flag"
	"fmt"

	"Ai-Novel/application/user/interfaces/rpc/internal/config"
	"Ai-Novel/application/user/interfaces/rpc/internal/server"
	"Ai-Novel/application/user/interfaces/rpc/internal/svc"
	"Ai-Novel/application/user/interfaces/rpc/user"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var configFile = flag.String("f", "etc/user-rpc.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)
	ctx := svc.NewServiceContext(c)

	// 注册自定义日志
	zlog.InitLogger(c.ServiceConf)

	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		user.RegisterUserRpcServer(grpcServer, server.NewUserRpcServer(ctx))

		if c.Mode == service.DevMode || c.Mode == service.TestMode {
			reflection.Register(grpcServer)
		}
	})
	defer s.Stop()

	fmt.Printf("Starting rpc server at %s...\n", c.ListenOn)
	s.Start()
}
