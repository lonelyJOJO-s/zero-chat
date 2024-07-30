package main

import (
	"flag"
	"fmt"
	"os"

	"zero-chat/app/user/cmd/rpc/internal/config"
	friendserviceServer "zero-chat/app/user/cmd/rpc/internal/server/friendservice"
	groupserviceServer "zero-chat/app/user/cmd/rpc/internal/server/groupservice"
	userserviceServer "zero-chat/app/user/cmd/rpc/internal/server/userservice"
	"zero-chat/app/user/cmd/rpc/internal/svc"
	"zero-chat/app/user/cmd/rpc/pb"
	"zero-chat/common/interceptor/rpcserver"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var configFile = flag.String("f", "etc/usercenter.yaml", "the config file")

func main() {
	flag.Parse()
	logx.SetWriter(logx.NewWriter(os.Stdout))
	var c config.Config
	conf.MustLoad(*configFile, &c)
	ctx := svc.NewServiceContext(c)

	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		pb.RegisterUserServiceServer(grpcServer, userserviceServer.NewUserServiceServer(ctx))
		pb.RegisterFriendServiceServer(grpcServer, friendserviceServer.NewFriendServiceServer(ctx))
		pb.RegisterGroupServiceServer(grpcServer, groupserviceServer.NewGroupServiceServer(ctx))

		if c.Mode == service.DevMode || c.Mode == service.TestMode {
			reflection.Register(grpcServer)
		}
	})
	defer s.Stop()
	s.AddUnaryInterceptors(rpcserver.LoggerInterceptor)

	fmt.Printf("Starting rpc server at %s...\n", c.ListenOn)
	s.Start()
}
