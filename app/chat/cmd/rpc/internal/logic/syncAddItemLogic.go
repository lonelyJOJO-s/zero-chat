package logic

import (
	"context"

	"zero-chat/app/chat/cmd/rpc/internal/svc"
	"zero-chat/app/chat/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type SyncAddItemLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSyncAddItemLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SyncAddItemLogic {
	return &SyncAddItemLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// sync table
func (l *SyncAddItemLogic) SyncAddItem(in *pb.SyncAddItemReq) (*pb.SyncAddItemResp, error) {
	// todo: add your logic here and delete this line

	return &pb.SyncAddItemResp{}, nil
}
