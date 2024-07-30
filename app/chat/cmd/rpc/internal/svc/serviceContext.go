package svc

import (
	"zero-chat/app/chat/cmd/rpc/internal/config"
	"zero-chat/app/chat/cmd/rpc/internal/im"
	"zero-chat/app/user/cmd/rpc/client/friendservice"
	"zero-chat/app/user/cmd/rpc/client/groupservice"

	"github.com/aliyun/aliyun-tablestore-go-sdk/timeline"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config    config.Config
	RdsClient *redis.Redis
	// StoreTable model.StoreTableModel
	// SyncTable  model.SyncTableModel
	IM               *im.IM
	FriendServiceRpc friendservice.FriendService
	GroupServiceRpc  groupservice.GroupService
}

func NewServiceContext(c config.Config) *ServiceContext {
	// client := tablestore.NewClient(c.TableStore.Endpoint, "zero-chat", c.TableStore.OtsAkEnv, c.TableStore.OtsSkEnv)
	storeBuilder := timeline.StoreOption{
		Endpoint:  c.TableStore.Endpoint,
		Instance:  c.TableStore.Instance,
		TableName: "im_timeline_store",
		AkId:      c.TableStore.OtsAkEnv,
		AkSecret:  c.TableStore.OtsSkEnv,
		TTL:       365 * 24 * 3600, // Data time to alive, eg: almost one year
	}
	syncBuilder := timeline.StoreOption{
		Endpoint:  c.TableStore.Endpoint,
		Instance:  c.TableStore.Instance,
		TableName: "im_timeline_sync",
		AkId:      c.TableStore.OtsAkEnv,
		AkSecret:  c.TableStore.OtsSkEnv,
		TTL:       14 * 24 * 3600, // Data time to alive, eg: almost one month
	}
	im, err := im.NewIm(storeBuilder, syncBuilder, timeline.DefaultStreamAdapter)
	if err != nil {
		panic("im load err" + err.Error())
	}
	return &ServiceContext{
		Config:    c,
		RdsClient: redis.MustNewRedis(c.Redis.RedisConf),
		// StoreTable: model.NewStoreTableModel("im_timeline_store", client),
		// SyncTable:  model.NewSyncTableModel("im_timeline_sync", client),
		IM:               im,
		FriendServiceRpc: friendservice.NewFriendService(zrpc.MustNewClient(c.UsercenterRpcConf)),
		GroupServiceRpc:  groupservice.NewGroupService(zrpc.MustNewClient(c.UsercenterRpcConf)),
	}
}
