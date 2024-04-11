package svc

import (
	"zero-chat/app/chat/cmd/rpc/internal/config"
	"zero-chat/app/chat/model"

	"github.com/aliyun/aliyun-tablestore-go-sdk/tablestore"
	"github.com/zeromicro/go-zero/core/stores/redis"
)

type ServiceContext struct {
	Config     config.Config
	RdsClient  *redis.Redis
	StoreTable model.StoreTableModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	client := tablestore.NewClient(c.TableStore.Endpoint, "zero-chat", c.TableStore.OtsAkEnv, c.TableStore.OtsSkEnv)
	return &ServiceContext{
		Config:     c,
		RdsClient:  redis.MustNewRedis(c.Redis.RedisConf),
		StoreTable: model.NewStoreTableModel("im_timeline_store_table", client),
	}
}
