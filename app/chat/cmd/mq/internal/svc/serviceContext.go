package svc

import (
	"zero-chat/app/chat/cmd/mq/internal/config"
	"zero-chat/app/chat/cmd/rpc/tableservice"
	"zero-chat/app/user/cmd/rpc/client/friendservice"
	"zero-chat/app/user/cmd/rpc/client/groupservice"

	"github.com/zeromicro/go-queue/kq"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config           config.Config
	ChatServiceRpc   tableservice.TableService
	FriendServiceRpc friendservice.FriendService
	GroupServiceRpc  groupservice.GroupService
	Redis            *redis.Redis
	KqPusherClient   *kq.Pusher
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:           c,
		ChatServiceRpc:   tableservice.NewTableService(zrpc.MustNewClient(c.ChatRpcConf)),
		FriendServiceRpc: friendservice.NewFriendService(zrpc.MustNewClient(c.UserRpcConf)),
		GroupServiceRpc:  groupservice.NewGroupService(zrpc.MustNewClient(c.UserRpcConf)),
		Redis:            redis.MustNewRedis(c.Redis),
		KqPusherClient:   kq.NewPusher(c.KqPusherConf.Brokers, c.KqPusherConf.Topic),
	}
}
