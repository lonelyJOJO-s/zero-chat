package logic

import (
	"context"

	"zero-chat/app/chat/cmd/rpc/internal/svc"
	"zero-chat/app/chat/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type SendLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSendLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SendLogic {
	return &SendLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SendLogic) Send(in *pb.SendReq) (*pb.Null, error) {
	// todo: add your logic here and delete this line

	return &pb.Null{}, nil
}
