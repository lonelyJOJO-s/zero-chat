package svc

import (
	"zero-chat/app/chat/cmd/mq/internal/config"
	"zero-chat/app/chat/cmd/rpc/tableservice"
	"zero-chat/app/user/cmd/rpc/client/friendservice"

	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config           config.Config
	ChatServiceRpc   tableservice.TableService
	FriendServiceRpc friendservice.FriendService
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:           c,
		ChatServiceRpc:   tableservice.NewTableService(zrpc.MustNewClient(c.ChatRpcConf)),
		FriendServiceRpc: friendservice.NewFriendService(zrpc.MustNewClient(c.UserRpcConf)),
	}
}
