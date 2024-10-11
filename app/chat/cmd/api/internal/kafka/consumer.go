package kafka

import (
	"context"
	"zero-chat/app/chat/cmd/api/internal/svc"
	"zero-chat/app/chat/cmd/api/internal/ws"
)

type MessageBackMq struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewMessageBackMq(ctx context.Context, svcCtx *svc.ServiceContext) *MessageBackMq {
	return &MessageBackMq{
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *MessageBackMq) Consume(ctx context.Context, key, val string) error {
	ws.WsServer.Broadcast <- []byte(val)
	return nil
}
