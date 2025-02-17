package config

import (
	"github.com/zeromicro/go-queue/kq"
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	rest.RestConf
	UsercenterRpcConf zrpc.RpcClientConf
	ChatRpcConf       zrpc.RpcClientConf
	JwtAuth           struct {
		AccessSecret string
		AccessExpire int64
	}
	KqPusherConf struct {
		Brokers []string
		Topic   string
	}
	KqConsumerConf kq.KqConf
}
