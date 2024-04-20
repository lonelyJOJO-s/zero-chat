package main

import (
	"flag"
	"fmt"

	"zero-chat/app/chat/cmd/api/internal/config"
	"zero-chat/app/chat/cmd/api/internal/handler"
	"zero-chat/app/chat/cmd/api/internal/kafka"
	"zero-chat/app/chat/cmd/api/internal/svc"
	"zero-chat/app/chat/cmd/api/internal/ws"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
)

var configFile = flag.String("f", "etc/chat.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)

	server := rest.MustNewServer(c.RestConf, rest.WithCors("*"))
	defer server.Stop()

	ctx := svc.NewServiceContext(c)
	// start ws server
	go ws.WsServer.Run(ctx.GroupServiceRpc)
	// start kafka listen
	go kafka.StartConsume(c)
	handler.RegisterHandlers(server, ctx)

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}
