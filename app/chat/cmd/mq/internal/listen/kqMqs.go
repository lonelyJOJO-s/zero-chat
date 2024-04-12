package listen

import (
	"context"
	"zero-chat/app/chat/cmd/mq/internal/config"

	kqMq "zero-chat/app/chat/cmd/mq/internal/mqs/kq"
	"zero-chat/app/chat/cmd/mq/internal/svc"

	"github.com/zeromicro/go-queue/kq"
	"github.com/zeromicro/go-zero/core/service"
)

// pub sub use kq (kafka)
func KqMqs(c config.Config, ctx context.Context, svcContext *svc.ServiceContext) []service.Service {

	return []service.Service{
		//Listening for changes in consumption flow status
		kq.MustNewQueue(c.KqConsumerConf, kqMq.NewMessageTransferMq(ctx, svcContext)),
		//.....
	}

}
