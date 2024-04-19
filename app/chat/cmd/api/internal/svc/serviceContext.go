package svc

import (
	"zero-chat/app/chat/cmd/api/internal/config"
	"zero-chat/app/chat/cmd/rpc/tableservice"
	"zero-chat/app/user/cmd/rpc/client/friendservice"
	"zero-chat/app/user/cmd/rpc/client/groupservice"
	"zero-chat/app/user/cmd/rpc/client/userservice"

	"github.com/IBM/sarama"
	"github.com/zeromicro/go-queue/kq"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config           config.Config
	UserServiceRpc   userservice.UserService
	GroupServiceRpc  groupservice.GroupService
	FriendServiceRpc friendservice.FriendService
	ChatServiceRpc   tableservice.TableService
	KqPusherClient   *kq.Pusher
	KqConsumerClient *sarama.Consumer
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:           c,
		UserServiceRpc:   userservice.NewUserService(zrpc.MustNewClient(c.UsercenterRpcConf)),
		GroupServiceRpc:  groupservice.NewGroupService(zrpc.MustNewClient(c.UsercenterRpcConf)),
		FriendServiceRpc: friendservice.NewFriendService(zrpc.MustNewClient(c.UsercenterRpcConf)),
		ChatServiceRpc:   tableservice.NewTableService(zrpc.MustNewClient(c.ChatRpcConf)),
		KqPusherClient:   kq.NewPusher(c.KqPusherConf.Brokers, c.KqPusherConf.Topic),
	}
}
