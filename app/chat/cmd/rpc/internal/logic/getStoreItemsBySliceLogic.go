package logic

import (
	"context"

	"zero-chat/app/chat/cmd/rpc/internal/svc"
	"zero-chat/app/chat/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetStoreItemsBySliceLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetStoreItemsBySliceLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetStoreItemsBySliceLogic {
	return &GetStoreItemsBySliceLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetStoreItemsBySliceLogic) GetStoreItemsBySlice(in *pb.GetStoreItemsBySliceReq) (*pb.GetStoreItemsBySliceResp, error) {
	// todo: add your logic here and delete this line

	return &pb.GetStoreItemsBySliceResp{}, nil
}
