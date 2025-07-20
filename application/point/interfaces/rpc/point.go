package main

import (
	"Ai-Novel/common/zlog"
	"flag"
	"fmt"

	"Ai-Novel/application/point/interfaces/rpc/internal/config"
	"Ai-Novel/application/point/interfaces/rpc/internal/server"
	"Ai-Novel/application/point/interfaces/rpc/internal/svc"
	"Ai-Novel/application/point/interfaces/rpc/point"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var configFile = flag.String("f", "etc/point-rpc.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)
	ctx := svc.NewServiceContext(c)

	// 注册自定义日志
	zlog.InitLogger(c.ServiceConf)

	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		point.RegisterPointRpcServer(grpcServer, server.NewPointRpcServer(ctx))

		if c.Mode == service.DevMode || c.Mode == service.TestMode {
			reflection.Register(grpcServer)
		}
	})
	defer s.Stop()

	//UserRpcServer := userrpc.NewUserRpc(zrpc.MustNewClient(c.UserRpc))
	//ping, err := UserRpcServer.Ping(context.Background(), &user.Ping{
	//	Message: "hello",
	//})
	//zlog.Infof("ping: %v, err: %v", ping, err)
	//if err != nil {
	//	return
	//}

	fmt.Printf("Starting rpc server at %s...\n", c.ListenOn)
	s.Start()
}
