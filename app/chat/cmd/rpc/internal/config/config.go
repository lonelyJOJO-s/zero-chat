package config

import "github.com/zeromicro/go-zero/zrpc"

type Config struct {
	zrpc.RpcServerConf
	TableStore struct {
		OtsAkEnv string
		OtsSkEnv string
		Endpoint string
	}
}
