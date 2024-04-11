package logic

import (
	"context"

	"zero-chat/app/chat/cmd/rpc/internal/svc"
	"zero-chat/app/chat/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type SyncGetUnreadItemsLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSyncGetUnreadItemsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SyncGetUnreadItemsLogic {
	return &SyncGetUnreadItemsLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SyncGetUnreadItemsLogic) SyncGetUnreadItems(in *pb.SyncGetUnreadItemsReq) (*pb.SyncGetUnreadItemsResp, error) {
	// todo: add your logic here and delete this line

	return &pb.SyncGetUnreadItemsResp{}, nil
}
