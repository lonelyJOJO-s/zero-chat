package kafka

import (
	"context"
	"zero-chat/app/chat/cmd/api/internal/config"
	"zero-chat/app/chat/cmd/api/internal/svc"

	"github.com/zeromicro/go-queue/kq"
	"github.com/zeromicro/go-zero/core/service"
)

func Consumers(c config.Config, ctx context.Context, svcContext *svc.ServiceContext) []service.Service {

	return []service.Service{
		//Listening for changes in consumption flow status
		kq.MustNewQueue(c.KqConsumerConf, NewMessageBackMq(ctx, svcContext)),
		//.....
	}

}
